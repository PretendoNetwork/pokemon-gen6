// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/v2/subscriber/types"
)

const (
	// ProtocolID is the protocol ID for the Subscriber protocol
	ProtocolID = 0x6E

	// MethodHello is the method ID for the method Hello
	MethodHello = 0x1

	// MethodPostContent is the method ID for the method PostContent
	MethodPostContent = 0x2

	// MethodGetContent is the method ID for the method GetContent
	MethodGetContent = 0x3

	// MethodFollow is the method ID for the method Follow
	MethodFollow = 0x4

	// MethodUnfollowAllAndFollow is the method ID for the method UnfollowAllAndFollow
	MethodUnfollowAllAndFollow = 0x5

	// MethodUnfollow is the method ID for the method Unfollow
	MethodUnfollow = 0x6

	// MethodGetFollowing is the method ID for the method GetFollowing
	MethodGetFollowing = 0x7

	// MethodGetFollower is the method ID for the method GetFollower
	MethodGetFollower = 0x8

	// MethodGetNumFollowers is the method ID for the method GetNumFollowers
	MethodGetNumFollowers = 0x9

	// MethodGetTimeline is the method ID for the method GetTimeline
	MethodGetTimeline = 0xA

	// MethodDeleteContent is the method ID for the method DeleteContent
	MethodDeleteContent = 0xB

	// MethodGetContentMulti is the method ID for the method GetContentMulti
	MethodGetContentMulti = 0xC

	// MethodUpdateUserStatus is the method ID for the method UpdateUserStatus
	MethodUpdateUserStatus = 0xD

	// MethodGetFriendUserStatuses is the method ID for the method GetFriendUserStatuses
	MethodGetFriendUserStatuses = 0xE

	// MethodGetUserStatuses is the method ID for the method GetUserStatuses
	MethodGetUserStatuses = 0xF
)

// Protocol stores all the RMC method handlers for the Subscriber protocol and listens for requests
type Protocol struct {
	endpoint              nex.EndpointInterface
	Hello                 func(err error, packet nex.PacketInterface, callID uint32, unknown *types.String) (*nex.RMCMessage, *nex.Error)
	PostContent           func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberPostContentParam) (*nex.RMCMessage, *nex.Error)
	GetContent            func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberGetContentParam) (*nex.RMCMessage, *nex.Error)
	Follow                func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	UnfollowAllAndFollow  func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	Unfollow              func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	GetFollowing          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	GetFollower           func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	GetNumFollowers       func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	GetTimeline           func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	DeleteContent         func(err error, packet nex.PacketInterface, callID uint32, unknown1 *types.List[*types.String], unknown2 *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetContentMulti       func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*subscriber_types.SubscriberGetContentParam]) (*nex.RMCMessage, *nex.Error)
	UpdateUserStatus      func(err error, packet nex.PacketInterface, callID uint32, unknown1 *types.List[*subscriber_types.Unknown], unknown2 *types.List[*types.PrimitiveU8]) (*nex.RMCMessage, *nex.Error)
	GetFriendUserStatuses func(err error, packet nex.PacketInterface, callID uint32, unknown *types.List[*types.PrimitiveU8]) (*nex.RMCMessage, *nex.Error)
	GetUserStatuses       func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID], unknown *types.List[*types.PrimitiveU8]) (*nex.RMCMessage, *nex.Error)
	Patches               nex.ServiceProtocol
	PatchedMethods        []uint32
}

// Interface implements the methods present on the Subscriber protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerHello(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerPostContent(handler func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberPostContentParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetContent(handler func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberGetContentParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerFollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnfollowAllAndFollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnfollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetFollowing(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetFollower(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNumFollowers(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetTimeline(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteContent(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 *types.List[*types.String], unknown2 *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetContentMulti(handler func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*subscriber_types.SubscriberGetContentParam]) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateUserStatus(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 *types.List[*subscriber_types.Unknown], unknown2 *types.List[*types.PrimitiveU8]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetFriendUserStatuses(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.List[*types.PrimitiveU8]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetUserStatuses(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID], unknown *types.List[*types.PrimitiveU8]) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerHello sets the handler for the Hello method
func (protocol *Protocol) SetHandlerHello(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.Hello = handler
}

// SetHandlerPostContent sets the handler for the PostContent method
func (protocol *Protocol) SetHandlerPostContent(handler func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberPostContentParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.PostContent = handler
}

// SetHandlerGetContent sets the handler for the GetContent method
func (protocol *Protocol) SetHandlerGetContent(handler func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberGetContentParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetContent = handler
}

// SetHandlerFollow sets the handler for the Follow method
func (protocol *Protocol) SetHandlerFollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.Follow = handler
}

// SetHandlerUnfollowAllAndFollow sets the handler for the UnfollowAllAndFollow method
func (protocol *Protocol) SetHandlerUnfollowAllAndFollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.UnfollowAllAndFollow = handler
}

// SetHandlerUnfollow sets the handler for the Unfollow method
func (protocol *Protocol) SetHandlerUnfollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unfollow = handler
}

// SetHandlerGetFollowing sets the handler for the GetFollowing method
func (protocol *Protocol) SetHandlerGetFollowing(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetFollowing = handler
}

// SetHandlerGetFollower sets the handler for the GetFollower method
func (protocol *Protocol) SetHandlerGetFollower(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetFollower = handler
}

// SetHandlerGetNumFollowers sets the handler for the GetNumFollowers method
func (protocol *Protocol) SetHandlerGetNumFollowers(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNumFollowers = handler
}

// SetHandlerGetTimeline sets the handler for the GetTimeline method
func (protocol *Protocol) SetHandlerGetTimeline(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetTimeline = handler
}

// SetHandlerDeleteContent sets the handler for the DeleteContent method
func (protocol *Protocol) SetHandlerDeleteContent(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 *types.List[*types.String], unknown2 *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteContent = handler
}

// SetHandlerGetContentMulti sets the handler for the GetContentMulti method
func (protocol *Protocol) SetHandlerGetContentMulti(handler func(err error, packet nex.PacketInterface, callID uint32, params *types.List[*subscriber_types.SubscriberGetContentParam]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetContentMulti = handler
}

// SetHandlerUpdateUserStatus sets the handler for the UpdateUserStatus method
func (protocol *Protocol) SetHandlerUpdateUserStatus(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 *types.List[*subscriber_types.Unknown], unknown2 *types.List[*types.PrimitiveU8]) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateUserStatus = handler
}

// SetHandlerGetFriendUserStatuses sets the handler for the GetFriendUserStatuses method
func (protocol *Protocol) SetHandlerGetFriendUserStatuses(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.List[*types.PrimitiveU8]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetFriendUserStatuses = handler
}

// SetHandlerGetUserStatuses sets the handler for the GetUserStatuses method
func (protocol *Protocol) SetHandlerGetUserStatuses(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID], unknown *types.List[*types.PrimitiveU8]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetUserStatuses = handler
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
	case MethodHello:
		protocol.handleHello(packet)
	case MethodPostContent:
		protocol.handlePostContent(packet)
	case MethodGetContent:
		protocol.handleGetContent(packet)
	case MethodFollow:
		protocol.handleFollow(packet)
	case MethodUnfollowAllAndFollow:
		protocol.handleUnfollowAllAndFollow(packet)
	case MethodUnfollow:
		protocol.handleUnfollow(packet)
	case MethodGetFollowing:
		protocol.handleGetFollowing(packet)
	case MethodGetFollower:
		protocol.handleGetFollower(packet)
	case MethodGetNumFollowers:
		protocol.handleGetNumFollowers(packet)
	case MethodGetTimeline:
		protocol.handleGetTimeline(packet)
	case MethodDeleteContent:
		protocol.handleDeleteContent(packet)
	case MethodGetContentMulti:
		protocol.handleGetContentMulti(packet)
	case MethodUpdateUserStatus:
		protocol.handleUpdateUserStatus(packet)
	case MethodGetFriendUserStatuses:
		protocol.handleGetFriendUserStatuses(packet)
	case MethodGetUserStatuses:
		protocol.handleGetUserStatuses(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Subscriber method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Subscriber protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
