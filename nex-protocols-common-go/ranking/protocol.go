package ranking

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
)

type CommonProtocol struct {
	endpoint                                          nex.EndpointInterface
	protocol                                          ranking.Interface
	GetCommonData                                     func(uniqueID *types.PrimitiveU64) (*types.Buffer, error)
	UploadCommonData                                  func(pid *types.PID, uniqueID *types.PrimitiveU64, commonData *types.Buffer) error
	InsertRankingByPIDAndRankingScoreData             func(pid *types.PID, rankingScoreData *ranking_types.RankingScoreData, uniqueID *types.PrimitiveU64) error
	GetRankingsAndCountByCategoryAndRankingOrderParam func(category *types.PrimitiveU32, rankingOrderParam *ranking_types.RankingOrderParam) (*types.List[*ranking_types.RankingRankData], uint32, error)
	OnAfterGetCachedTopXRanking                       func(packet nex.PacketInterface, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam)
	OnAfterGetCachedTopXRankings                      func(packet nex.PacketInterface, categories *types.List[*types.PrimitiveU32], orderParams *types.List[*ranking_types.RankingOrderParam])
	OnAfterGetCommonData                              func(packet nex.PacketInterface, uniqueID *types.PrimitiveU64)
	OnAfterGetRanking                                 func(packet nex.PacketInterface, rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64, principalID *types.PID)
	OnAfterUploadCommonData                           func(packet nex.PacketInterface, commonData *types.Buffer, uniqueID *types.PrimitiveU64)
	OnAfterUploadScore                                func(packet nex.PacketInterface, scoreData *ranking_types.RankingScoreData, uniqueID *types.PrimitiveU64)
}

// NewCommonProtocol returns a new CommonProtocol
func NewCommonProtocol(protocol ranking.Interface) *CommonProtocol {
	commonProtocol := &CommonProtocol{
		endpoint: protocol.Endpoint(),
		protocol: protocol,
	}

	protocol.SetHandlerGetCachedTopXRanking(commonProtocol.getCachedTopXRanking)
	protocol.SetHandlerGetCachedTopXRankings(commonProtocol.getCachedTopXRankings)
	protocol.SetHandlerGetCommonData(commonProtocol.getCommonData)
	protocol.SetHandlerGetRanking(commonProtocol.getRanking)
	protocol.SetHandlerUploadCommonData(commonProtocol.uploadCommonData)
	protocol.SetHandlerUploadScore(commonProtocol.uploadScore)

	return commonProtocol
}
