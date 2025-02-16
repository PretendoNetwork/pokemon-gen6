// Package protocol implements the Subscription protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleReplaceTargetAndGetSubscriptionData(packet nex.PacketInterface) {
	if protocol.ReplaceTargetAndGetSubscriptionData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SubscriptionProtocol::ReplaceTargetAndGetSubscriptionData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.ReplaceTargetAndGetSubscriptionData(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
