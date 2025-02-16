// Package protocol implements the Screening protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Screening protocol
	ProtocolID = 0x7C

	// MethodReportDataStoreContent is the method ID for the method ReportDataStoreContent
	MethodReportDataStoreContent = 0x1

	// MethodReportUser is the method ID for the method ReportUser
	MethodReportUser = 0x2
)

// Protocol handles the Screening protocol
type Protocol struct {
	endpoint               nex.EndpointInterface
	ReportDataStoreContent func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) // TODO - Unknown request/response format
	ReportUser             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) // TODO - Unknown request/response format
	Patches                nex.ServiceProtocol
	PatchedMethods         []uint32
}

// Interface implements the methods present on the Screening protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerReportDataStoreContent(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
	SetHandlerReportUser(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerReportDataStoreContent sets the handler for the ReportDataStoreContent method
func (protocol *Protocol) SetHandlerReportDataStoreContent(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReportDataStoreContent = handler
}

// SetHandlerReportUser sets the handler for the ReportUser method
func (protocol *Protocol) SetHandlerReportUser(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReportUser = handler
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
	case MethodReportDataStoreContent:
		protocol.handleReportDataStoreContent(packet)
	case MethodReportUser:
		protocol.handleReportUser(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Screening method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Screening protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
