package ranking

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
)

func (commonProtocol *CommonProtocol) uploadCommonData(err error, packet nex.PacketInterface, callID uint32, commonData *types.Buffer, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.UploadCommonData == nil {
		common_globals.Logger.Warning("Ranking::UploadCommonData missing UploadCommonData!")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Ranking.InvalidArgument, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	err = commonProtocol.UploadCommonData(connection.PID(), uniqueID, commonData)
	if err != nil {
		common_globals.Logger.Critical(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Ranking.Unknown, "change_error")
	}

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = ranking.ProtocolID
	rmcResponse.MethodID = ranking.MethodUploadCommonData
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterUploadCommonData != nil {
		go commonProtocol.OnAfterUploadCommonData(packet, commonData, uniqueID)
	}

	return rmcResponse, nil
}
