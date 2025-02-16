// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetTournamentResult(packet nex.PacketInterface) {
	if protocol.GetTournamentResult == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtensionSuperSmashBros4::GetTournamentResult not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::GetTournamentResult STUBBED")

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.GetTournamentResult(nil, packet, callID, packet.Payload())
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
