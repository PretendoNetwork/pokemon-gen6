// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking2/types"
)

func (protocol *Protocol) handleGetRankingByPrincipalID(packet nex.PacketInterface) {
	if protocol.GetRankingByPrincipalID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking2::GetRankingByPrincipalID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	getParam := ranking2_types.NewRanking2GetParam()
	principalIDList := types.NewList[*types.PID]()
	principalIDList.Type = types.NewPID(0)

	var err error

	err = getParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPrincipalID(fmt.Errorf("Failed to read getParam from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = principalIDList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRankingByPrincipalID(fmt.Errorf("Failed to read principalIDList from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRankingByPrincipalID(nil, packet, callID, getParam, principalIDList)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
