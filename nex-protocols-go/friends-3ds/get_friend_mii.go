// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/v2/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetFriendMii(packet nex.PacketInterface) {
	if protocol.GetFriendMii == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends3DS::GetFriendMii not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	friends := types.NewList[*friends_3ds_types.FriendInfo]()
	friends.Type = friends_3ds_types.NewFriendInfo()

	err := friends.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetFriendMii(fmt.Errorf("Failed to read friends from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetFriendMii(nil, packet, callID, friends)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
