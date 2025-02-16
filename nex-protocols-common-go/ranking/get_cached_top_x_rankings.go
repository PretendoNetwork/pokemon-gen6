package ranking

import (
	"time"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
)

func (commonProtocol *CommonProtocol) getCachedTopXRankings(err error, packet nex.PacketInterface, callID uint32, categories *types.List[*types.PrimitiveU32], orderParams *types.List[*ranking_types.RankingOrderParam]) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.GetRankingsAndCountByCategoryAndRankingOrderParam == nil {
		common_globals.Logger.Warning("Ranking::GetCachedTopXRankings missing GetRankingsAndCountByCategoryAndRankingOrderParam!")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Ranking.InvalidArgument, "change_error")
	}

	// TODO - Is this true?
	if categories.Length() != orderParams.Length() {
		return nil, nex.NewError(nex.ResultCodes.Ranking.InvalidArgument, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	pResult := types.NewList[*ranking_types.RankingCachedResult]()

	pResult.Type = ranking_types.NewRankingCachedResult()

	for i, category := range categories.Slice() {
		orderParam, err := orderParams.Get(i)
		if err != nil {
			common_globals.Logger.Error(err.Error())
			return nil, nex.NewError(nex.ResultCodes.Ranking.InvalidArgument, "change_error")
		}

		rankDataList, totalCount, err := commonProtocol.GetRankingsAndCountByCategoryAndRankingOrderParam(category, orderParam)
		if err != nil {
			common_globals.Logger.Critical(err.Error())
			return nil, nex.NewError(nex.ResultCodes.Ranking.Unknown, "change_error")
		}

		if totalCount == 0 || rankDataList.Length() == 0 {
			return nil, nex.NewError(nex.ResultCodes.Ranking.NotFound, "change_error")
		}

		result := ranking_types.NewRankingCachedResult()

		result.RankingResult.RankDataList = rankDataList
		result.RankingResult.TotalCount = types.NewPrimitiveU32(totalCount)
		result.RankingResult.SinceTime = types.NewDateTime(0x1F40420000) // * 2000-01-01T00:00:00.000Z, this is what the real server sends back

		result.CreatedTime = types.NewDateTime(0).Now()
		// * The real server sends the "CreatedTime" + 5 minutes.
		// * It doesn't change, even on subsequent requests, until after the
		// * ExpiredTime has passed (seemingly what the "cached" means).
		// * Whether we need to replicate this idk, but in case, here's a note.
		result.ExpiredTime = types.NewDateTime(0).FromTimestamp(time.Now().UTC().Add(time.Minute * time.Duration(5)))
		// * This is the length Ultimate NES Remix uses
		// TODO - Does this matter? and are other games different?
		result.MaxLength = types.NewPrimitiveU8(10)

		pResult.Append(result)
	}

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pResult.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = ranking.ProtocolID
	rmcResponse.MethodID = ranking.MethodGetCachedTopXRankings
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterGetCachedTopXRankings != nil {
		go commonProtocol.OnAfterGetCachedTopXRankings(packet, categories, orderParams)
	}

	return rmcResponse, nil
}
