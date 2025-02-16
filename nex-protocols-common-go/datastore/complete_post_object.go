package datastore

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func (commonProtocol *CommonProtocol) completePostObject(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParam) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.minIOClient == nil {
		common_globals.Logger.Warning("MinIOClient not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if commonProtocol.GetObjectInfoByDataID == nil {
		common_globals.Logger.Warning("GetObjectInfoByDataID not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if commonProtocol.GetObjectOwnerByDataID == nil {
		common_globals.Logger.Warning("GetObjectOwnerByDataID not defined")
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

	if commonProtocol.DeleteObjectByDataID == nil {
		common_globals.Logger.Warning("DeleteObjectByDataID not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.Unknown, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	// * If GetObjectInfoByDataID returns data then that means
	// * the object has already been marked as uploaded. So do
	// * nothing
	objectInfo, _ := commonProtocol.GetObjectInfoByDataID(param.DataID)
	if objectInfo != nil {
		return nil, nex.NewError(nex.ResultCodes.DataStore.PermissionDenied, "change_error")
	}

	// * Only allow an objects owner to make this request
	ownerPID, errCode := commonProtocol.GetObjectOwnerByDataID(param.DataID)
	if errCode != nil {
		return nil, errCode
	}

	if ownerPID != connection.PID().LegacyValue() {
		return nil, nex.NewError(nex.ResultCodes.DataStore.PermissionDenied, "change_error")
	}

	bucket := commonProtocol.S3Bucket
	key := fmt.Sprintf("%s/%d.bin", commonProtocol.s3DataKeyBase, param.DataID)

	if param.IsSuccess.Value {
		objectSizeS3, err := commonProtocol.S3ObjectSize(bucket, key)
		if err != nil {
			common_globals.Logger.Error(err.Error())
			return nil, nex.NewError(nex.ResultCodes.DataStore.NotFound, "change_error")
		}

		objectSizeDB, errCode := commonProtocol.GetObjectSizeByDataID(param.DataID)
		if errCode != nil {
			return nil, errCode
		}

		if objectSizeS3 != uint64(objectSizeDB) {
			common_globals.Logger.Errorf("Object with DataID %d did not upload correctly! Mismatched sizes", param.DataID)
			// TODO - Is this a good error?
			return nil, nex.NewError(nex.ResultCodes.DataStore.Unknown, "change_error")
		}

		errCode = commonProtocol.UpdateObjectUploadCompletedByDataID(param.DataID, true)
		if errCode != nil {
			return nil, errCode
		}
	} else {
		errCode := commonProtocol.DeleteObjectByDataID(param.DataID)
		if errCode != nil {
			return nil, errCode
		}
	}

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodCompletePostObject
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterCompletePostObject != nil {
		go commonProtocol.OnAfterCompletePostObject(packet, param)
	}

	return rmcResponse, nil
}
