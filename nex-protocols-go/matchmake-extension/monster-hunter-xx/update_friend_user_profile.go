// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	matchmake_extension_monster_hunter_xx_types "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/monster-hunter-xx/types"
)

func (protocol *Protocol) handleUpdateFriendUserProfile(packet nex.PacketInterface) {
	if protocol.UpdateFriendUserProfile == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtensionMonsterHunterXX::UpdateFriendUserProfile not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	param := matchmake_extension_monster_hunter_xx_types.NewFriendUserParam()

	err := param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateFriendUserProfile(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateFriendUserProfile(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
