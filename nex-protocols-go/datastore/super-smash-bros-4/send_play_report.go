// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSendPlayReport(packet nex.PacketInterface) {
	if protocol.SendPlayReport == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperSmashBros4::SendPlayReport not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	playReport := types.NewList[*types.PrimitiveS32]()
	playReport.Type = types.NewPrimitiveS32(0)

	err := playReport.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SendPlayReport(fmt.Errorf("Failed to read playReport from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SendPlayReport(nil, packet, callID, playReport)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
