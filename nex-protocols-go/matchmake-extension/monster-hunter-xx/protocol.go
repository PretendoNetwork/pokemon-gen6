// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	matchmake_extension_monster_hunter_x_x_types "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/monster-hunter-xx/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Matchmake Extension (Monster Hunter XX) protocol
	ProtocolID = 0x6D

	// MethodUpdateFriendUserProfile is the method ID for the UpdateFriendUserProfile method
	MethodUpdateFriendUserProfile = 0x36

	// MethodGetFriendUserProfiles is the method ID for the GetFriendUserProfiles method
	MethodGetFriendUserProfiles = 0x37

	// MethodAddFriends is the method ID for the AddFriends method
	MethodAddFriends = 0x39

	// MethodRemoveFriend is the method ID for the RemoveFriend method
	MethodRemoveFriend = 0x3A

	// MethodFindCommunityByOwner is the method ID for the FindCommunityByOwner method
	MethodFindCommunityByOwner = 0x3B
)

var patchedMethods = []uint32{
	MethodUpdateFriendUserProfile,
	MethodGetFriendUserProfiles,
	MethodAddFriends,
	MethodRemoveFriend,
	MethodFindCommunityByOwner,
}

type matchmakeExtensionProtocol = matchmake_extension.Protocol

// Protocol stores all the RMC method handlers for the Matchmake Extension (Monster Hunter XX) protocol and listens for requests
// Embeds the Matchmake Extension protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	matchmakeExtensionProtocol
	UpdateFriendUserProfile func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_extension_monster_hunter_x_x_types.FriendUserParam) (*nex.RMCMessage, *nex.Error)
	GetFriendUserProfiles   func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)
	AddFriends              func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)
	RemoveFriend            func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, *nex.Error)
	FindCommunityByOwner    func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.matchmakeExtensionProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodUpdateFriendUserProfile:
		protocol.handleUpdateFriendUserProfile(packet)
	case MethodGetFriendUserProfiles:
		protocol.handleGetFriendUserProfiles(packet)
	case MethodAddFriends:
		protocol.handleAddFriends(packet)
	case MethodRemoveFriend:
		protocol.handleRemoveFriend(packet)
	case MethodFindCommunityByOwner:
		protocol.handleFindCommunityByOwner(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Matchmake Extension (Monster Hunter XX) method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new MatchmakeExtensionMonsterHunterXX protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.matchmakeExtensionProtocol.SetEndpoint(endpoint)

	return protocol
}
