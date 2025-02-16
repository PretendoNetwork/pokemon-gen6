// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetFriendUserStatuses(packet nex.PacketInterface) {
	if protocol.GetFriendUserStatuses == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Subscriber::GetFriendUserStatuses not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	unknown := types.NewList[*types.PrimitiveU8]()
	unknown.Type = types.NewPrimitiveU8(0)

	err := unknown.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetFriendUserStatuses(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetFriendUserStatuses(nil, packet, callID, unknown)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
