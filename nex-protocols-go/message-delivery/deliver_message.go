// Package protocol implements the Message Deliver protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleDeliverMessage(packet nex.PacketInterface) {
	if protocol.DeliverMessage == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MessageDelivery::DeliverMessage not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	oUserMessage := types.NewAnyDataHolder()

	err := oUserMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeliverMessage(fmt.Errorf("Failed to read oUserMessage from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeliverMessage(nil, packet, callID, oUserMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
