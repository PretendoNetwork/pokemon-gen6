// Package protocol implements the Debug protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Debug protocol
	ProtocolID = 0x76

	// MethodEnableAPIRecorder is the method ID for the method EnableAPIRecorder
	MethodUnk1 = 0x1

	// MethodEnableAPIRecorder is the method ID for the method EnableAPIRecorder
	MethodUnk2 = 0x2
)

// Protocol handles the Debug protocol
type Protocol struct {
	endpoint                         nex.EndpointInterface
	Unk1                			 func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	Unk2                			 func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	Patches                          nex.ServiceProtocol
	PatchedMethods                   []uint32
}

// Interface implements the methods present on the Debug protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerUnk1(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnk2(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerEnableAPIRecorder sets the handler for the EnableAPIRecorder method
func (protocol *Protocol) SetHandlerUnk1(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk1 = handler
}

// SetHandlerEnableAPIRecorder sets the handler for the EnableAPIRecorder method
func (protocol *Protocol) SetHandlerUnk2(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk1 = handler
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
	case MethodUnk1:
		protocol.handleUnk1(packet)
	case MethodUnk2:
		protocol.handleUnk2(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Debug method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Debug protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
