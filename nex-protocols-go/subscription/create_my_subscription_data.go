// Package protocol implements the Subscription protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleCreateMySubscriptionData(packet nex.PacketInterface) {
	if protocol.CreateMySubscriptionData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SubscriptionProtocol::CreateMySubscriptionData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	unk := types.NewPrimitiveU64(0)

	err := unk.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateMySubscriptionData(fmt.Errorf("Failed to read unk from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	// * This is done since the server doesn't need to care about the data here (it's game-specific),
	// * so we just pass it along to store however the handler wants
	// TODO - Is this really the best way to do this?
	content := parametersStream.ReadRemaining()

	rmcMessage, rmcError := protocol.CreateMySubscriptionData(nil, packet, callID, unk, content)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
