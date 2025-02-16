package ticket_granting

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	ticket_granting "github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting"
)

func (commonProtocol *CommonProtocol) requestTicket(err error, packet nex.PacketInterface, callID uint32, idSource *types.PID, idTarget *types.PID) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint().(*nex.PRUDPEndPoint)

	sourceAccount, errorCode := endpoint.AccountDetailsByPID(idSource)
	if errorCode != nil && errorCode.ResultCode != nex.ResultCodes.Core.AccessDenied {
		return nil, errorCode
	}

	targetAccount, errorCode := endpoint.AccountDetailsByPID(idTarget)
	if errorCode != nil && errorCode.ResultCode != nex.ResultCodes.Core.AccessDenied {
		return nil, errorCode
	}

	encryptedTicket, errorCode := generateTicket(sourceAccount, targetAccount, commonProtocol.SessionKeyLength, endpoint)
	if errorCode != nil && errorCode.ResultCode != nex.ResultCodes.Core.AccessDenied {
		return nil, errorCode
	}

	// * From the wiki:
	// *
	// * "If the source or target pid is invalid, the %retval% field is set to Core::AccessDenied and the ticket is empty."
	retval := types.NewQResultSuccess(nex.ResultCodes.Core.Unknown)
	bufResponse := types.NewBuffer(encryptedTicket)

	if errorCode != nil {
		retval = types.NewQResultError(nex.ResultCodes.Core.AccessDenied)
		bufResponse = types.NewBuffer([]byte{})
	}

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	retval.WriteTo(rmcResponseStream)
	bufResponse.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = ticket_granting.ProtocolID
	rmcResponse.MethodID = ticket_granting.MethodRequestTicket
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterRequestTicket != nil {
		go commonProtocol.OnAfterRequestTicket(packet, idSource, idTarget)
	}

	return rmcResponse, nil
}
