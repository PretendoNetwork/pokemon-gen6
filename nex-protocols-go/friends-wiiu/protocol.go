// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/v2/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Friends (WiiU) protocol
	ProtocolID = 0x66

	// MethodUpdateAndGetAllInformation is the method ID for method UpdateAndGetAllInformation
	MethodUpdateAndGetAllInformation = 0x1

	// MethodAddFriend is the method ID for method AddFriend
	MethodAddFriend = 0x2

	// MethodAddFriendByName is the method ID for method AddFriendByName
	MethodAddFriendByName = 0x3

	// MethodRemoveFriend is the method ID for method RemoveFriend
	MethodRemoveFriend = 0x4

	// MethodAddFriendRequest is the method ID for method AddFriendRequest
	MethodAddFriendRequest = 0x5

	// MethodCancelFriendRequest is the method ID for method CancelFriendRequest
	MethodCancelFriendRequest = 0x6

	// MethodAcceptFriendRequest is the method ID for method AcceptFriendRequest
	MethodAcceptFriendRequest = 0x7

	// MethodDeleteFriendRequest is the method ID for method DeleteFriendRequest
	MethodDeleteFriendRequest = 0x8

	// MethodDenyFriendRequest is the method ID for method DenyFriendRequest
	MethodDenyFriendRequest = 0x9

	// MethodMarkFriendRequestsAsReceived is the method ID for method MarkFriendRequestsAsReceived
	MethodMarkFriendRequestsAsReceived = 0xA

	// MethodAddBlackList is the method ID for method AddBlackList
	MethodAddBlackList = 0xB

	// MethodRemoveBlackList is the method ID for method RemoveBlackList
	MethodRemoveBlackList = 0xC

	// MethodUpdatePresence is the method ID for method UpdatePresence
	MethodUpdatePresence = 0xD

	// MethodUpdateMii is the method ID for method UpdateMii
	MethodUpdateMii = 0xE

	// MethodUpdateComment is the method ID for method UpdateComment
	MethodUpdateComment = 0xF

	// MethodUpdatePreference is the method ID for method UpdatePreference
	MethodUpdatePreference = 0x10

	// MethodGetBasicInfo is the method ID for method GetBasicInfo
	MethodGetBasicInfo = 0x11

	// MethodDeletePersistentNotification is the method ID for method DeletePersistentNotification
	MethodDeletePersistentNotification = 0x12

	// MethodCheckSettingStatus is the method ID for method CheckSettingStatus
	MethodCheckSettingStatus = 0x13

	// MethodGetRequestBlockSettings is the method ID for method GetRequestBlockSettings
	MethodGetRequestBlockSettings = 0x14
)

// Protocol stores all the RMC method handlers for the Friends (WiiU) protocol and listens for requests
type Protocol struct {
	endpoint                     nex.EndpointInterface
	UpdateAndGetAllInformation   func(err error, packet nex.PacketInterface, callID uint32, nnaInfo *friends_wiiu_types.NNAInfo, presence *friends_wiiu_types.NintendoPresenceV2, birthday *types.DateTime) (*nex.RMCMessage, *nex.Error)
	AddFriend                    func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, *nex.Error)
	AddFriendByName              func(err error, packet nex.PacketInterface, callID uint32, username *types.String) (*nex.RMCMessage, *nex.Error)
	RemoveFriend                 func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, *nex.Error)
	AddFriendRequest             func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, unknown2 *types.PrimitiveU8, message *types.String, unknown4 *types.PrimitiveU8, unknown5 *types.String, gameKey *friends_wiiu_types.GameKey, unknown6 *types.DateTime) (*nex.RMCMessage, *nex.Error)
	CancelFriendRequest          func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	AcceptFriendRequest          func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	DeleteFriendRequest          func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	DenyFriendRequest            func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	MarkFriendRequestsAsReceived func(err error, packet nex.PacketInterface, callID uint32, ids *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error)
	AddBlackList                 func(err error, packet nex.PacketInterface, callID uint32, blacklistedPrincipal *friends_wiiu_types.BlacklistedPrincipal) (*nex.RMCMessage, *nex.Error)
	RemoveBlackList              func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, *nex.Error)
	UpdatePresence               func(err error, packet nex.PacketInterface, callID uint32, presence *friends_wiiu_types.NintendoPresenceV2) (*nex.RMCMessage, *nex.Error)
	UpdateMii                    func(err error, packet nex.PacketInterface, callID uint32, mii *friends_wiiu_types.MiiV2) (*nex.RMCMessage, *nex.Error)
	UpdateComment                func(err error, packet nex.PacketInterface, callID uint32, comment *friends_wiiu_types.Comment) (*nex.RMCMessage, *nex.Error)
	UpdatePreference             func(err error, packet nex.PacketInterface, callID uint32, preference *friends_wiiu_types.PrincipalPreference) (*nex.RMCMessage, *nex.Error)
	GetBasicInfo                 func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)
	DeletePersistentNotification func(err error, packet nex.PacketInterface, callID uint32, notifications *types.List[*friends_wiiu_types.PersistentNotification]) (*nex.RMCMessage, *nex.Error)
	CheckSettingStatus           func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetRequestBlockSettings      func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
	Patches                      nex.ServiceProtocol
	PatchedMethods               []uint32
}

// Interface implements the methods present on the Friends WiiU protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerUpdateAndGetAllInformation(handler func(err error, packet nex.PacketInterface, callID uint32, nnaInfo *friends_wiiu_types.NNAInfo, presence *friends_wiiu_types.NintendoPresenceV2, birthday *types.DateTime) (*nex.RMCMessage, *nex.Error))
	SetHandlerAddFriend(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerAddFriendByName(handler func(err error, packet nex.PacketInterface, callID uint32, username *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerRemoveFriend(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerAddFriendRequest(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, unknown2 *types.PrimitiveU8, message *types.String, unknown4 *types.PrimitiveU8, unknown5 *types.String, gameKey *friends_wiiu_types.GameKey, unknown6 *types.DateTime) (*nex.RMCMessage, *nex.Error))
	SetHandlerCancelFriendRequest(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerAcceptFriendRequest(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteFriendRequest(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerDenyFriendRequest(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerMarkFriendRequestsAsReceived(handler func(err error, packet nex.PacketInterface, callID uint32, ids *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error))
	SetHandlerAddBlackList(handler func(err error, packet nex.PacketInterface, callID uint32, blacklistedPrincipal *friends_wiiu_types.BlacklistedPrincipal) (*nex.RMCMessage, *nex.Error))
	SetHandlerRemoveBlackList(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdatePresence(handler func(err error, packet nex.PacketInterface, callID uint32, presence *friends_wiiu_types.NintendoPresenceV2) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateMii(handler func(err error, packet nex.PacketInterface, callID uint32, mii *friends_wiiu_types.MiiV2) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateComment(handler func(err error, packet nex.PacketInterface, callID uint32, comment *friends_wiiu_types.Comment) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdatePreference(handler func(err error, packet nex.PacketInterface, callID uint32, preference *friends_wiiu_types.PrincipalPreference) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetBasicInfo(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeletePersistentNotification(handler func(err error, packet nex.PacketInterface, callID uint32, notifications *types.List[*friends_wiiu_types.PersistentNotification]) (*nex.RMCMessage, *nex.Error))
	SetHandlerCheckSettingStatus(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRequestBlockSettings(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerUpdateAndGetAllInformation sets the handler for the UpdateAndGetAllInformation method
func (protocol *Protocol) SetHandlerUpdateAndGetAllInformation(handler func(err error, packet nex.PacketInterface, callID uint32, nnaInfo *friends_wiiu_types.NNAInfo, presence *friends_wiiu_types.NintendoPresenceV2, birthday *types.DateTime) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateAndGetAllInformation = handler
}

// SetHandlerAddFriend sets the handler for the AddFriend method
func (protocol *Protocol) SetHandlerAddFriend(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.AddFriend = handler
}

// SetHandlerAddFriendByName sets the handler for the AddFriendByName method
func (protocol *Protocol) SetHandlerAddFriendByName(handler func(err error, packet nex.PacketInterface, callID uint32, username *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.AddFriendByName = handler
}

// SetHandlerRemoveFriend sets the handler for the RemoveFriend method
func (protocol *Protocol) SetHandlerRemoveFriend(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.RemoveFriend = handler
}

// SetHandlerAddFriendRequest sets the handler for the AddFriendRequest method
func (protocol *Protocol) SetHandlerAddFriendRequest(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, unknown2 *types.PrimitiveU8, message *types.String, unknown4 *types.PrimitiveU8, unknown5 *types.String, gameKey *friends_wiiu_types.GameKey, unknown6 *types.DateTime) (*nex.RMCMessage, *nex.Error)) {
	protocol.AddFriendRequest = handler
}

// SetHandlerCancelFriendRequest sets the handler for the CancelFriendRequest method
func (protocol *Protocol) SetHandlerCancelFriendRequest(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.CancelFriendRequest = handler
}

// SetHandlerAcceptFriendRequest sets the handler for the AcceptFriendRequest method
func (protocol *Protocol) SetHandlerAcceptFriendRequest(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.AcceptFriendRequest = handler
}

// SetHandlerDeleteFriendRequest sets the handler for the DeleteFriendRequest method
func (protocol *Protocol) SetHandlerDeleteFriendRequest(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteFriendRequest = handler
}

// SetHandlerDenyFriendRequest sets the handler for the DenyFriendRequest method
func (protocol *Protocol) SetHandlerDenyFriendRequest(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.DenyFriendRequest = handler
}

// SetHandlerMarkFriendRequestsAsReceived sets the handler for the MarkFriendRequestsAsReceived method
func (protocol *Protocol) SetHandlerMarkFriendRequestsAsReceived(handler func(err error, packet nex.PacketInterface, callID uint32, ids *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, *nex.Error)) {
	protocol.MarkFriendRequestsAsReceived = handler
}

// SetHandlerAddBlackList sets the handler for the AddBlackList method
func (protocol *Protocol) SetHandlerAddBlackList(handler func(err error, packet nex.PacketInterface, callID uint32, blacklistedPrincipal *friends_wiiu_types.BlacklistedPrincipal) (*nex.RMCMessage, *nex.Error)) {
	protocol.AddBlackList = handler
}

// SetHandlerRemoveBlackList sets the handler for the RemoveBlackList method
func (protocol *Protocol) SetHandlerRemoveBlackList(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.RemoveBlackList = handler
}

// SetHandlerUpdatePresence sets the handler for the UpdatePresence method
func (protocol *Protocol) SetHandlerUpdatePresence(handler func(err error, packet nex.PacketInterface, callID uint32, presence *friends_wiiu_types.NintendoPresenceV2) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdatePresence = handler
}

// SetHandlerUpdateMii sets the handler for the UpdateMii method
func (protocol *Protocol) SetHandlerUpdateMii(handler func(err error, packet nex.PacketInterface, callID uint32, mii *friends_wiiu_types.MiiV2) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateMii = handler
}

// SetHandlerUpdateComment sets the handler for the UpdateComment method
func (protocol *Protocol) SetHandlerUpdateComment(handler func(err error, packet nex.PacketInterface, callID uint32, comment *friends_wiiu_types.Comment) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateComment = handler
}

// SetHandlerUpdatePreference sets the handler for the UpdatePreference method
func (protocol *Protocol) SetHandlerUpdatePreference(handler func(err error, packet nex.PacketInterface, callID uint32, preference *friends_wiiu_types.PrincipalPreference) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdatePreference = handler
}

// SetHandlerGetBasicInfo sets the handler for the GetBasicInfo method
func (protocol *Protocol) SetHandlerGetBasicInfo(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetBasicInfo = handler
}

// SetHandlerDeletePersistentNotification sets the handler for the DeletePersistentNotification method
func (protocol *Protocol) SetHandlerDeletePersistentNotification(handler func(err error, packet nex.PacketInterface, callID uint32, notifications *types.List[*friends_wiiu_types.PersistentNotification]) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeletePersistentNotification = handler
}

// SetHandlerCheckSettingStatus sets the handler for the CheckSettingStatus method
func (protocol *Protocol) SetHandlerCheckSettingStatus(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.CheckSettingStatus = handler
}

// SetHandlerGetRequestBlockSettings sets the handler for the GetRequestBlockSettings method
func (protocol *Protocol) SetHandlerGetRequestBlockSettings(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRequestBlockSettings = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if protocol.Patches != nil && slices.Contains(protocol.PatchedMethods, message.MethodID) {
		protocol.Patches.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodUpdateAndGetAllInformation:
		protocol.handleUpdateAndGetAllInformation(packet)
	case MethodAddFriend:
		protocol.handleAddFriend(packet)
	case MethodAddFriendByName:
		protocol.handleAddFriendByName(packet)
	case MethodRemoveFriend:
		protocol.handleRemoveFriend(packet)
	case MethodAddFriendRequest:
		protocol.handleAddFriendRequest(packet)
	case MethodCancelFriendRequest:
		protocol.handleCancelFriendRequest(packet)
	case MethodAcceptFriendRequest:
		protocol.handleAcceptFriendRequest(packet)
	case MethodDeleteFriendRequest:
		protocol.handleDeleteFriendRequest(packet)
	case MethodDenyFriendRequest:
		protocol.handleDenyFriendRequest(packet)
	case MethodMarkFriendRequestsAsReceived:
		protocol.handleMarkFriendRequestsAsReceived(packet)
	case MethodAddBlackList:
		protocol.handleAddBlackList(packet)
	case MethodRemoveBlackList:
		protocol.handleRemoveBlackList(packet)
	case MethodUpdatePresence:
		protocol.handleUpdatePresence(packet)
	case MethodUpdateMii:
		protocol.handleUpdateMii(packet)
	case MethodUpdateComment:
		protocol.handleUpdateComment(packet)
	case MethodUpdatePreference:
		protocol.handleUpdatePreference(packet)
	case MethodGetBasicInfo:
		protocol.handleGetBasicInfo(packet)
	case MethodDeletePersistentNotification:
		protocol.handleDeletePersistentNotification(packet)
	case MethodCheckSettingStatus:
		protocol.handleCheckSettingStatus(packet)
	case MethodGetRequestBlockSettings:
		protocol.handleGetRequestBlockSettings(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Friends (WiiU) method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Friends (WiiU) protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
