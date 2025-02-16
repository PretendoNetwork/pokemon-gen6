// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUnfollowAllAndFollow(packet nex.PacketInterface) {
	if protocol.UnfollowAllAndFollow == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Subscriber::UnfollowAllAndFollow not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	globals.Logger.Warning("Subscriber::UnfollowAllAndFollow STUBBED")

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.UnfollowAllAndFollow(nil, packet, callID, packet.Payload())
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
