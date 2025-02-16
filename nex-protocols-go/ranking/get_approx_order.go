// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
)

func (protocol *Protocol) handleGetApproxOrder(packet nex.PacketInterface) {
	if protocol.GetApproxOrder == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::GetApproxOrder not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	category := types.NewPrimitiveU32(0)
	orderParam := ranking_types.NewRankingOrderParam()
	score := types.NewPrimitiveU32(0)
	uniqueID := types.NewPrimitiveU64(0)
	principalID := types.NewPID(0)

	var err error

	err = category.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetApproxOrder(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = orderParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetApproxOrder(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = score.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetApproxOrder(fmt.Errorf("Failed to read score from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetApproxOrder(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = principalID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetApproxOrder(fmt.Errorf("Failed to read principalID from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetApproxOrder(nil, packet, callID, category, orderParam, score, uniqueID, principalID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
