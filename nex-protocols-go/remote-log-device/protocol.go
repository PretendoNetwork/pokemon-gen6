// Package protocol implements the Remote Log Device protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the RemoteLogDevice protocol
	ProtocolID = 0x1

	// MethodLog is the method ID for the method Log
	MethodLog = 0x1
)

// Protocol handles the RemoteLogDevice protocol
type Protocol struct {
	endpoint       nex.EndpointInterface
	Log            func(err error, packet nex.PacketInterface, callID uint32, strLine *types.String) (*nex.RMCMessage, *nex.Error)
	Patches        nex.ServiceProtocol
	PatchedMethods []uint32
}

// Interface implements the methods present on the Remote Log Device protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerLog(handler func(err error, packet nex.PacketInterface, callID uint32, strLine *types.String) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerLog sets the handler for the Log method
func (protocol *Protocol) SetHandlerLog(handler func(err error, packet nex.PacketInterface, callID uint32, strLine *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.Log = handler
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
	case MethodLog:
		protocol.handleLog(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported RemoteLogDevice method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Remote Log Device protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
