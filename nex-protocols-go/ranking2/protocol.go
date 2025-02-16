// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking2/types"
)

const (
	// ProtocolID is the protocol ID for the Ranking2 protocol
	ProtocolID = 0x7A

	// MethodPutScore is the method ID for the method PutScore
	MethodPutScore = 0x1

	// MethodGetCommonData is the method ID for the method GetCommonData
	MethodGetCommonData = 0x2

	// MethodPutCommonData is the method ID for the method PutCommonData
	MethodPutCommonData = 0x3

	// MethodDeleteCommonData is the method ID for the method DeleteCommonData
	MethodDeleteCommonData = 0x4

	// MethodGetRanking is the method ID for the method GetRanking
	MethodGetRanking = 0x5

	// MethodGetRankingByPrincipalID is the method ID for the method GetRankingByPrincipalID
	MethodGetRankingByPrincipalID = 0x6

	// MethodGetCategorySetting is the method ID for the method GetCategorySetting
	MethodGetCategorySetting = 0x7

	// MethodGetRankingChart is the method ID for the method GetRankingChart
	MethodGetRankingChart = 0x8

	// MethodGetRankingCharts is the method ID for the method GetRankingCharts
	MethodGetRankingCharts = 0x9

	// MethodGetEstimateScoreRank is the method ID for the method GetEstimateScoreRank
	MethodGetEstimateScoreRank = 0xA
)

// Protocol stores all the RMC method handlers for the Ranking2 protocol and listens for requests
type Protocol struct {
	endpoint                nex.EndpointInterface
	PutScore                func(err error, packet nex.PacketInterface, callID uint32, scoreDataList *types.List[*ranking2_types.Ranking2ScoreData], nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetCommonData           func(err error, packet nex.PacketInterface, callID uint32, optionFlags *types.PrimitiveU32, principalID *types.PID, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	PutCommonData           func(err error, packet nex.PacketInterface, callID uint32, commonData *ranking2_types.Ranking2CommonData, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	DeleteCommonData        func(err error, packet nex.PacketInterface, callID uint32, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetRanking              func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam) (*nex.RMCMessage, *nex.Error)
	GetRankingByPrincipalID func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam, principalIDList *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)
	GetCategorySetting      func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetRankingChart         func(err error, packet nex.PacketInterface, callID uint32, info *ranking2_types.Ranking2ChartInfoInput) (*nex.RMCMessage, *nex.Error)
	GetRankingCharts        func(err error, packet nex.PacketInterface, callID uint32, infoArray *types.List[*ranking2_types.Ranking2ChartInfoInput]) (*nex.RMCMessage, *nex.Error)
	GetEstimateScoreRank    func(err error, packet nex.PacketInterface, callID uint32, input *ranking2_types.Ranking2EstimateScoreRankInput) (*nex.RMCMessage, *nex.Error)
	Patches                 nex.ServiceProtocol
	PatchedMethods          []uint32
}

// Interface implements the methods present on the Ranking2 protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerPutScore(handler func(err error, packet nex.PacketInterface, callID uint32, scoreDataList *types.List[*ranking2_types.Ranking2ScoreData], nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, optionFlags *types.PrimitiveU32, principalID *types.PID, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerPutCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, commonData *ranking2_types.Ranking2CommonData, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRanking(handler func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRankingByPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam, principalIDList *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetCategorySetting(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRankingChart(handler func(err error, packet nex.PacketInterface, callID uint32, info *ranking2_types.Ranking2ChartInfoInput) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRankingCharts(handler func(err error, packet nex.PacketInterface, callID uint32, infoArray *types.List[*ranking2_types.Ranking2ChartInfoInput]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetEstimateScoreRank(handler func(err error, packet nex.PacketInterface, callID uint32, input *ranking2_types.Ranking2EstimateScoreRankInput) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerPutScore sets the handler for the PutScore method
func (protocol *Protocol) SetHandlerPutScore(handler func(err error, packet nex.PacketInterface, callID uint32, scoreDataList *types.List[*ranking2_types.Ranking2ScoreData], nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.PutScore = handler
}

// SetHandlerGetCommonData sets the handler for the GetCommonData method
func (protocol *Protocol) SetHandlerGetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, optionFlags *types.PrimitiveU32, principalID *types.PID, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetCommonData = handler
}

// SetHandlerPutCommonData sets the handler for the PutCommonData method
func (protocol *Protocol) SetHandlerPutCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, commonData *ranking2_types.Ranking2CommonData, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.PutCommonData = handler
}

// SetHandlerDeleteCommonData sets the handler for the DeleteCommonData method
func (protocol *Protocol) SetHandlerDeleteCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteCommonData = handler
}

// SetHandlerGetRanking sets the handler for the GetRanking method
func (protocol *Protocol) SetHandlerGetRanking(handler func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRanking = handler
}

// SetHandlerGetRankingByPrincipalID sets the handler for the GetRankingByPrincipalID method
func (protocol *Protocol) SetHandlerGetRankingByPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam, principalIDList *types.List[*types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRankingByPrincipalID = handler
}

// SetHandlerGetCategorySetting sets the handler for the GetCategorySetting method
func (protocol *Protocol) SetHandlerGetCategorySetting(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetCategorySetting = handler
}

// SetHandlerGetRankingChart sets the handler for the GetRankingChart method
func (protocol *Protocol) SetHandlerGetRankingChart(handler func(err error, packet nex.PacketInterface, callID uint32, info *ranking2_types.Ranking2ChartInfoInput) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRankingChart = handler
}

// SetHandlerGetRankingCharts sets the handler for the GetRankingCharts method
func (protocol *Protocol) SetHandlerGetRankingCharts(handler func(err error, packet nex.PacketInterface, callID uint32, infoArray *types.List[*ranking2_types.Ranking2ChartInfoInput]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRankingCharts = handler
}

// SetHandlerGetEstimateScoreRank sets the handler for the GetEstimateScoreRank method
func (protocol *Protocol) SetHandlerGetEstimateScoreRank(handler func(err error, packet nex.PacketInterface, callID uint32, input *ranking2_types.Ranking2EstimateScoreRankInput) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetEstimateScoreRank = handler
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
	case MethodPutScore:
		protocol.handlePutScore(packet)
	case MethodGetCommonData:
		protocol.handleGetCommonData(packet)
	case MethodPutCommonData:
		protocol.handlePutCommonData(packet)
	case MethodDeleteCommonData:
		protocol.handleDeleteCommonData(packet)
	case MethodGetRanking:
		protocol.handleGetRanking(packet)
	case MethodGetRankingByPrincipalID:
		protocol.handleGetRankingByPrincipalID(packet)
	case MethodGetCategorySetting:
		protocol.handleGetCategorySetting(packet)
	case MethodGetRankingChart:
		protocol.handleGetRankingChart(packet)
	case MethodGetRankingCharts:
		protocol.handleGetRankingCharts(packet)
	case MethodGetEstimateScoreRank:
		protocol.handleGetEstimateScoreRank(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Ranking2 method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Ranking2 protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
