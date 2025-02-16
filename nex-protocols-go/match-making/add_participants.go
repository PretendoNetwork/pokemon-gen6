// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleAddParticipants(packet nex.PacketInterface) {
	if protocol.AddParticipants == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::AddParticipants not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	idGathering := types.NewPrimitiveU32(0)
	lstPrincipals := types.NewList[*types.PID]()
	lstPrincipals.Type = types.NewPID(0)
	strMessage := types.NewString("")

	var err error

	err = idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddParticipants(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = lstPrincipals.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddParticipants(fmt.Errorf("Failed to read lstPrincipals from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddParticipants(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AddParticipants(nil, packet, callID, idGathering, lstPrincipals, strMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
