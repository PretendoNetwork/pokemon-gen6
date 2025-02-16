package datastore

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func (commonProtocol *CommonProtocol) rateObject(err error, packet nex.PacketInterface, callID uint32, target *datastore_types.DataStoreRatingTarget, param *datastore_types.DataStoreRateObjectParam, fetchRatings *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error) {
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

	objectInfo, errCode := commonProtocol.GetObjectInfoByDataIDWithPassword(target.DataID, param.AccessPassword)
	if errCode != nil {
		return nil, errCode
	}

	errCode = commonProtocol.VerifyObjectPermission(objectInfo.OwnerID, connection.PID(), objectInfo.Permission)
	if errCode != nil {
		return nil, errCode
	}

	pRating, errCode := commonProtocol.RateObjectWithPassword(target.DataID, target.Slot, param.RatingValue, param.AccessPassword)
	if errCode != nil {
		return nil, errCode
	}

	// * This is kinda backwards. Server returns
	// * the rating by default, so we check if
	// * the client DOESN'T want it and then just
	// * zero it out
	if !fetchRatings.Value {
		pRating = datastore_types.NewDataStoreRatingInfo()
	}

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pRating.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodRateObject
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterRateObject != nil {
		go commonProtocol.OnAfterRateObject(packet, target, param, fetchRatings)
	}

	return rmcResponse, nil
}
