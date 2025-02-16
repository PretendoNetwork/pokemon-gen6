// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Secure Connection protocol
	ProtocolID = 0xB

	// MethodRegister is the method ID for the method Register
	MethodRegister = 0x1

	// MethodRequestConnectionData is the method ID for the method RequestConnectionData
	MethodRequestConnectionData = 0x2

	// MethodRequestURLs is the method ID for the method RequestURLs
	MethodRequestURLs = 0x3

	// MethodRegisterEx is the method ID for the method RegisterEx
	MethodRegisterEx = 0x4

	// MethodTestConnectivity is the method ID for the method TestConnectivity
	MethodTestConnectivity = 0x5

	// MethodUpdateURLs is the method ID for the method UpdateURLs
	MethodUpdateURLs = 0x6

	// MethodReplaceURL is the method ID for the method ReplaceURL
	MethodReplaceURL = 0x7

	// MethodSendReport is the method ID for the method SendReport
	MethodSendReport = 0x8
)

// Protocol stores all the RMC method handlers for the Secure Connection protocol and listens for requests
type Protocol struct {
	endpoint              nex.EndpointInterface
	Register              func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error)
	RequestConnectionData func(err error, packet nex.PacketInterface, callID uint32, cidTarget *types.PrimitiveU32, pidTarget *types.PID) (*nex.RMCMessage, *nex.Error)
	RequestURLs           func(err error, packet nex.PacketInterface, callID uint32, cidTarget *types.PrimitiveU32, pidTarget *types.PID) (*nex.RMCMessage, *nex.Error)
	RegisterEx            func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL], hCustomData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	TestConnectivity      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	UpdateURLs            func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error)
	ReplaceURL            func(err error, packet nex.PacketInterface, callID uint32, target *types.StationURL, url *types.StationURL) (*nex.RMCMessage, *nex.Error)
	SendReport            func(err error, packet nex.PacketInterface, callID uint32, reportID *types.PrimitiveU32, reportData *types.QBuffer) (*nex.RMCMessage, *nex.Error)
	Patches               nex.ServiceProtocol
	PatchedMethods        []uint32
}

// Interface implements the methods present on the Secure Connection protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerRegister(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error))
	SetHandlerRequestConnectionData(handler func(err error, packet nex.PacketInterface, callID uint32, cidTarget *types.PrimitiveU32, pidTarget *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerRequestURLs(handler func(err error, packet nex.PacketInterface, callID uint32, cidTarget *types.PrimitiveU32, pidTarget *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerRegisterEx(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL], hCustomData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerTestConnectivity(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateURLs(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error))
	SetHandlerReplaceURL(handler func(err error, packet nex.PacketInterface, callID uint32, target *types.StationURL, url *types.StationURL) (*nex.RMCMessage, *nex.Error))
	SetHandlerSendReport(handler func(err error, packet nex.PacketInterface, callID uint32, reportID *types.PrimitiveU32, reportData *types.QBuffer) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerRegister sets the handler for the Register method
func (protocol *Protocol) SetHandlerRegister(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error)) {
	protocol.Register = handler
}

// SetHandlerRequestConnectionData sets the handler for the RequestConnectionData method
func (protocol *Protocol) SetHandlerRequestConnectionData(handler func(err error, packet nex.PacketInterface, callID uint32, cidTarget *types.PrimitiveU32, pidTarget *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.RequestConnectionData = handler
}

// SetHandlerRequestURLs sets the handler for the RequestURLs method
func (protocol *Protocol) SetHandlerRequestURLs(handler func(err error, packet nex.PacketInterface, callID uint32, cidTarget *types.PrimitiveU32, pidTarget *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.RequestURLs = handler
}

// SetHandlerRegisterEx sets the handler for the RegisterEx method
func (protocol *Protocol) SetHandlerRegisterEx(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL], hCustomData *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.RegisterEx = handler
}

// SetHandlerTestConnectivity sets the handler for the TestConnectivity method
func (protocol *Protocol) SetHandlerTestConnectivity(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.TestConnectivity = handler
}

// SetHandlerUpdateURLs sets the handler for the UpdateURLs method
func (protocol *Protocol) SetHandlerUpdateURLs(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateURLs = handler
}

// SetHandlerReplaceURL sets the handler for the ReplaceURL method
func (protocol *Protocol) SetHandlerReplaceURL(handler func(err error, packet nex.PacketInterface, callID uint32, target *types.StationURL, url *types.StationURL) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReplaceURL = handler
}

// SetHandlerSendReport sets the handler for the SendReport method
func (protocol *Protocol) SetHandlerSendReport(handler func(err error, packet nex.PacketInterface, callID uint32, reportID *types.PrimitiveU32, reportData *types.QBuffer) (*nex.RMCMessage, *nex.Error)) {
	protocol.SendReport = handler
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
	case MethodRegister:
		protocol.handleRegister(packet)
	case MethodRequestConnectionData:
		protocol.handleRequestConnectionData(packet)
	case MethodRequestURLs:
		protocol.handleRequestURLs(packet)
	case MethodRegisterEx:
		protocol.handleRegisterEx(packet)
	case MethodTestConnectivity:
		protocol.handleTestConnectivity(packet)
	case MethodUpdateURLs:
		protocol.handleUpdateURLs(packet)
	case MethodReplaceURL:
		protocol.handleReplaceURL(packet)
	case MethodSendReport:
		protocol.handleSendReport(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported SecureConnection method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Secure Connection protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
