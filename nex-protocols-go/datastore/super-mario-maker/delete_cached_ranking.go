// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleDeleteCachedRanking(packet nex.PacketInterface) {
	if protocol.DeleteCachedRanking == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::DeleteCachedRanking not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	rankingType := types.NewString("")
	rankingArgs := types.NewList[*types.String]()
	rankingArgs.Type = types.NewString("")

	var err error

	err = rankingType.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteCachedRanking(fmt.Errorf("Failed to read rankingType from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = rankingArgs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteCachedRanking(fmt.Errorf("Failed to read rankingArgs from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeleteCachedRanking(nil, packet, callID, rankingType, rankingArgs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
