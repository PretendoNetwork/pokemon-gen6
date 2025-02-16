// Package protocol implements the StorageManager protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the StorageManager protocol
	ProtocolID = 0x6E

	// MethodAcquireCardID is the method ID for the method AcquireCardID
	MethodAcquireCardID = 0x4

	// MethodActivateWithCardID is the method ID for the method ActivateWithCardID
	MethodActivateWithCardID = 0x5
)

// Protocol stores all the RMC method handlers for the StorageManager protocol and listens for requests
type Protocol struct {
	endpoint           nex.EndpointInterface
	AcquireCardID      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	ActivateWithCardID func(err error, packet nex.PacketInterface, callID uint32, unknown *types.PrimitiveU8, cardID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	Patches            nex.ServiceProtocol
	PatchedMethods     []uint32
}

// Interface implements the methods present on the StorageManager protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerAcquireCardID(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerActivateWithCardID(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.PrimitiveU8, cardID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerAcquireCardID sets the handler for the AcquireCardID method
func (protocol *Protocol) SetHandlerAcquireCardID(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.AcquireCardID = handler
}

// SetHandlerActivateWithCardID sets the handler for the ActivateWithCardID method
func (protocol *Protocol) SetHandlerActivateWithCardID(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.PrimitiveU8, cardID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.ActivateWithCardID = handler
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
	case MethodAcquireCardID:
		protocol.handleAcquireCardID(packet)
	case MethodActivateWithCardID:
		protocol.handleActivateWithCardID(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported StorageManager method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new StorageManager protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
