// Package protocol implements the AAUser protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	aauser_types "github.com/PretendoNetwork/nex-protocols-go/v2/aa-user/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the AAUser protocol
	ProtocolID = 0x7B

	// MethodRegisterApplication is the method ID for the method RegisterApplication
	MethodRegisterApplication = 0x1

	// MethodUnregisterApplication is the method ID for the method UnregisterApplication
	MethodUnregisterApplication = 0x2

	// MethodSetApplicationInfo is the method ID for the method RegisterApplication
	MethodSetApplicationInfo = 0x3

	// MethodGetApplicationInfo is the method ID for the method GetApplicationInfo
	MethodGetApplicationInfo = 0x4
)

// Protocol stores all the RMC method handlers for the AAUser protocol and listens for requests
type Protocol struct {
	endpoint              nex.EndpointInterface
	RegisterApplication   func(err error, packet nex.PacketInterface, callID uint32, titleID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	UnregisterApplication func(err error, packet nex.PacketInterface, callID uint32, titleID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	SetApplicationInfo    func(err error, packet nex.PacketInterface, callID uint32, applicationInfo *types.List[*aauser_types.ApplicationInfo]) (*nex.RMCMessage, *nex.Error)
	GetApplicationInfo    func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	Patches               nex.ServiceProtocol
	PatchedMethods        []uint32
}

// Interface implements the methods present on the AAUser Protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerRegisterApplication(handler func(err error, packet nex.PacketInterface, callID uint32, titleID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnregisterApplication(handler func(err error, packet nex.PacketInterface, callID uint32, titleID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerSetApplicationInfo(handler func(err error, packet nex.PacketInterface, callID uint32, applicationInfo *types.List[*aauser_types.ApplicationInfo]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetApplicationInfo(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerRegisterApplication sets the handler for the RegisterApplication method
func (protocol *Protocol) SetHandlerRegisterApplication(handler func(err error, packet nex.PacketInterface, callID uint32, titleID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.RegisterApplication = handler
}

// SetHandlerUnregisterApplication sets the handler for the UnregisterApplication method
func (protocol *Protocol) SetHandlerUnregisterApplication(handler func(err error, packet nex.PacketInterface, callID uint32, titleID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.UnregisterApplication = handler
}

// SetHandlerSetApplicationInfo sets the handler for the SetApplicationInfo method
func (protocol *Protocol) SetHandlerSetApplicationInfo(handler func(err error, packet nex.PacketInterface, callID uint32, applicationInfo *types.List[*aauser_types.ApplicationInfo]) (*nex.RMCMessage, *nex.Error)) {
	protocol.SetApplicationInfo = handler
}

// SetHandlerGetApplicationInfo sets the handler for the GetApplicationInfo method
func (protocol *Protocol) SetHandlerGetApplicationInfo(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetApplicationInfo = handler
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
	case MethodRegisterApplication:
		protocol.handleRegisterApplication(packet)
	case MethodUnregisterApplication:
		protocol.handleUnregisterApplication(packet)
	case MethodSetApplicationInfo:
		protocol.handleSetApplicationInfo(packet)
	case MethodGetApplicationInfo:
		protocol.handleGetApplicationInfo(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported AAUser method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new AAUser protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
