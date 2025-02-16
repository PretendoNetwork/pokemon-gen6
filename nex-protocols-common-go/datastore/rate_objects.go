package datastore

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func (commonProtocol *CommonProtocol) rateObjects(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*datastore_types.DataStoreRatingTarget], params *types.List[*datastore_types.DataStoreRateObjectParam], transactional *types.PrimitiveBool, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.GetObjectInfoByDataIDWithPassword == nil {
		common_globals.Logger.Warning("GetObjectInfoByDataIDWithPassword not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if commonProtocol.RateObjectWithPassword == nil {
		common_globals.Logger.Warning("RateObjectWithPassword not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.Unknown, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	pRatings := types.NewList[*datastore_types.DataStoreRatingInfo]()
	pResults := types.NewList[*types.QResult]()

	pRatings.Type = datastore_types.NewDataStoreRatingInfo()
	pResults.Type = types.NewQResult(0)

	// * Real DataStore does not actually check this.
	// * I just didn't feel like working out the
	// * logic for differing sized lists. So force
	// * them to always be the same
	if targets.Length() != params.Length() {
		return nil, nex.NewError(nex.ResultCodes.DataStore.InvalidArgument, "change_error")
	}

	var errorCode *nex.Error

	targets.Each(func(i int, target *datastore_types.DataStoreRatingTarget) bool {
		param, err := params.Get(i)
		if err != nil {
			errorCode = nex.NewError(nex.ResultCodes.DataStore.InvalidArgument, "change_error")
			return true
		}

		objectInfo, errCode := commonProtocol.GetObjectInfoByDataIDWithPassword(target.DataID, param.AccessPassword)
		if errCode != nil {
			errorCode = errCode
			return true
		}

		errCode = commonProtocol.VerifyObjectPermission(objectInfo.OwnerID, connection.PID(), objectInfo.Permission)
		if errCode != nil {
			errorCode = errCode
			return true
		}

		rating, errCode := commonProtocol.RateObjectWithPassword(target.DataID, target.Slot, param.RatingValue, param.AccessPassword)
		if errCode != nil {
			errorCode = errCode
			return true
		}

		if fetchRatings.Value {
			pRatings.Append(rating)
		}

		return false
	})

	if errorCode != nil {
		return nil, errorCode
	}

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pRatings.WriteTo(rmcResponseStream)
	pResults.WriteTo(rmcResponseStream) // * pResults is ALWAYS empty in SMM?

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodRateObjects
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterRateObjects != nil {
		go commonProtocol.OnAfterRateObjects(packet, targets, params, transactional, fetchRatings)
	}

	return rmcResponse, nil
}
