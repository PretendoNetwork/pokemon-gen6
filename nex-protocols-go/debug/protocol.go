// Package protocol implements the Debug protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Debug protocol
	ProtocolID = 0x74

	// MethodEnableAPIRecorder is the method ID for the method EnableAPIRecorder
	MethodEnableAPIRecorder = 0x1

	// MethodDisableAPIRecorder is the method ID for the method DisableAPIRecorder
	MethodDisableAPIRecorder = 0x2

	// MethodIsAPIRecorderEnabled is the method ID for the method IsAPIRecorderEnabled
	MethodIsAPIRecorderEnabled = 0x3

	// MethodGetAPICalls is the method ID for the method GetAPICalls
	MethodGetAPICalls = 0x4

	// MethodSetExcludeJoinedMatchmakeSession is the method ID for the method SetExcludeJoinedMatchmakeSession
	MethodSetExcludeJoinedMatchmakeSession = 0x5

	// MethodGetExcludeJoinedMatchmakeSession is the method ID for the method GetExcludeJoinedMatchmakeSession
	MethodGetExcludeJoinedMatchmakeSession = 0x6

	// MethodGetAPICallSummary is the method ID for the method GetAPICallSummary
	MethodGetAPICallSummary = 0x7
)

// Protocol handles the Debug protocol
type Protocol struct {
	endpoint                         nex.EndpointInterface
	EnableAPIRecorder                func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	DisableAPIRecorder               func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	IsAPIRecorderEnabled             func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetAPICalls                      func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID], unknown *types.DateTime, unknown2 *types.DateTime) (*nex.RMCMessage, *nex.Error)
	SetExcludeJoinedMatchmakeSession func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) // TODO - Unknown request/response format
	GetExcludeJoinedMatchmakeSession func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) // TODO - Unknown request/response format
	GetAPICallSummary                func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) // TODO - Unknown request/response format
	Patches                          nex.ServiceProtocol
	PatchedMethods                   []uint32
}

// Interface implements the methods present on the Debug protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerEnableAPIRecorder(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerDisableAPIRecorder(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerIsAPIRecorderEnabled(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetAPICalls(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID], unknown *types.DateTime, unknown2 *types.DateTime) (*nex.RMCMessage, *nex.Error))
	SetHandlerSetExcludeJoinedMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetExcludeJoinedMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetAPICallSummary(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
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
func (protocol *Protocol) SetHandlerEnableAPIRecorder(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.EnableAPIRecorder = handler
}

// SetHandlerDisableAPIRecorder sets the handler for the DisableAPIRecorder method
func (protocol *Protocol) SetHandlerDisableAPIRecorder(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.DisableAPIRecorder = handler
}

// SetHandlerIsAPIRecorderEnabled sets the handler for the IsAPIRecorderEnabled method
func (protocol *Protocol) SetHandlerIsAPIRecorderEnabled(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.IsAPIRecorderEnabled = handler
}

// SetHandlerGetAPICalls sets the handler for the GetAPICalls method
func (protocol *Protocol) SetHandlerGetAPICalls(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID], unknown *types.DateTime, unknown2 *types.DateTime) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetAPICalls = handler
}

// SetHandlerSetExcludeJoinedMatchmakeSession sets the handler for the SetExcludeJoinedMatchmakeSession method
func (protocol *Protocol) SetHandlerSetExcludeJoinedMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.SetExcludeJoinedMatchmakeSession = handler
}

// SetHandlerGetExcludeJoinedMatchmakeSession sets the handler for the GetExcludeJoinedMatchmakeSession method
func (protocol *Protocol) SetHandlerGetExcludeJoinedMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetExcludeJoinedMatchmakeSession = handler
}

// SetHandlerGetAPICallSummary sets the handler for the GetAPICallSummary method
func (protocol *Protocol) SetHandlerGetAPICallSummary(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetAPICallSummary = handler
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
	case MethodEnableAPIRecorder:
		protocol.handleEnableAPIRecorder(packet)
	case MethodDisableAPIRecorder:
		protocol.handleDisableAPIRecorder(packet)
	case MethodIsAPIRecorderEnabled:
		protocol.handleIsAPIRecorderEnabled(packet)
	case MethodGetAPICalls:
		protocol.handleGetAPICalls(packet)
	case MethodSetExcludeJoinedMatchmakeSession:
		protocol.handleSetExcludeJoinedMatchmakeSession(packet)
	case MethodGetExcludeJoinedMatchmakeSession:
		protocol.handleGetExcludeJoinedMatchmakeSession(packet)
	case MethodGetAPICallSummary:
		protocol.handleGetAPICallSummary(packet)
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
