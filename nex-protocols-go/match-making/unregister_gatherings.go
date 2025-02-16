// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUnregisterGatherings(packet nex.PacketInterface) {
	if protocol.UnregisterGatherings == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::UnregisterGatherings not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	lstGatherings := types.NewList[*types.PrimitiveU32]()
	lstGatherings.Type = types.NewPrimitiveU32(0)

	err := lstGatherings.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UnregisterGatherings(fmt.Errorf("Failed to read lstGatherings from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UnregisterGatherings(nil, packet, callID, lstGatherings)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
