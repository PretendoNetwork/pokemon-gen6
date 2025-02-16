package datastore

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

func (commonProtocol *CommonProtocol) getMetas(err error, packet nex.PacketInterface, callID uint32, dataIDs *types.List[*types.PrimitiveU64], param *datastore_types.DataStoreGetMetaParam) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.GetObjectInfoByDataID == nil {
		common_globals.Logger.Warning("GetObjectInfoByDataID not defined")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.Unknown, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	// TODO - Verify if param.PersistenceTarget is respected? It wouldn't make sense here but who knows

	pMetaInfo := types.NewList[*datastore_types.DataStoreMetaInfo]()
	pResults := types.NewList[*types.QResult]()

	pMetaInfo.Type = datastore_types.NewDataStoreMetaInfo()
	pResults.Type = types.NewQResult(0)

	// * param has an AccessPassword, but it goes unchecked here.
	// * The password would need to be the same for every object
	// * in the input array, which doesn't make any sense. Assuming
	// * it's unused until proven otherwise

	dataIDs.Each(func(_ int, dataID *types.PrimitiveU64) bool {
		objectInfo, errCode := commonProtocol.GetObjectInfoByDataID(dataID)

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
	rmcResponse.MethodID = datastore.MethodGetMetas
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterGetMetas != nil {
		go commonProtocol.OnAfterGetMetas(packet, dataIDs, param)
	}

	return rmcResponse, nil
}
