// Package protocol implements the Utility protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	utility_types "github.com/PretendoNetwork/nex-protocols-go/v2/utility/types"
)

const (
	// ProtocolID is the protocol ID for the Utility protocol
	ProtocolID = 0x6E

	// MethodAcquireNexUniqueID is the method ID for the method AcquireNexUniqueID
	MethodAcquireNexUniqueID = 0x1

	// MethodAcquireNexUniqueIDWithPassword is the method ID for the method AcquireNexUniqueIDWithPassword
	MethodAcquireNexUniqueIDWithPassword = 0x2

	// MethodAssociateNexUniqueIDWithMyPrincipalID is the method ID for the method AssociateNexUniqueIDWithMyPrincipalID
	MethodAssociateNexUniqueIDWithMyPrincipalID = 0x3

	// MethodAssociateNexUniqueIDsWithMyPrincipalID is the method ID for the method AssociateNexUniqueIDsWithMyPrincipalID
	MethodAssociateNexUniqueIDsWithMyPrincipalID = 0x4

	// MethodGetAssociatedNexUniqueIDWithMyPrincipalID is the method ID for the method GetAssociatedNexUniqueIDWithMyPrincipalID
	MethodGetAssociatedNexUniqueIDWithMyPrincipalID = 0x5

	// MethodGetAssociatedNexUniqueIDsWithMyPrincipalID is the method ID for the method GetAssociatedNexUniqueIDsWithMyPrincipalID
	MethodGetAssociatedNexUniqueIDsWithMyPrincipalID = 0x6

	// MethodGetIntegerSettings is the method ID for the method GetIntegerSettings
	MethodGetIntegerSettings = 0x7

	// MethodGetStringSettings is the method ID for the method GetStringSettings
	MethodGetStringSettings = 0x8
)

// Protocol stores all the RMC method handlers for the Utility protocol and listens for requests
type Protocol struct {
	endpoint                                   nex.EndpointInterface
	AcquireNexUniqueID                         func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	AcquireNexUniqueIDWithPassword             func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	AssociateNexUniqueIDWithMyPrincipalID      func(err error, packet nex.PacketInterface, callID uint32, uniqueIDInfo *utility_types.UniqueIDInfo) (*nex.RMCMessage, *nex.Error)
	AssociateNexUniqueIDsWithMyPrincipalID     func(err error, packet nex.PacketInterface, callID uint32, uniqueIDInfo *types.List[*utility_types.UniqueIDInfo]) (*nex.RMCMessage, *nex.Error)
	GetAssociatedNexUniqueIDWithMyPrincipalID  func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetAssociatedNexUniqueIDsWithMyPrincipalID func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetIntegerSettings                         func(err error, packet nex.PacketInterface, callID uint32, integerSettingIndex *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetStringSettings                          func(err error, packet nex.PacketInterface, callID uint32, stringSettingIndex *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	Patches                                    nex.ServiceProtocol
	PatchedMethods                             []uint32
}

// Interface implements the methods present on the Utility protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerAcquireNexUniqueID(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerAcquireNexUniqueIDWithPassword(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerAssociateNexUniqueIDWithMyPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueIDInfo *utility_types.UniqueIDInfo) (*nex.RMCMessage, *nex.Error))
	SetHandlerAssociateNexUniqueIDsWithMyPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueIDInfo *types.List[*utility_types.UniqueIDInfo]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetAssociatedNexUniqueIDWithMyPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetAssociatedNexUniqueIDsWithMyPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetIntegerSettings(handler func(err error, packet nex.PacketInterface, callID uint32, integerSettingIndex *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetStringSettings(handler func(err error, packet nex.PacketInterface, callID uint32, stringSettingIndex *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerAcquireNexUniqueID sets the handler for the AcquireNexUniqueID method
func (protocol *Protocol) SetHandlerAcquireNexUniqueID(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.AcquireNexUniqueID = handler
}

// SetHandlerAcquireNexUniqueIDWithPassword sets the handler for the AcquireNexUniqueIDWithPassword method
func (protocol *Protocol) SetHandlerAcquireNexUniqueIDWithPassword(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.AcquireNexUniqueIDWithPassword = handler
}

// SetHandlerAssociateNexUniqueIDWithMyPrincipalID sets the handler for the AssociateNexUniqueIDWithMyPrincipalID method
func (protocol *Protocol) SetHandlerAssociateNexUniqueIDWithMyPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueIDInfo *utility_types.UniqueIDInfo) (*nex.RMCMessage, *nex.Error)) {
	protocol.AssociateNexUniqueIDWithMyPrincipalID = handler
}

// SetHandlerAssociateNexUniqueIDsWithMyPrincipalID sets the handler for the AssociateNexUniqueIDsWithMyPrincipalID method
func (protocol *Protocol) SetHandlerAssociateNexUniqueIDsWithMyPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueIDInfo *types.List[*utility_types.UniqueIDInfo]) (*nex.RMCMessage, *nex.Error)) {
	protocol.AssociateNexUniqueIDsWithMyPrincipalID = handler
}

// SetHandlerGetAssociatedNexUniqueIDWithMyPrincipalID sets the handler for the GetAssociatedNexUniqueIDWithMyPrincipalID method
func (protocol *Protocol) SetHandlerGetAssociatedNexUniqueIDWithMyPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetAssociatedNexUniqueIDWithMyPrincipalID = handler
}

// SetHandlerGetAssociatedNexUniqueIDsWithMyPrincipalID sets the handler for the GetAssociatedNexUniqueIDsWithMyPrincipalID method
func (protocol *Protocol) SetHandlerGetAssociatedNexUniqueIDsWithMyPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetAssociatedNexUniqueIDsWithMyPrincipalID = handler
}

// SetHandlerGetIntegerSettings sets the handler for the GetIntegerSettings method
func (protocol *Protocol) SetHandlerGetIntegerSettings(handler func(err error, packet nex.PacketInterface, callID uint32, integerSettingIndex *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetIntegerSettings = handler
}

// SetHandlerGetStringSettings sets the handler for the GetStringSettings method
func (protocol *Protocol) SetHandlerGetStringSettings(handler func(err error, packet nex.PacketInterface, callID uint32, stringSettingIndex *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetStringSettings = handler
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
	case MethodAcquireNexUniqueID:
		protocol.handleAcquireNexUniqueID(packet)
	case MethodAcquireNexUniqueIDWithPassword:
		protocol.handleAcquireNexUniqueIDWithPassword(packet)
	case MethodAssociateNexUniqueIDWithMyPrincipalID:
		protocol.handleAssociateNexUniqueIDWithMyPrincipalID(packet)
	case MethodAssociateNexUniqueIDsWithMyPrincipalID:
		protocol.handleAssociateNexUniqueIDsWithMyPrincipalID(packet)
	case MethodGetAssociatedNexUniqueIDWithMyPrincipalID:
		protocol.handleGetAssociatedNexUniqueIDWithMyPrincipalID(packet)
	case MethodGetAssociatedNexUniqueIDsWithMyPrincipalID:
		protocol.handleGetAssociatedNexUniqueIDsWithMyPrincipalID(packet)
	case MethodGetIntegerSettings:
		protocol.handleGetIntegerSettings(packet)
	case MethodGetStringSettings:
		protocol.handleGetStringSettings(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Utility method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Utility protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
