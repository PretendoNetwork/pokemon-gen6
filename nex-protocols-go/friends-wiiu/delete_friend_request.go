// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleDeleteFriendRequest(packet nex.PacketInterface) {
	if protocol.DeleteFriendRequest == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "FriendsWiiU::DeleteFriendRequest not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	id := types.NewPrimitiveU64(0)

	err := id.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteFriendRequest(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeleteFriendRequest(nil, packet, callID, id)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
