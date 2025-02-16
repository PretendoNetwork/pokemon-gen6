// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
)

func (protocol *Protocol) handleChangeAttributes(packet nex.PacketInterface) {
	if protocol.ChangeAttributes == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::ChangeAttributes not implemented")

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
	changeParam := ranking_types.NewRankingChangeAttributesParam()
	uniqueID := types.NewPrimitiveU64(0)

	var err error

	err = category.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangeAttributes(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = changeParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangeAttributes(fmt.Errorf("Failed to read changeParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangeAttributes(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ChangeAttributes(nil, packet, callID, category, changeParam, uniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
