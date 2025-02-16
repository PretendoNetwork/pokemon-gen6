// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleFindByNameRegex(packet nex.PacketInterface) {
	if protocol.FindByNameRegex == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::FindByNameRegex not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	uiGroups := types.NewPrimitiveU32(0)
	strRegex := types.NewString("")
	resultRange := types.NewResultRange()

	var err error

	err = uiGroups.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByNameRegex(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strRegex.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByNameRegex(fmt.Errorf("Failed to read strRegex from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByNameRegex(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.FindByNameRegex(nil, packet, callID, uiGroups, strRegex, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
