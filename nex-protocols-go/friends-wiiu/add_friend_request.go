// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/v2/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleAddFriendRequest(packet nex.PacketInterface) {
	if protocol.AddFriendRequest == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "FriendsWiiU::AddFriendRequest not implemented")

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
	unknown2 := types.NewPrimitiveU8(0)
	message := types.NewString("")
	unknown4 := types.NewPrimitiveU8(0)
	unknown5 := types.NewString("")
	gameKey := friends_wiiu_types.NewGameKey()
	unknown6 := types.NewDateTime(0)

	var err error

	err = pid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddFriendRequest(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddFriendRequest(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = message.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddFriendRequest(fmt.Errorf("Failed to read message from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown4.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddFriendRequest(fmt.Errorf("Failed to read unknown4 from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown5.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddFriendRequest(fmt.Errorf("Failed to read unknown5 from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = gameKey.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddFriendRequest(fmt.Errorf("Failed to read gameKey from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown6.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddFriendRequest(fmt.Errorf("Failed to read unknown6 from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AddFriendRequest(nil, packet, callID, pid, unknown2, message, unknown4, unknown5, gameKey, unknown6)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
