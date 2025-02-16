// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Ticket Granting protocol
	ProtocolID = 0xA

	// MethodLogin is the method ID for the method Login
	MethodLogin = 0x1

	// MethodLoginEx is the method ID for the method LoginEx
	MethodLoginEx = 0x2

	// MethodRequestTicket is the method ID for the method RequestTicket
	MethodRequestTicket = 0x3

	// MethodGetPID is the method ID for the method GetPID
	MethodGetPID = 0x4

	// MethodGetName is the method ID for the method GetName
	MethodGetName = 0x5

	// MethodLoginWithContext is the method ID for the method LoginWithContext
	MethodLoginWithContext = 0x6
)

// Protocol stores all the RMC method handlers for the Ticket Granting protocol and listens for requests
type Protocol struct {
	endpoint         nex.EndpointInterface
	Login            func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String) (*nex.RMCMessage, *nex.Error)
	LoginEx          func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String, oExtraData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	RequestTicket    func(err error, packet nex.PacketInterface, callID uint32, idSource *types.PID, idTarget *types.PID) (*nex.RMCMessage, *nex.Error)
	GetPID           func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String) (*nex.RMCMessage, *nex.Error)
	GetName          func(err error, packet nex.PacketInterface, callID uint32, id *types.PID) (*nex.RMCMessage, *nex.Error)
	LoginWithContext func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	Patches          nex.ServiceProtocol
	PatchedMethods   []uint32
}

// Interface implements the methods present on the Ticket Granting protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerLogin(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerLoginEx(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String, oExtraData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerRequestTicket(handler func(err error, packet nex.PacketInterface, callID uint32, idSource *types.PID, idTarget *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetPID(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetName(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerLoginWithContext(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerLogin sets the handler for the Login method
func (protocol *Protocol) SetHandlerLogin(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.Login = handler
}

// SetHandlerLoginEx sets the handler for the LoginEx method
func (protocol *Protocol) SetHandlerLoginEx(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String, oExtraData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.LoginEx = handler
}

// SetHandlerRequestTicket sets the handler for the RequestTicket method
func (protocol *Protocol) SetHandlerRequestTicket(handler func(err error, packet nex.PacketInterface, callID uint32, idSource *types.PID, idTarget *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.RequestTicket = handler
}

// SetHandlerGetPID sets the handler for the GetPID method
func (protocol *Protocol) SetHandlerGetPID(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetPID = handler
}

// SetHandlerGetName sets the handler for the GetName method
func (protocol *Protocol) SetHandlerGetName(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetName = handler
}

// SetHandlerLoginWithContext sets the handler for the LoginWithContext method
func (protocol *Protocol) SetHandlerLoginWithContext(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.LoginWithContext = handler
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
	case MethodLogin:
		protocol.handleLogin(packet)
	case MethodLoginEx:
		protocol.handleLoginEx(packet)
	case MethodRequestTicket:
		protocol.handleRequestTicket(packet)
	case MethodGetPID:
		protocol.handleGetPID(packet)
	case MethodGetName:
		protocol.handleGetName(packet)
	case MethodLoginWithContext:
		protocol.handleLoginWithContext(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Ticket Granting method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Ticket Granting protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
