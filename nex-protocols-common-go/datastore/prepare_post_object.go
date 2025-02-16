package datastore

import (
	"fmt"
	"time"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func (commonProtocol *CommonProtocol) preparePostObject(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePreparePostParam) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.InitializeObjectByPreparePostParam == nil {
		common_globals.Logger.Warning("InitializeObjectByPreparePostParam not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if commonProtocol.InitializeObjectRatingWithSlot == nil {
		common_globals.Logger.Warning("InitializeObjectRatingWithSlot not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if commonProtocol.S3Presigner == nil {
		common_globals.Logger.Warning("S3Presigner not defined")
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
		common_globals.Logger.Errorf("Error on object init: %s", errCode.Error())
		return nil, errCode
	}

	// TODO - Should this be moved to InitializeObjectByPreparePostParam?
	param.RatingInitParams.Each(func(_ int, ratingInitParamWithSlot *datastore_types.DataStoreRatingInitParamWithSlot) bool {
		errCode = commonProtocol.InitializeObjectRatingWithSlot(dataID, ratingInitParamWithSlot)
		if errCode != nil {
			common_globals.Logger.Errorf("Error on rating init: %s", errCode.Error())

			return true
		}

		return false
	})

	if errCode != nil {
		return nil, errCode
	}

	bucket := commonProtocol.S3Bucket
	key := fmt.Sprintf("%s/%d.bin", commonProtocol.s3DataKeyBase, dataID)

	URL, formData, err := commonProtocol.S3Presigner.PostObject(bucket, key, time.Minute*15)
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.OperationNotAllowed, "change_error")
	}

	requestHeaders, errCode := commonProtocol.S3PostRequestHeaders()
	if errCode != nil {
		return nil, errCode
	}

	pReqPostInfo := datastore_types.NewDataStoreReqPostInfo()

	pReqPostInfo.DataID = types.NewPrimitiveU64(dataID)
	pReqPostInfo.URL = types.NewString(URL.String())
	pReqPostInfo.RequestHeaders = types.NewList[*datastore_types.DataStoreKeyValue]()
	pReqPostInfo.FormFields = types.NewList[*datastore_types.DataStoreKeyValue]()
	pReqPostInfo.RootCACert = types.NewBuffer(commonProtocol.RootCACert)

	pReqPostInfo.RequestHeaders.Type = datastore_types.NewDataStoreKeyValue()
	pReqPostInfo.RequestHeaders.SetFromData(requestHeaders)

	pReqPostInfo.FormFields.Type = datastore_types.NewDataStoreKeyValue()

	for key, value := range formData {
		field := datastore_types.NewDataStoreKeyValue()
		field.Key = types.NewString(key)
		field.Value = types.NewString(value)

		pReqPostInfo.FormFields.Append(field)
	}

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pReqPostInfo.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodPreparePostObject
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterPreparePostObject != nil {
		go commonProtocol.OnAfterPreparePostObject(packet, param)
	}

	return rmcResponse, nil
}
