// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Friends protocol
	ProtocolID = 0x14

	// MethodAddFriend is the method ID for method AddFriend
	MethodAddFriend = 0x1

	// MethodAddFriendByName is the method ID for method AddFriendByName
	MethodAddFriendByName = 0x2

	// MethodAddFriendWithDetails is the method ID for method AddFriendWithDetails
	MethodAddFriendWithDetails = 0x3

	// MethodAddFriendByNameWithDetails is the method ID for method AddFriendByNameWithDetails
	MethodAddFriendByNameWithDetails = 0x4

	// MethodAcceptFriendship is the method ID for method AcceptFriendship
	MethodAcceptFriendship = 0x5

	// MethodDeclineFriendship is the method ID for method DeclineFriendship
	MethodDeclineFriendship = 0x6

	// MethodBlackList is the method ID for method BlackList
	MethodBlackList = 0x7

	// MethodBlackListByName is the method ID for method BlackListByName
	MethodBlackListByName = 0x8

	// MethodClearRelationship is the method ID for method ClearRelationship
	MethodClearRelationship = 0x9

	// MethodUpdateDetails is the method ID for method UpdateDetails
	MethodUpdateDetails = 0xA

	// MethodGetList is the method ID for method GetList
	MethodGetList = 0xB

	// MethodGetDetailedList is the method ID for method GetDetailedList
	MethodGetDetailedList = 0xC

	// MethodGetRelationships is the method ID for method GetRelationships
	MethodGetRelationships = 0xD
)

// Protocol handles the Friends QRV protocol
type Protocol struct {
	endpoint                   nex.EndpointInterface
	AddFriend                  func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	AddFriendByName            func(err error, packet nex.PacketInterface, callID uint32, strPlayerName *types.String, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	AddFriendWithDetails       func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	AddFriendByNameWithDetails func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error) // TODO - Is this the right signature?
	AcceptFriendship           func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	DeclineFriendship          func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	BlackList                  func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	BlackListByName            func(err error, packet nex.PacketInterface, callID uint32, strPlayerName *types.String, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	ClearRelationship          func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	UpdateDetails              func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetList                    func(err error, packet nex.PacketInterface, callID uint32, byRelationship *types.PrimitiveU8, bReversed *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetDetailedList            func(err error, packet nex.PacketInterface, callID uint32, byRelationship *types.PrimitiveU8, bReversed *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetRelationships           func(err error, packet nex.PacketInterface, callID uint32, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	Patches                    nex.ServiceProtocol
	PatchedMethods             []uint32
}

// Interface implements the methods present on the Friends protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerAddFriend(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerAddFriendByName(handler func(err error, packet nex.PacketInterface, callID uint32, strPlayerName *types.String, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerAddFriendWithDetails(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerAddFriendByNameWithDetails(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerAcceptFriendship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeclineFriendship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerBlackList(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerBlackListByName(handler func(err error, packet nex.PacketInterface, callID uint32, strPlayerName *types.String, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerClearRelationship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateDetails(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetList(handler func(err error, packet nex.PacketInterface, callID uint32, byRelationship *types.PrimitiveU8, bReversed *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetDetailedList(handler func(err error, packet nex.PacketInterface, callID uint32, byRelationship *types.PrimitiveU8, bReversed *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRelationships(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerAddFriend sets the handler for the AddFriend method
func (protocol *Protocol) SetHandlerAddFriend(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.AddFriend = handler
}

// SetHandlerAddFriendByName sets the handler for the AddFriendByName method
func (protocol *Protocol) SetHandlerAddFriendByName(handler func(err error, packet nex.PacketInterface, callID uint32, strPlayerName *types.String, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.AddFriendByName = handler
}

// SetHandlerAddFriendWithDetails sets the handler for the AddFriendWithDetails method
func (protocol *Protocol) SetHandlerAddFriendWithDetails(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.AddFriendWithDetails = handler
}

// SetHandlerAddFriendByNameWithDetails sets the handler for the AddFriendByNameWithDetails method
func (protocol *Protocol) SetHandlerAddFriendByNameWithDetails(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.AddFriendByNameWithDetails = handler
}

// SetHandlerAcceptFriendship sets the handler for the AcceptFriendship method
func (protocol *Protocol) SetHandlerAcceptFriendship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.AcceptFriendship = handler
}

// SetHandlerDeclineFriendship sets the handler for the DeclineFriendship method
func (protocol *Protocol) SetHandlerDeclineFriendship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeclineFriendship = handler
}

// SetHandlerBlackList sets the handler for the BlackList method
func (protocol *Protocol) SetHandlerBlackList(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.BlackList = handler
}

// SetHandlerBlackListByName sets the handler for the BlackListByName method
func (protocol *Protocol) SetHandlerBlackListByName(handler func(err error, packet nex.PacketInterface, callID uint32, strPlayerName *types.String, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.BlackListByName = handler
}

// SetHandlerClearRelationship sets the handler for the ClearRelationship method
func (protocol *Protocol) SetHandlerClearRelationship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ClearRelationship = handler
}

// SetHandlerUpdateDetails sets the handler for the UpdateDetails method
func (protocol *Protocol) SetHandlerUpdateDetails(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateDetails = handler
}

// SetHandlerGetList sets the handler for the GetList method
func (protocol *Protocol) SetHandlerGetList(handler func(err error, packet nex.PacketInterface, callID uint32, byRelationship *types.PrimitiveU8, bReversed *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetList = handler
}

// SetHandlerGetDetailedList sets the handler for the GetDetailedList method
func (protocol *Protocol) SetHandlerGetDetailedList(handler func(err error, packet nex.PacketInterface, callID uint32, byRelationship *types.PrimitiveU8, bReversed *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetDetailedList = handler
}

// SetHandlerGetRelationships sets the handler for the GetRelationships method
func (protocol *Protocol) SetHandlerGetRelationships(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRelationships = handler
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
	case MethodAddFriend:
		protocol.handleAddFriend(packet)
	case MethodAddFriendByName:
		protocol.handleAddFriendByName(packet)
	case MethodAddFriendWithDetails:
		protocol.handleAddFriendWithDetails(packet)
	case MethodAddFriendByNameWithDetails:
		protocol.handleAddFriendByNameWithDetails(packet)
	case MethodAcceptFriendship:
		protocol.handleAcceptFriendship(packet)
	case MethodDeclineFriendship:
		protocol.handleDeclineFriendship(packet)
	case MethodBlackList:
		protocol.handleBlackList(packet)
	case MethodBlackListByName:
		protocol.handleBlackListByName(packet)
	case MethodClearRelationship:
		protocol.handleClearRelationship(packet)
	case MethodUpdateDetails:
		protocol.handleUpdateDetails(packet)
	case MethodGetList:
		protocol.handleGetList(packet)
	case MethodGetDetailedList:
		protocol.handleGetDetailedList(packet)
	case MethodGetRelationships:
		protocol.handleGetRelationships(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Friends method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Friends protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
