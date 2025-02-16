// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
)

func (protocol *Protocol) handleGetCachedTopXRankings(packet nex.PacketInterface) {
	if protocol.GetCachedTopXRankings == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::GetCachedTopXRankings not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	categories := types.NewList[*types.PrimitiveU32]()
	categories.Type = types.NewPrimitiveU32(0)
	orderParams := types.NewList[*ranking_types.RankingOrderParam]()
	orderParams.Type = ranking_types.NewRankingOrderParam()

	var err error

	err = categories.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCachedTopXRankings(fmt.Errorf("Failed to read categories from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = orderParams.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCachedTopXRankings(fmt.Errorf("Failed to read orderParams from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetCachedTopXRankings(nil, packet, callID, categories, orderParams)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
