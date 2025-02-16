// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleDeleteFromDeletions(packet nex.PacketInterface) {
	if protocol.DeleteFromDeletions == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMakingExt::DeleteFromDeletions not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	lstDeletions := types.NewList[*types.PrimitiveU32]()
	lstDeletions.Type = types.NewPrimitiveU32(0)
	pid := types.NewPID(0)

	var err error

	err = lstDeletions.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteFromDeletions(fmt.Errorf("Failed to read lstDeletions from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = pid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteFromDeletions(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeleteFromDeletions(nil, packet, callID, lstDeletions, pid)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
