// Package protocol implements the Health protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Health protocol
	ProtocolID = 0x12

	// MethodPingDaemon is the method ID for the method PingDaemon
	MethodPingDaemon = 0x1

	// MethodPingDatabase is the method ID for the method PingDatabase
	MethodPingDatabase = 0x2

	// MethodRunSanityCheck is the method ID for the method RunSanityCheck
	MethodRunSanityCheck = 0x3

	// MethodFixSanityErrors is the method ID for the method FixSanityErrors
	MethodFixSanityErrors = 0x4
)

// Protocol handles the Health protocol
type Protocol struct {
	endpoint        nex.EndpointInterface
	PingDaemon      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	PingDatabase    func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	RunSanityCheck  func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	FixSanityErrors func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	Patches         nex.ServiceProtocol
	PatchedMethods  []uint32
}

// Interface implements the methods present on the Health protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerPingDaemon(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerPingDatabase(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerRunSanityCheck(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerFixSanityErrors(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerPingDaemon sets the handler for the PingDaemon method
func (protocol *Protocol) SetHandlerPingDaemon(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.PingDaemon = handler
}

// SetHandlerPingDatabase sets the handler for the PingDatabase method
func (protocol *Protocol) SetHandlerPingDatabase(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.PingDatabase = handler
}

// SetHandlerRunSanityCheck sets the handler for the RunSanityCheck method
func (protocol *Protocol) SetHandlerRunSanityCheck(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.RunSanityCheck = handler
}

// SetHandlerFixSanityErrors sets the handler for the FixSanityErrors method
func (protocol *Protocol) SetHandlerFixSanityErrors(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.FixSanityErrors = handler
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
	case MethodPingDaemon:
		protocol.handlePingDaemon(packet)
	case MethodPingDatabase:
		protocol.handlePingDatabase(packet)
	case MethodRunSanityCheck:
		protocol.handleRunSanityCheck(packet)
	case MethodFixSanityErrors:
		protocol.handleFixSanityErrors(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Health method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Health protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
