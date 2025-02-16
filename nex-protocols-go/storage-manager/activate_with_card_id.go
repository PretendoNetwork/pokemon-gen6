// Package protocol implements the StorageManager protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleActivateWithCardID(packet nex.PacketInterface) {
	if protocol.ActivateWithCardID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "StorageManager::ActivateWithCardID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	unknown := types.NewPrimitiveU8(0)
	cardID := types.NewPrimitiveU64(0)

	var err error

	err = unknown.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ActivateWithCardID(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = cardID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ActivateWithCardID(fmt.Errorf("Failed to read cardID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ActivateWithCardID(nil, packet, callID, unknown, cardID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
