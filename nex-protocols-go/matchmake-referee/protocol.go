// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-referee/types"
)

const (
	// ProtocolID is the protocol ID for the Message Delivery protocol
	ProtocolID = 0x78

	// MethodStartRound is the method ID for the method StartRound
	MethodStartRound = 0x1

	// MethodGetStartRoundParam is the method ID for the method GetStartRoundParam
	MethodGetStartRoundParam = 0x2

	// MethodEndRound is the method ID for the method EndRound
	MethodEndRound = 0x3

	// MethodEndRoundWithoutReport is the method ID for the method EndRoundWithoutReport
	MethodEndRoundWithoutReport = 0x4

	// MethodGetRoundParticipants is the method ID for the method GetRoundParticipants
	MethodGetRoundParticipants = 0x5

	// MethodGetNotSummarizedRound is the method ID for the method GetNotSummarizedRound
	MethodGetNotSummarizedRound = 0x6

	// MethodGetRound is the method ID for the method GetRound
	MethodGetRound = 0x7

	// MethodGetStatsPrimary is the method ID for the method GetStatsPrimary
	MethodGetStatsPrimary = 0x8

	// MethodGetStatsPrimaries is the method ID for the method GetStatsPrimaries
	MethodGetStatsPrimaries = 0x9

	// MethodGetStatsAll is the method ID for the method GetStatsAll
	MethodGetStatsAll = 0xA

	// MethodCreateStats is the method ID for the method CreateStats
	MethodCreateStats = 0xB

	// MethodGetOrCreateStats is the method ID for the method GetOrCreateStats
	MethodGetOrCreateStats = 0xC

	// MethodResetStats is the method ID for the method ResetStats
	MethodResetStats = 0xD
)

// Protocol stores all the RMC method handlers for the Matchmake Referee protocol and listens for requests
type Protocol struct {
	endpoint              nex.EndpointInterface
	StartRound            func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStartRoundParam) (*nex.RMCMessage, *nex.Error)
	GetStartRoundParam    func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	EndRound              func(err error, packet nex.PacketInterface, callID uint32, endRoundParam *matchmake_referee_types.MatchmakeRefereeEndRoundParam) (*nex.RMCMessage, *nex.Error)
	EndRoundWithoutReport func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetRoundParticipants  func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetNotSummarizedRound func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetRound              func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetStatsPrimary       func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, *nex.Error)
	GetStatsPrimaries     func(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*matchmake_referee_types.MatchmakeRefereeStatsTarget]) (*nex.RMCMessage, *nex.Error)
	GetStatsAll           func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, *nex.Error)
	CreateStats           func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, *nex.Error)
	GetOrCreateStats      func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, *nex.Error)
	ResetStats            func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	Patches               nex.ServiceProtocol
	PatchedMethods        []uint32
}

// Interface implements the methods present on the Matchmake Referee protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerStartRound(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStartRoundParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetStartRoundParam(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerEndRound(handler func(err error, packet nex.PacketInterface, callID uint32, endRoundParam *matchmake_referee_types.MatchmakeRefereeEndRoundParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerEndRoundWithoutReport(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRoundParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNotSummarizedRound(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRound(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetStatsPrimary(handler func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetStatsPrimaries(handler func(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*matchmake_referee_types.MatchmakeRefereeStatsTarget]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetStatsAll(handler func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, *nex.Error))
	SetHandlerCreateStats(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetOrCreateStats(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerResetStats(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerStartRound sets the handler for the StartRound method
func (protocol *Protocol) SetHandlerStartRound(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStartRoundParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.StartRound = handler
}

// SetHandlerGetStartRoundParam sets the handler for the GetStartRoundParam method
func (protocol *Protocol) SetHandlerGetStartRoundParam(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetStartRoundParam = handler
}

// SetHandlerEndRound sets the handler for the EndRound method
func (protocol *Protocol) SetHandlerEndRound(handler func(err error, packet nex.PacketInterface, callID uint32, endRoundParam *matchmake_referee_types.MatchmakeRefereeEndRoundParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.EndRound = handler
}

// SetHandlerEndRoundWithoutReport sets the handler for the EndRoundWithoutReport method
func (protocol *Protocol) SetHandlerEndRoundWithoutReport(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.EndRoundWithoutReport = handler
}

// SetHandlerGetRoundParticipants sets the handler for the GetRoundParticipants method
func (protocol *Protocol) SetHandlerGetRoundParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRoundParticipants = handler
}

// SetHandlerGetNotSummarizedRound sets the handler for the GetNotSummarizedRound method
func (protocol *Protocol) SetHandlerGetNotSummarizedRound(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNotSummarizedRound = handler
}

// SetHandlerGetRound sets the handler for the GetRound method
func (protocol *Protocol) SetHandlerGetRound(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRound = handler
}

// SetHandlerGetStatsPrimary sets the handler for the GetStatsPrimary method
func (protocol *Protocol) SetHandlerGetStatsPrimary(handler func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetStatsPrimary = handler
}

// SetHandlerGetStatsPrimaries sets the handler for the GetStatsPrimaries method
func (protocol *Protocol) SetHandlerGetStatsPrimaries(handler func(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*matchmake_referee_types.MatchmakeRefereeStatsTarget]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetStatsPrimaries = handler
}

// SetHandlerGetStatsAll sets the handler for the GetStatsAll method
func (protocol *Protocol) SetHandlerGetStatsAll(handler func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetStatsAll = handler
}

// SetHandlerCreateStats sets the handler for the CreateStats method
func (protocol *Protocol) SetHandlerCreateStats(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.CreateStats = handler
}

// SetHandlerGetOrCreateStats sets the handler for the GetOrCreateStats method
func (protocol *Protocol) SetHandlerGetOrCreateStats(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetOrCreateStats = handler
}

// SetHandlerResetStats sets the handler for the ResetStats method
func (protocol *Protocol) SetHandlerResetStats(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.ResetStats = handler
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
	case MethodStartRound:
		protocol.handleStartRound(packet)
	case MethodGetStartRoundParam:
		protocol.handleGetStartRoundParam(packet)
	case MethodEndRound:
		protocol.handleEndRound(packet)
	case MethodEndRoundWithoutReport:
		protocol.handleEndRoundWithoutReport(packet)
	case MethodGetRoundParticipants:
		protocol.handleGetRoundParticipants(packet)
	case MethodGetNotSummarizedRound:
		protocol.handleGetNotSummarizedRound(packet)
	case MethodGetRound:
		protocol.handleGetRound(packet)
	case MethodGetStatsPrimary:
		protocol.handleGetStatsPrimary(packet)
	case MethodGetStatsPrimaries:
		protocol.handleGetStatsPrimaries(packet)
	case MethodGetStatsAll:
		protocol.handleGetStatsAll(packet)
	case MethodCreateStats:
		protocol.handleCreateStats(packet)
	case MethodGetOrCreateStats:
		protocol.handleGetOrCreateStats(packet)
	case MethodResetStats:
		protocol.handleResetStats(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported MatchmakeReferee method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Matchmake Referee protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
