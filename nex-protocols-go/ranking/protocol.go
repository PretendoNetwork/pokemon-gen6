// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
)

const (
	// ProtocolID is the protocol ID for the Ranking protocol
	ProtocolID = 0x70

	// MethodUploadScore is the method ID for method UploadScore
	MethodUploadScore = 0x1

	// MethodDeleteScore is the method ID for method DeleteScore
	MethodDeleteScore = 0x2

	// MethodDeleteAllScores is the method ID for method DeleteAllScores
	MethodDeleteAllScores = 0x3

	// MethodUploadCommonData is the method ID for method UploadCommonData
	MethodUploadCommonData = 0x4

	// MethodDeleteCommonData is the method ID for method DeleteCommonData
	MethodDeleteCommonData = 0x5

	// MethodGetCommonData is the method ID for method GetCommonData
	MethodGetCommonData = 0x6

	// MethodChangeAttributes is the method ID for method ChangeAttributes
	MethodChangeAttributes = 0x7

	// MethodChangeAllAttributes is the method ID for method ChangeAllAttributes
	MethodChangeAllAttributes = 0x8

	// MethodGetRanking is the method ID for method GetRanking
	MethodGetRanking = 0x9

	// MethodGetApproxOrder is the method ID for method GetApproxOrder
	MethodGetApproxOrder = 0xA

	// MethodGetStats is the method ID for method GetStats
	MethodGetStats = 0xB

	// MethodGetRankingByPIDList is the method ID for method GetRankingByPIDList
	MethodGetRankingByPIDList = 0xC

	// MethodGetRankingByUniqueIDList is the method ID for method GetRankingByUniqueIDList
	MethodGetRankingByUniqueIDList = 0xD

	// MethodGetCachedTopXRanking is the method ID for method GetCachedTopXRanking
	MethodGetCachedTopXRanking = 0xE

	// MethodGetCachedTopXRankings is the method ID for method GetCachedTopXRankings
	MethodGetCachedTopXRankings = 0xF
)

// Protocol stores all the RMC method handlers for the Ranking protocol and listens for requests
type Protocol struct {
	endpoint                 nex.EndpointInterface
	UploadScore              func(err error, packet nex.PacketInterface, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	DeleteScore              func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	DeleteAllScores          func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	UploadCommonData         func(err error, packet nex.PacketInterface, callID uint32, commonData *types.Buffer, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	DeleteCommonData         func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetCommonData            func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	ChangeAttributes         func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	ChangeAllAttributes      func(err error, packet nex.PacketInterface, callID uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetRanking               func(err error, packet nex.PacketInterface, callID uint32, rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64, principalID *types.PID) (*nex.RMCMessage, *nex.Error)
	GetApproxOrder           func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, score *types.PrimitiveU32, uniqueID *types.PrimitiveU64, principalID *types.PID) (*nex.RMCMessage, *nex.Error)
	GetStats                 func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, flags *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	GetRankingByPIDList      func(err error, packet nex.PacketInterface, callID uint32, principalIDList *types.List[*types.PID], rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetRankingByUniqueIDList func(err error, packet nex.PacketInterface, callID uint32, nexUniqueIDList *types.List[*types.PrimitiveU64], rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
	GetCachedTopXRanking     func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam) (*nex.RMCMessage, *nex.Error)
	GetCachedTopXRankings    func(err error, packet nex.PacketInterface, callID uint32, categories *types.List[*types.PrimitiveU32], orderParams *types.List[*ranking_types.RankingOrderParam]) (*nex.RMCMessage, *nex.Error)
	Patches                  nex.ServiceProtocol
	PatchedMethods           []uint32
}

// Interface implements the methods present on the Ranking protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerUploadScore(handler func(err error, packet nex.PacketInterface, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteScore(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteAllScores(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerUploadCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, commonData *types.Buffer, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerChangeAttributes(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerChangeAllAttributes(handler func(err error, packet nex.PacketInterface, callID uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRanking(handler func(err error, packet nex.PacketInterface, callID uint32, rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64, principalID *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetApproxOrder(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, score *types.PrimitiveU32, uniqueID *types.PrimitiveU64, principalID *types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetStats(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, flags *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRankingByPIDList(handler func(err error, packet nex.PacketInterface, callID uint32, principalIDList *types.List[*types.PID], rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRankingByUniqueIDList(handler func(err error, packet nex.PacketInterface, callID uint32, nexUniqueIDList *types.List[*types.PrimitiveU64], rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetCachedTopXRanking(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetCachedTopXRankings(handler func(err error, packet nex.PacketInterface, callID uint32, categories *types.List[*types.PrimitiveU32], orderParams *types.List[*ranking_types.RankingOrderParam]) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerUploadScore sets the handler for the UploadScore method
func (protocol *Protocol) SetHandlerUploadScore(handler func(err error, packet nex.PacketInterface, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.UploadScore = handler
}

// SetHandlerDeleteScore sets the handler for the DeleteScore method
func (protocol *Protocol) SetHandlerDeleteScore(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteScore = handler
}

// SetHandlerDeleteAllScores sets the handler for the DeleteAllScores method
func (protocol *Protocol) SetHandlerDeleteAllScores(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteAllScores = handler
}

// SetHandlerUploadCommonData sets the handler for the UploadCommonData method
func (protocol *Protocol) SetHandlerUploadCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, commonData *types.Buffer, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.UploadCommonData = handler
}

// SetHandlerDeleteCommonData sets the handler for the DeleteCommonData method
func (protocol *Protocol) SetHandlerDeleteCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteCommonData = handler
}

// SetHandlerGetCommonData sets the handler for the GetCommonData method
func (protocol *Protocol) SetHandlerGetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetCommonData = handler
}

// SetHandlerChangeAttributes sets the handler for the ChangeAttributes method
func (protocol *Protocol) SetHandlerChangeAttributes(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.ChangeAttributes = handler
}

// SetHandlerChangeAllAttributes sets the handler for the ChangeAllAttributes method
func (protocol *Protocol) SetHandlerChangeAllAttributes(handler func(err error, packet nex.PacketInterface, callID uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.ChangeAllAttributes = handler
}

// SetHandlerGetRanking sets the handler for the GetRanking method
func (protocol *Protocol) SetHandlerGetRanking(handler func(err error, packet nex.PacketInterface, callID uint32, rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64, principalID *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRanking = handler
}

// SetHandlerGetApproxOrder sets the handler for the GetApproxOrder method
func (protocol *Protocol) SetHandlerGetApproxOrder(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, score *types.PrimitiveU32, uniqueID *types.PrimitiveU64, principalID *types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetApproxOrder = handler
}

// SetHandlerGetStats sets the handler for the GetStats method
func (protocol *Protocol) SetHandlerGetStats(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, flags *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetStats = handler
}

// SetHandlerGetRankingByPIDList sets the handler for the GetRankingByPIDList method
func (protocol *Protocol) SetHandlerGetRankingByPIDList(handler func(err error, packet nex.PacketInterface, callID uint32, principalIDList *types.List[*types.PID], rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRankingByPIDList = handler
}

// SetHandlerGetRankingByUniqueIDList sets the handler for the GetRankingByUniqueIDList method
func (protocol *Protocol) SetHandlerGetRankingByUniqueIDList(handler func(err error, packet nex.PacketInterface, callID uint32, nexUniqueIDList *types.List[*types.PrimitiveU64], rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRankingByUniqueIDList = handler
}

// SetHandlerGetCachedTopXRanking sets the handler for the GetCachedTopXRanking method
func (protocol *Protocol) SetHandlerGetCachedTopXRanking(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetCachedTopXRanking = handler
}

// SetHandlerGetCachedTopXRankings sets the handler for the GetCachedTopXRankings method
func (protocol *Protocol) SetHandlerGetCachedTopXRankings(handler func(err error, packet nex.PacketInterface, callID uint32, categories *types.List[*types.PrimitiveU32], orderParams *types.List[*ranking_types.RankingOrderParam]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetCachedTopXRankings = handler
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
	case MethodUploadScore:
		protocol.handleUploadScore(packet)
	case MethodDeleteScore:
		protocol.handleDeleteScore(packet)
	case MethodDeleteAllScores:
		protocol.handleDeleteAllScores(packet)
	case MethodUploadCommonData:
		protocol.handleUploadCommonData(packet)
	case MethodDeleteCommonData:
		protocol.handleDeleteCommonData(packet)
	case MethodGetCommonData:
		protocol.handleGetCommonData(packet)
	case MethodChangeAttributes:
		protocol.handleChangeAttributes(packet)
	case MethodChangeAllAttributes:
		protocol.handleChangeAllAttributes(packet)
	case MethodGetRanking:
		protocol.handleGetRanking(packet)
	case MethodGetApproxOrder:
		protocol.handleGetApproxOrder(packet)
	case MethodGetStats:
		protocol.handleGetStats(packet)
	case MethodGetRankingByPIDList:
		protocol.handleGetRankingByPIDList(packet)
	case MethodGetRankingByUniqueIDList:
		protocol.handleGetRankingByUniqueIDList(packet)
	case MethodGetCachedTopXRanking:
		protocol.handleGetCachedTopXRanking(packet)
	case MethodGetCachedTopXRankings:
		protocol.handleGetCachedTopXRankings(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Ranking method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Ranking protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
