// Package protocol implements the Message Deliver protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Message Delivery protocol
	ProtocolID = 0x1B

	// MethodDeliverMessage is the method ID for the method DeliverMessage
	MethodDeliverMessage = 0x1
)

// Protocol stores all the RMC method handlers for the Message Delivery protocol and listens for requests
type Protocol struct {
	endpoint       nex.EndpointInterface
	DeliverMessage func(err error, packet nex.PacketInterface, callID uint32, oUserMessage *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	Patches        nex.ServiceProtocol
	PatchedMethods []uint32
}

// Interface implements the methods present on the Message Deliver protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerDeliverMessage(handler func(err error, packet nex.PacketInterface, callID uint32, oUserMessage *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerDeliverMessage sets the handler for the DeliverMessage method
func (protocol *Protocol) SetHandlerDeliverMessage(handler func(err error, packet nex.PacketInterface, callID uint32, oUserMessage *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeliverMessage = handler
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
	case MethodDeliverMessage:
		protocol.handleDeliverMessage(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported MessageDelivery method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Message Delivery protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
