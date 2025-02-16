package ranking

import (
	"time"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
)

func (commonProtocol *CommonProtocol) getCachedTopXRanking(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.GetRankingsAndCountByCategoryAndRankingOrderParam == nil {
		common_globals.Logger.Warning("Ranking::GetCachedTopXRanking missing GetRankingsAndCountByCategoryAndRankingOrderParam!")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

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

	pResult := ranking_types.NewRankingCachedResult()

	pResult.RankingResult.RankDataList = rankDataList
	pResult.RankingResult.TotalCount = types.NewPrimitiveU32(totalCount)
	pResult.RankingResult.SinceTime = types.NewDateTime(0x1F40420000) // * 2000-01-01T00:00:00.000Z, this is what the real server sends back

	pResult.CreatedTime = types.NewDateTime(0).Now()
	// * The real server sends the "CreatedTime" + 5 minutes.
	// * It doesn't change, even on subsequent requests, until after the
	// * ExpiredTime has passed (seemingly what the "cached" means).
	// * Whether we need to replicate this idk, but in case, here's a note.
	pResult.ExpiredTime = types.NewDateTime(0).FromTimestamp(time.Now().UTC().Add(time.Minute * time.Duration(5)))
	// * This is the length Ultimate NES Remix uses
	// TODO - Does this matter? and are other games different?
	pResult.MaxLength = types.NewPrimitiveU8(10)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pResult.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = ranking.ProtocolID
	rmcResponse.MethodID = ranking.MethodGetCachedTopXRanking
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterGetCachedTopXRanking != nil {
		go commonProtocol.OnAfterGetCachedTopXRanking(packet, category, orderParam)
	}

	return rmcResponse, nil
}
