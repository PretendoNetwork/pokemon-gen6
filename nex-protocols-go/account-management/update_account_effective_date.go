// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUpdateAccountEffectiveDate(packet nex.PacketInterface) {
	if protocol.UpdateAccountEffectiveDate == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::UpdateAccountEffectiveDate not implemented")

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
	dtEffectiveFrom := types.NewDateTime(0)
	strNotEffectiveMessage := types.NewString("")

	var err error

	err = idPrincipal.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAccountEffectiveDate(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = dtEffectiveFrom.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAccountEffectiveDate(fmt.Errorf("Failed to read dtEffectiveFrom from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strNotEffectiveMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAccountEffectiveDate(fmt.Errorf("Failed to read strNotEffectiveMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateAccountEffectiveDate(nil, packet, callID, idPrincipal, dtEffectiveFrom, strNotEffectiveMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
