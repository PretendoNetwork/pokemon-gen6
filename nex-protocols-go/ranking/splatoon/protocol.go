// Package protocol implements the RankingSplatoon protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Ranking (Splatoon) protocol
	ProtocolID = 0x70

	// MethodGetCompetitionRankingScore is the method ID for the GetCompetitionRankingScore method
	MethodGetCompetitionRankingScore = 0x10

	// MethodGetcompetitionRankingScoreByPeriodList is the method ID for the GetcompetitionRankingScoreByPeriodList method
	MethodGetcompetitionRankingScoreByPeriodList = 0x11

	// MethodUploadCompetitionRankingScore is the method ID for the UploadCompetitionRankingScore method
	MethodUploadCompetitionRankingScore = 0x12

	// MethodDeleteCompetitionRankingScore is the method ID for the DeleteCompetitionRankingScore method
	MethodDeleteCompetitionRankingScore = 0x13
)

var patchedMethods = []uint32{
	MethodGetCompetitionRankingScore,
	MethodGetcompetitionRankingScoreByPeriodList,
	MethodUploadCompetitionRankingScore,
	MethodDeleteCompetitionRankingScore,
}

type rankingProtocol = ranking.Protocol

// Protocol stores all the RMC method handlers for the Ranking (Splatoon) protocol and listens for requests
// Embeds the Ranking protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	rankingProtocol
	GetCompetitionRankingScore             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	GetcompetitionRankingScoreByPeriodList func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	UploadCompetitionRankingScore          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	DeleteCompetitionRankingScore          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.rankingProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodGetCompetitionRankingScore:
		protocol.handleGetCompetitionRankingScore(packet)
	case MethodGetcompetitionRankingScoreByPeriodList:
		protocol.handleGetcompetitionRankingScoreByPeriodList(packet)
	case MethodUploadCompetitionRankingScore:
		protocol.handleUploadCompetitionRankingScore(packet)
	case MethodDeleteCompetitionRankingScore:
		protocol.handleDeleteCompetitionRankingScore(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Ranking (Splatoon) method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new RankingSplatoon protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.rankingProtocol.SetEndpoint(endpoint)

	return protocol
}
