// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking2/types"
)

func (protocol *Protocol) handleGetRankingCharts(packet nex.PacketInterface) {
	if protocol.GetRankingCharts == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking2::GetRankingCharts not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	infoArray := types.NewList[*ranking2_types.Ranking2ChartInfoInput]()
	infoArray.Type = ranking2_types.NewRanking2ChartInfoInput()

	err := infoArray.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingCharts(fmt.Errorf("Failed to read infoArray from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRankingCharts(nil, packet, callID, infoArray)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
