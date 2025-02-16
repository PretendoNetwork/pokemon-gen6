// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUpdatePreference(packet nex.PacketInterface) {
	if protocol.UpdatePreference == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends3DS::UpdatePreference not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	publicMode := types.NewPrimitiveBool(false)
	showGame := types.NewPrimitiveBool(false)
	showPlayedGame := types.NewPrimitiveBool(false)

	var err error

	err = publicMode.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdatePreference(fmt.Errorf("Failed to read publicMode from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = showGame.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdatePreference(fmt.Errorf("Failed to read showGame from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = showPlayedGame.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdatePreference(fmt.Errorf("Failed to read showPlayedGame from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdatePreference(nil, packet, callID, publicMode, showGame, showPlayedGame)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
