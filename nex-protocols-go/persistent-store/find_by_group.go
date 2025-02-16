// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleFindByGroup(packet nex.PacketInterface) {
	if protocol.FindByGroup == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "PersistentStore::FindByGroup not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	uiGroup := types.NewPrimitiveU32(0)

	err := uiGroup.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByGroup(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.FindByGroup(nil, packet, callID, uiGroup)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
