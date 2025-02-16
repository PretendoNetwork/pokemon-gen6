// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
)

func (protocol *Protocol) handleAutoMatchmakeWithSearchCriteriaPostpone(packet nex.PacketInterface) {
	if protocol.AutoMatchmakeWithSearchCriteriaPostpone == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::AutoMatchmakeWithSearchCriteriaPostpone not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	lstSearchCriteria := types.NewList[*match_making_types.MatchmakeSessionSearchCriteria]()
	lstSearchCriteria.Type = match_making_types.NewMatchmakeSessionSearchCriteria()
	anyGathering := types.NewAnyDataHolder()
	strMessage := types.NewString("")

	var err error

	err = lstSearchCriteria.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AutoMatchmakeWithSearchCriteriaPostpone(fmt.Errorf("Failed to read lstSearchCriteria from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = anyGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AutoMatchmakeWithSearchCriteriaPostpone(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AutoMatchmakeWithSearchCriteriaPostpone(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AutoMatchmakeWithSearchCriteriaPostpone(nil, packet, callID, lstSearchCriteria, anyGathering, strMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
