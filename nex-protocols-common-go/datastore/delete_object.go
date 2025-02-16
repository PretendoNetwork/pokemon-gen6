package datastore

import (
	"github.com/PretendoNetwork/nex-go/v2"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func (commonProtocol *CommonProtocol) deleteObject(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreDeleteParam) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.GetObjectInfoByDataID == nil {
		common_globals.Logger.Warning("GetObjectInfoByDataID not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if commonProtocol.DeleteObjectByDataIDWithPassword == nil {
		common_globals.Logger.Warning("DeleteObjectByDataIDWithPassword not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.Unknown, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	metaInfo, errCode := commonProtocol.GetObjectInfoByDataID(param.DataID)
	if errCode != nil {
		return nil, errCode
	}

	errCode = commonProtocol.VerifyObjectPermission(metaInfo.OwnerID, connection.PID(), metaInfo.DelPermission)
	if errCode != nil {
		return nil, errCode
	}

	errCode = commonProtocol.DeleteObjectByDataIDWithPassword(param.DataID, param.UpdatePassword)
	if errCode != nil {
		return nil, errCode
	}

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodDeleteObject
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterDeleteObject != nil {
		go commonProtocol.OnAfterDeleteObject(packet, param)
	}

	return rmcResponse, nil
}
