// Package protocol implements the RankingMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
	ranking_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/mario-kart-8/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Ranking (Mario Kart 8) protocol
	ProtocolID = 0x70

	// MethodGetCompetitionRankingScore is the method ID for the GetCompetitionRankingScore method
	MethodGetCompetitionRankingScore = 0xE

	// MethodUploadCompetitionRankingScore is the method ID for the UploadCompetitionRankingScore method
	MethodUploadCompetitionRankingScore = 0xF

	// MethodGetCompetitionInfo is the method ID for the GetCompetitionInfo method
	MethodGetCompetitionInfo = 0x10
)

var patchedMethods = []uint32{
	MethodGetCompetitionRankingScore,
	MethodUploadCompetitionRankingScore,
	MethodGetCompetitionInfo,
}

type rankingProtocol = ranking.Protocol

// Protocol stores all the RMC method handlers for the Ranking (Mario Kart 8) protocol and listens for requests
// Embeds the Ranking protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	rankingProtocol
	GetCompetitionRankingScore    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	UploadCompetitionRankingScore func(err error, packet nex.PacketInterface, callID uint32, param *ranking_mario_kart8_types.CompetitionRankingUploadScoreParam) (*nex.RMCMessage, *nex.Error)
	GetCompetitionInfo            func(err error, packet nex.PacketInterface, callID uint32, param *ranking_mario_kart8_types.CompetitionRankingInfoGetParam) (*nex.RMCMessage, *nex.Error)
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
	case MethodUploadCompetitionRankingScore:
		protocol.handleUploadCompetitionRankingScore(packet)
	case MethodGetCompetitionInfo:
		protocol.handleGetCompetitionInfo(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Ranking (Mario Kart 8) method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new RankingMarioKart8 protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.rankingProtocol.SetEndpoint(endpoint)

	return protocol
}
