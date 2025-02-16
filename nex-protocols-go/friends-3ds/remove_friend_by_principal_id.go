// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRemoveFriendByPrincipalID(packet nex.PacketInterface) {
	if protocol.RemoveFriendByPrincipalID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends3DS::RemoveFriendByPrincipalID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pid := types.NewPID(0)

	err := pid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RemoveFriendByPrincipalID(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RemoveFriendByPrincipalID(nil, packet, callID, pid)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
