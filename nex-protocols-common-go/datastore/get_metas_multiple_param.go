package datastore

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func (commonProtocol *CommonProtocol) getMetasMultipleParam(err error, packet nex.PacketInterface, callID uint32, params *types.List[*datastore_types.DataStoreGetMetaParam]) (*nex.RMCMessage, *nex.Error) {
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

	pMetaInfo := types.NewList[*datastore_types.DataStoreMetaInfo]()
	pResults := types.NewList[*types.QResult]()

	pMetaInfo.Type = datastore_types.NewDataStoreMetaInfo()
	pResults.Type = types.NewQResult(0)

	params.Each(func(_ int, param *datastore_types.DataStoreGetMetaParam) bool {
		var objectInfo *datastore_types.DataStoreMetaInfo
		var errCode *nex.Error

		// * Real server ignores PersistenceTarget if DataID is set
		if param.DataID.Value == 0 {
			objectInfo, errCode = commonProtocol.GetObjectInfoByPersistenceTargetWithPassword(param.PersistenceTarget, param.AccessPassword)
		} else {
			objectInfo, errCode = commonProtocol.GetObjectInfoByDataIDWithPassword(param.DataID, param.AccessPassword)
		}

		if errCode != nil {
			objectInfo = datastore_types.NewDataStoreMetaInfo()

			pResults.Append(types.NewQResultError(errCode.ResultCode))
		} else {
			errCode = commonProtocol.VerifyObjectPermission(objectInfo.OwnerID, connection.PID(), objectInfo.Permission)
			if errCode != nil {
				objectInfo = datastore_types.NewDataStoreMetaInfo()

				pResults.Append(types.NewQResultError(errCode.ResultCode))
			} else {
				pResults.Append(types.NewQResultSuccess(nex.ResultCodes.DataStore.Unknown))
			}

			objectInfo.FilterPropertiesByResultOption(param.ResultOption)
		}

		pMetaInfo.Append(objectInfo)

		return false
	})

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pMetaInfo.WriteTo(rmcResponseStream)
	pResults.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = datastore.ProtocolID
	rmcResponse.MethodID = datastore.MethodGetMetasMultipleParam
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterGetMetasMultipleParam != nil {
		go commonProtocol.OnAfterGetMetasMultipleParam(packet, params)
	}

	return rmcResponse, nil
}
