package secureconnection

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	secure_connection "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"

	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
)

func (commonProtocol *CommonProtocol) sendReport(err error, packet nex.PacketInterface, callID uint32, reportID *types.PrimitiveU32, reportData *types.QBuffer) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.CreateReportDBRecord == nil {
		common_globals.Logger.Warning("SecureConnection::SendReport missing CreateReportDBRecord!")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Critical(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.Unknown, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	err = commonProtocol.CreateReportDBRecord(connection.PID(), reportID, reportData)
	if err != nil {
		common_globals.Logger.Critical(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.Unknown, "change_error")
	}

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = secure_connection.ProtocolID
	rmcResponse.MethodID = secure_connection.MethodSendReport
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterSendReport != nil {
		go commonProtocol.OnAfterSendReport(packet, reportID, reportData)
	}

	return rmcResponse, nil
}
