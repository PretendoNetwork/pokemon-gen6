// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleDisableAccount(packet nex.PacketInterface) {
	if protocol.DisableAccount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::DisableAccount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	idPrincipal := types.NewPID(0)
	dtUntil := types.NewDateTime(0)
	strMessage := types.NewString("")

	var err error

	err = idPrincipal.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DisableAccount(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = dtUntil.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DisableAccount(fmt.Errorf("Failed to read dtUntil from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DisableAccount(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DisableAccount(nil, packet, callID, idPrincipal, dtUntil, strMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
