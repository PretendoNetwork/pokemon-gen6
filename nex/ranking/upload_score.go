package nex_ranking

import (
	"github.com/PretendoNetwork/nex-go"
	ranking "github.com/PretendoNetwork/nex-protocols-go/ranking"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
	"github.com/PretendoNetwork/pokemon-gen6/database"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
)

func UploadScore(err error, client *nex.Client, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID uint64) {
	rmcResponse := nex.NewRMCResponse(ranking.ProtocolID, callID)

	if err != nil {
		globals.Logger.Error(err.Error())
		rmcResponse.SetError(nex.Errors.Ranking.Unknown)
	}

	insertErr := database.InsertRankingByPIDAndRankingScoreData(client.PID(), scoreData)
	if insertErr != nil {
		globals.Logger.Error(insertErr.Error())
		rmcResponse.SetError(nex.Errors.Ranking.Unknown)
	} else {
		rmcResponse.SetSuccess(ranking.MethodUploadScore, nil)
	}

	rmcResponseBytes := rmcResponse.Bytes()

	responsePacket, _ := nex.NewPacketV1(client, nil)

	responsePacket.SetVersion(1)
	responsePacket.SetSource(0xA1)
	responsePacket.SetDestination(0xAF)
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	globals.SecureServer.Send(responsePacket)
}
