// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the MatchMakingExt protocol
	ProtocolID = 0x32

	// MethodEndParticipation is the method ID for the method EndParticipation
	MethodEndParticipation = 0x1

	// MethodGetParticipants is the method ID for the method GetParticipants
	MethodGetParticipants = 0x2

	// MethodGetDetailedParticipants is the method ID for the method GetDetailedParticipants
	MethodGetDetailedParticipants = 0x3

	// MethodGetParticipantsURLs is the method ID for the method GetParticipantsURLs
	MethodGetParticipantsURLs = 0x4

	// MethodGetGatheringRelations is the method ID for the method GetGatheringRelations
	MethodGetGatheringRelations = 0x5

	// MethodDeleteFromDeletions is the method ID for the method DeleteFromDeletions
	MethodDeleteFromDeletions = 0x6
)

// Protocol handles the MatchMakingExt protocol
type Protocol struct {
	endpoint                nex.EndpointInterface
	EndParticipation        func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)
	GetParticipants         func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, bOnlyActive *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetDetailedParticipants func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, bOnlyActive *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetParticipantsURLs     func(err error, packet nex.PacketInterface, callID uint32, lstGatherings *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
	GetGatheringRelations   func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32, descr *types.String) (*nex.RMCMessage, *nex.Error)
	DeleteFromDeletions     func(err error, packet nex.PacketInterface, callID uint32, lstDeletions *types.List[*types.PrimitiveU32], pid *types.PID) (*nex.RMCMessage, *nex.Error)
	Patches                 nex.ServiceProtocol
	PatchedMethods          []uint32
}

// Interface implements the methods present on the Match Making Ext protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerEndParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, bOnlyActive *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetDetailedParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, bOnlyActive *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetParticipantsURLs(handler func(err error, packet nex.PacketInterface, callID uint32, lstGatherings *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetGatheringRelations(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32, descr *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteFromDeletions(handler func(err error, packet nex.PacketInterface, callID uint32, lstDeletions *types.List[*types.PrimitiveU32], pid *types.PID) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerEndParticipation sets the handler for the EndParticipation method
func (protocol *Protocol) SetHandlerEndParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.EndParticipation = handler
}

// SetHandlerGetParticipants sets the handler for the GetParticipants method
func (protocol *Protocol) SetHandlerGetParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, bOnlyActive *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetParticipants = handler
}

// SetHandlerGetDetailedParticipants sets the handler for the GetDetailedParticipants method
func (protocol *Protocol) SetHandlerGetDetailedParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, bOnlyActive *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetDetailedParticipants = handler
}

// SetHandlerGetParticipantsURLs sets the handler for the GetParticipantsURLs method
func (protocol *Protocol) SetHandlerGetParticipantsURLs(handler func(err error, packet nex.PacketInterface, callID uint32, lstGatherings *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetParticipantsURLs = handler
}

// SetHandlerGetGatheringRelations sets the handler for the GetGatheringRelations method
func (protocol *Protocol) SetHandlerGetGatheringRelations(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32, descr *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetGatheringRelations = handler
}

// SetHandlerDeleteFromDeletions sets the handler for the DeleteFromDeletions method
func (protocol *Protocol) SetHandlerDeleteFromDeletions(handler func(err error, packet nex.PacketInterface, callID uint32, lstDeletions *types.List[*types.PrimitiveU32], pid *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteFromDeletions = handler
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
	case MethodEndParticipation:
		protocol.handleEndParticipation(packet)
	case MethodGetParticipants:
		protocol.handleGetParticipants(packet)
	case MethodGetDetailedParticipants:
		protocol.handleGetDetailedParticipants(packet)
	case MethodGetParticipantsURLs:
		protocol.handleGetParticipantsURLs(packet)
	case MethodGetGatheringRelations:
		protocol.handleGetGatheringRelations(packet)
	case MethodDeleteFromDeletions:
		protocol.handleDeleteFromDeletions(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported MatchMakingExt method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Match Making Ext protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
