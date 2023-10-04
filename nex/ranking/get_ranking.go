package nex_ranking

import (
	"github.com/PretendoNetwork/nex-go"
	ranking "github.com/PretendoNetwork/nex-protocols-go/ranking"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
	"github.com/PretendoNetwork/pokemon-gen6/database"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
)

func GetRanking(err error, client *nex.Client, callID uint32, rankingMode uint8, category uint32, orderParam *ranking_types.RankingOrderParam, uniqueID uint64, principalID uint32) {
	rmcResponse := nex.NewRMCResponse(ranking.ProtocolID, callID)

	if err != nil {
		globals.Logger.Error(err.Error())
		rmcResponse.SetError(nex.Errors.Ranking.Unknown)
	}

	rankDataListErr, rankDataList := database.GetRankingsByCategoryAndRankingOrderParam(category, orderParam)
	if rankDataListErr != nil {
		globals.Logger.Error(rankDataListErr.Error())
		rmcResponse.SetError(nex.Errors.Ranking.Unknown)
	}

	totalCountErr, totalCount := database.GetTotalRankingsByCategory(category)
	if totalCountErr != nil {
		globals.Logger.Error(totalCountErr.Error())
		rmcResponse.SetError(nex.Errors.Ranking.Unknown)
	}

	if totalCount == 0 || len(rankDataList) == 0 {
		rmcResponse.SetError(nex.Errors.Ranking.NotFound)
	}

	if rankDataListErr == nil && totalCountErr == nil && totalCount != 0 {
		pResult := ranking_types.NewRankingResult()

		pResult.RankDataList = rankDataList
		pResult.TotalCount = totalCount
		pResult.SinceTime = nex.NewDateTime(0x1f40420000) // * 2000-01-01T00:00:00.000Z, this is what the real server sends back

		rmcResponseStream := nex.NewStreamOut(globals.SecureServer)

		rmcResponseStream.WriteStructure(pResult)

		rmcResponseBody := rmcResponseStream.Bytes()

		rmcResponse.SetSuccess(ranking.MethodGetRanking, rmcResponseBody)
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
