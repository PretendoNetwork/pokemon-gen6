// Package protocol implements the Nintendo Badge Arcade Secure Connection protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetMaintenanceStatus(packet nex.PacketInterface) {
	if protocol.GetMaintenanceStatus == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SecureConnectionNintendoBadgeArcade::GetMaintenanceStatus not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.GetMaintenanceStatus(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
