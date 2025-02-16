package ranking

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
)

func (commonProtocol *CommonProtocol) getCommonData(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.GetCommonData == nil {
		common_globals.Logger.Warning("Ranking::GetCommonData missing GetCommonData!")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Ranking.InvalidArgument, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	commonData, err := commonProtocol.GetCommonData(uniqueID)
	if err != nil {
		return nil, nex.NewError(nex.ResultCodes.Ranking.NotFound, "change_error")
	}

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	commonData.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = ranking.ProtocolID
	rmcResponse.MethodID = ranking.MethodGetCommonData
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterGetCommonData != nil {
		go commonProtocol.OnAfterGetCommonData(packet, uniqueID)
	}

	return rmcResponse, nil
}
