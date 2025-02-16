// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
)

func (protocol *Protocol) handleGetRankingByUniqueIDList(packet nex.PacketInterface) {
	if protocol.GetRankingByUniqueIDList == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::GetRankingByUniqueIDList not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	nexUniqueIDList := types.NewList[*types.PrimitiveU64]()
	nexUniqueIDList.Type = types.NewPrimitiveU64(0)
	rankingMode := types.NewPrimitiveU8(0)
	category := types.NewPrimitiveU32(0)
	orderParam := ranking_types.NewRankingOrderParam()
	uniqueID := types.NewPrimitiveU64(0)

	var err error

	err = nexUniqueIDList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByUniqueIDList(fmt.Errorf("Failed to read nexUniqueIDList from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = rankingMode.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByUniqueIDList(fmt.Errorf("Failed to read rankingMode from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = category.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByUniqueIDList(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = orderParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByUniqueIDList(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByUniqueIDList(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRankingByUniqueIDList(nil, packet, callID, nexUniqueIDList, rankingMode, category, orderParam, uniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
