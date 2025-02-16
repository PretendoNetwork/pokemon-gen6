// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSetExcludeJoinedMatchmakeSession(packet nex.PacketInterface) {
	if protocol.SetExcludeJoinedMatchmakeSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Debug::SetExcludeJoinedMatchmakeSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	globals.Logger.Warning("Debug::SetExcludeJoinedMatchmakeSession STUBBED")

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.SetExcludeJoinedMatchmakeSession(nil, packet, callID, packet.Payload())
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
