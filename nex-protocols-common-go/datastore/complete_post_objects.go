package datastore

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
)

func (commonProtocol *CommonProtocol) completePostObjects(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.minIOClient == nil {
		common_globals.Logger.Warning("MinIOClient not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if commonProtocol.GetObjectSizeByDataID == nil {
		common_globals.Logger.Warning("GetObjectSizeByDataID not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if commonProtocol.UpdateObjectUploadCompletedByDataID == nil {
		common_globals.Logger.Warning("UpdateObjectUploadCompletedByDataID not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.Unknown, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	var errorCode *nex.Error

	dataIDs.Each(func(_ int, dataID *types.PrimitiveU64) bool {
		bucket := commonProtocol.S3Bucket
		key := fmt.Sprintf("%s/%d.bin", commonProtocol.s3DataKeyBase, dataID)

		objectSizeS3, err := commonProtocol.S3ObjectSize(bucket, key)
		if err != nil {
			common_globals.Logger.Error(err.Error())
			errorCode = nex.NewError(nex.ResultCodes.DataStore.NotFound, "change_error")

			return true
		}

		objectSizeDB, errCode := commonProtocol.GetObjectSizeByDataID(dataID)
		if errCode != nil {
			errorCode = errCode

			return true
		}

		if objectSizeS3 != uint64(objectSizeDB) {
			common_globals.Logger.Errorf("Object with DataID %d did not upload correctly! Mismatched sizes", dataID)
			// TODO - Is this a good error?
			errorCode = nex.NewError(nex.ResultCodes.DataStore.Unknown, "change_error")

			return true
		}

		errCode = commonProtocol.UpdateObjectUploadCompletedByDataID(dataID, true)
		if errCode != nil {
			errorCode = errCode

			return true
		}

		return false
	})

	if errorCode != nil {
		return nil, errorCode
	}

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodCompletePostObjects
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterCompletePostObjects != nil {
		go commonProtocol.OnAfterCompletePostObjects(packet, dataIDs)
	}

	return rmcResponse, nil
}
