package datastore

import (
	"github.com/PretendoNetwork/nex-go/v2"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func (commonProtocol *CommonProtocol) postMetaBinary(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error) {
	// * This method looks to function identically to DataStore::PreparePostObject,
	// * except the only difference being it doesn't return an S3 upload URL. This
	// * needs to be verified though, as there are other methods in the family such
	// * as DataStore::PostMetaBinaryWithDataID which make less sense in this context,
	// * unless those are just used to *update* a meta binary? Or maybe the DataID in
	// * those methods is a pre-allocated DataID from the server? Needs more testing

	if commonProtocol.InitializeObjectByPreparePostParam == nil {
		common_globals.Logger.Warning("InitializeObjectByPreparePostParam not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if commonProtocol.InitializeObjectRatingWithSlot == nil {
		common_globals.Logger.Warning("InitializeObjectRatingWithSlot not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.Unknown, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	// TODO - Need to verify what param.PersistenceInitParam.DeleteLastObject really means. It's often set to true even when it wouldn't make sense
	dataID, errCode := commonProtocol.InitializeObjectByPreparePostParam(connection.PID(), param)
	if errCode != nil {
		common_globals.Logger.Errorf("Error code on object init: %s", errCode.Error())
		return nil, errCode
	}

	// TODO - Should this be moved to InitializeObjectByPreparePostParam?
	param.RatingInitParams.Each(func(_ int, ratingInitParamWithSlot *datastore_types.DataStoreRatingInitParamWithSlot) bool {
		errCode = commonProtocol.InitializeObjectRatingWithSlot(dataID, ratingInitParamWithSlot)
		if errCode != nil {
			common_globals.Logger.Errorf("Error code on rating init: %s", errCode.Error())
			return true
		}

		return false
	})

	if errCode != nil {
		return nil, errCode
	}

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	rmcResponseStream.WritePrimitiveUInt64LE(dataID)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodPostMetaBinary
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterPostMetaBinary != nil {
		go commonProtocol.OnAfterPostMetaBinary(packet, param)
	}

	return rmcResponse, nil
}
