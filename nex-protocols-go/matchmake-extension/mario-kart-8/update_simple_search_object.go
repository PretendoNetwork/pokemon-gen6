// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	matchmake_extension_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/mario-kart-8/types"
)

func (protocol *Protocol) handleUpdateSimpleSearchObject(packet nex.PacketInterface) {
	if protocol.UpdateSimpleSearchObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtensionMarioKart8::UpdateSimpleSearchObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	objectID := types.NewPrimitiveU32(0)
	newObject := matchmake_extension_mario_kart8_types.NewSimpleSearchObject()

	var err error

	err = objectID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateSimpleSearchObject(fmt.Errorf("Failed to read objectID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = newObject.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateSimpleSearchObject(fmt.Errorf("Failed to read newObject from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateSimpleSearchObject(nil, packet, callID, objectID, newObject)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
