// Package protocol implements the RankingSplatoon protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetCompetitionRankingScore(packet nex.PacketInterface) {
	if protocol.GetCompetitionRankingScore == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "RankingSplatoon::GetCompetitionRankingScore not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	globals.Logger.Warning("RankingSplatoon::GetCompetitionRankingScore STUBBED")

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.GetCompetitionRankingScore(nil, packet, callID, packet.Payload())
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
