package datastore

import (
	"fmt"
	"time"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func (commonProtocol *CommonProtocol) prepareGetObject(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareGetParam) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.GetObjectInfoByDataID == nil {
		common_globals.Logger.Warning("GetObjectInfoByDataID not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if commonProtocol.S3Presigner == nil {
		common_globals.Logger.Warning("S3Presigner not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.Unknown, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	bucket := commonProtocol.S3Bucket
	key := fmt.Sprintf("%s/%d.bin", commonProtocol.s3DataKeyBase, param.DataID)

	objectInfo, errCode := commonProtocol.GetObjectInfoByDataID(param.DataID)
	if errCode != nil {
		return nil, errCode
	}

	errCode = commonProtocol.VerifyObjectPermission(objectInfo.OwnerID, connection.PID(), objectInfo.Permission)
	if errCode != nil {
		return nil, errCode
	}

	url, err := commonProtocol.S3Presigner.GetObject(bucket, key, time.Minute*15)
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.OperationNotAllowed, "change_error")
	}

	requestHeaders, errCode := commonProtocol.S3GetRequestHeaders()
	if errCode != nil {
		return nil, errCode
	}

	pReqGetInfo := datastore_types.NewDataStoreReqGetInfo()

	pReqGetInfo.URL = types.NewString(url.String())
	pReqGetInfo.RequestHeaders = types.NewList[*datastore_types.DataStoreKeyValue]()
	pReqGetInfo.Size = objectInfo.Size.Copy().(*types.PrimitiveU32)
	pReqGetInfo.RootCACert = types.NewBuffer(commonProtocol.RootCACert)
	pReqGetInfo.DataID = param.DataID

	pReqGetInfo.RequestHeaders.Type = datastore_types.NewDataStoreKeyValue()
	pReqGetInfo.RequestHeaders.SetFromData(requestHeaders)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pReqGetInfo.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodPrepareGetObject
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterPrepareGetObject != nil {
		go commonProtocol.OnAfterPrepareGetObject(packet, param)
	}

	return rmcResponse, nil
}
