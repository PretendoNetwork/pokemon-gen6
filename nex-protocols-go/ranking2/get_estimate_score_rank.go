// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking2/types"
)

func (protocol *Protocol) handleGetEstimateScoreRank(packet nex.PacketInterface) {
	if protocol.GetEstimateScoreRank == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking2::GetEstimateScoreRank not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	input := ranking2_types.NewRanking2EstimateScoreRankInput()

	err := input.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetEstimateScoreRank(fmt.Errorf("Failed to read input from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetEstimateScoreRank(nil, packet, callID, input)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
