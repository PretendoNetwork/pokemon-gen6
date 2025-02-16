package datastore

import (
	"github.com/PretendoNetwork/nex-go/v2"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func (commonProtocol *CommonProtocol) getMeta(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreGetMetaParam) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.GetObjectInfoByPersistenceTargetWithPassword == nil {
		common_globals.Logger.Warning("GetObjectInfoByPersistenceTargetWithPassword not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if commonProtocol.GetObjectInfoByDataIDWithPassword == nil {
		common_globals.Logger.Warning("GetObjectInfoByDataIDWithPassword not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.Unknown, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	var pMetaInfo *datastore_types.DataStoreMetaInfo
	var errCode *nex.Error

	// * Real server ignores PersistenceTarget if DataID is set
	if param.DataID.Value == 0 {
		pMetaInfo, errCode = commonProtocol.GetObjectInfoByPersistenceTargetWithPassword(param.PersistenceTarget, param.AccessPassword)
	} else {
		pMetaInfo, errCode = commonProtocol.GetObjectInfoByDataIDWithPassword(param.DataID, param.AccessPassword)
	}

	if errCode != nil {
		return nil, errCode
	}

	errCode = commonProtocol.VerifyObjectPermission(pMetaInfo.OwnerID, connection.PID(), pMetaInfo.Permission)
	if errCode != nil {
		return nil, errCode
	}

	pMetaInfo.FilterPropertiesByResultOption(param.ResultOption)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pMetaInfo.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodGetMeta
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterGetMeta != nil {
		go commonProtocol.OnAfterGetMeta(packet, param)
	}

	return rmcResponse, nil
}
