// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking2/types"
)

func (protocol *Protocol) handlePutScore(packet nex.PacketInterface) {
	if protocol.PutScore == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking2::PutScore not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	scoreDataList := types.NewList[*ranking2_types.Ranking2ScoreData]()
	scoreDataList.Type = ranking2_types.NewRanking2ScoreData()
	nexUniqueID := types.NewPrimitiveU64(0)

	var err error

	err = scoreDataList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PutScore(fmt.Errorf("Failed to read scoreDataList from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = nexUniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PutScore(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PutScore(nil, packet, callID, scoreDataList, nexUniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
