// Package protocol implements the NAT Traversal protocol
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
	ProtocolID = 0x3

	// MethodRequestProbeInitiation is the method ID for the method RequestProbeInitiation
	MethodRequestProbeInitiation = 0x1

	// MethodInitiateProbe is the method ID for the method InitiateProbe
	MethodInitiateProbe = 0x2

	// MethodRequestProbeInitiationExt is the method ID for the method RequestProbeInitiationExt
	MethodRequestProbeInitiationExt = 0x3

	// MethodReportNATTraversalResult is the method ID for the method ReportNATTraversalResult
	MethodReportNATTraversalResult = 0x4

	// MethodReportNATProperties is the method ID for the method ReportNATProperties
	MethodReportNATProperties = 0x5

	// MethodGetRelaySignatureKey is the method ID for the method GetRelaySignatureKey
	MethodGetRelaySignatureKey = 0x6

	// MethodReportNATTraversalResultDetail is the method ID for the method ReportNATTraversalResultDetail
	MethodReportNATTraversalResultDetail = 0x7
)

// Protocol stores all the RMC method handlers for the NAT Traversal protocol and listens for requests
type Protocol struct {
	endpoint                       nex.EndpointInterface
	RequestProbeInitiation         func(err error, packet nex.PacketInterface, callID uint32, urlTargetList *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error)
	InitiateProbe                  func(err error, packet nex.PacketInterface, callID uint32, urlStationToProbe *types.StationURL) (*nex.RMCMessage, *nex.Error)
	RequestProbeInitiationExt      func(err error, packet nex.PacketInterface, callID uint32, targetList *types.List[*types.String], stationToProbe *types.String) (*nex.RMCMessage, *nex.Error)
	ReportNATTraversalResult       func(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, rtt *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	ReportNATProperties            func(err error, packet nex.PacketInterface, callID uint32, natmapping *types.PrimitiveU32, natfiltering *types.PrimitiveU32, rtt *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetRelaySignatureKey           func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	ReportNATTraversalResultDetail func(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, detail *types.PrimitiveS32, rtt *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	Patches                        nex.ServiceProtocol
	PatchedMethods                 []uint32
}

// Interface implements the methods present on the NAT Traversal protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerRequestProbeInitiation(handler func(err error, packet nex.PacketInterface, callID uint32, urlTargetList *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error))
	SetHandlerInitiateProbe(handler func(err error, packet nex.PacketInterface, callID uint32, urlStationToProbe *types.StationURL) (*nex.RMCMessage, *nex.Error))
	SetHandlerRequestProbeInitiationExt(handler func(err error, packet nex.PacketInterface, callID uint32, targetList *types.List[*types.String], stationToProbe *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerReportNATTraversalResult(handler func(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, rtt *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerReportNATProperties(handler func(err error, packet nex.PacketInterface, callID uint32, natmapping *types.PrimitiveU32, natfiltering *types.PrimitiveU32, rtt *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRelaySignatureKey(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerReportNATTraversalResultDetail(handler func(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, detail *types.PrimitiveS32, rtt *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerRequestProbeInitiation sets the handler for the RequestProbeInitiation method
func (protocol *Protocol) SetHandlerRequestProbeInitiation(handler func(err error, packet nex.PacketInterface, callID uint32, urlTargetList *types.List[*types.StationURL]) (*nex.RMCMessage, *nex.Error)) {
	protocol.RequestProbeInitiation = handler
}

// SetHandlerInitiateProbe sets the handler for the InitiateProbe method
func (protocol *Protocol) SetHandlerInitiateProbe(handler func(err error, packet nex.PacketInterface, callID uint32, urlStationToProbe *types.StationURL) (*nex.RMCMessage, *nex.Error)) {
	protocol.InitiateProbe = handler
}

// SetHandlerRequestProbeInitiationExt sets the handler for the RequestProbeInitiationExt method
func (protocol *Protocol) SetHandlerRequestProbeInitiationExt(handler func(err error, packet nex.PacketInterface, callID uint32, targetList *types.List[*types.String], stationToProbe *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.RequestProbeInitiationExt = handler
}

// SetHandlerReportNATTraversalResult sets the handler for the ReportNATTraversalResult method
func (protocol *Protocol) SetHandlerReportNATTraversalResult(handler func(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, rtt *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReportNATTraversalResult = handler
}

// SetHandlerReportNATProperties sets the handler for the ReportNATProperties method
func (protocol *Protocol) SetHandlerReportNATProperties(handler func(err error, packet nex.PacketInterface, callID uint32, natmapping *types.PrimitiveU32, natfiltering *types.PrimitiveU32, rtt *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReportNATProperties = handler
}

// SetHandlerGetRelaySignatureKey sets the handler for the GetRelaySignatureKey method
func (protocol *Protocol) SetHandlerGetRelaySignatureKey(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRelaySignatureKey = handler
}

// SetHandlerReportNATTraversalResultDetail sets the handler for the ReportNATTraversalResultDetail method
func (protocol *Protocol) SetHandlerReportNATTraversalResultDetail(handler func(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, detail *types.PrimitiveS32, rtt *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReportNATTraversalResultDetail = handler
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
	case MethodRequestProbeInitiation:
		protocol.handleRequestProbeInitiation(packet)
	case MethodInitiateProbe:
		protocol.handleInitiateProbe(packet)
	case MethodRequestProbeInitiationExt:
		protocol.handleRequestProbeInitiationExt(packet)
	case MethodReportNATTraversalResult:
		protocol.handleReportNATTraversalResult(packet)
	case MethodReportNATProperties:
		protocol.handleReportNATProperties(packet)
	case MethodGetRelaySignatureKey:
		protocol.handleGetRelaySignatureKey(packet)
	case MethodReportNATTraversalResultDetail:
		protocol.handleReportNATTraversalResultDetail(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported NATTraversal method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new NAT Traversal NEX protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
