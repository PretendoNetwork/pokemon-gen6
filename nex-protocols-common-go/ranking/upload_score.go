package ranking

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
)

func (commonProtocol *CommonProtocol) uploadScore(err error, packet nex.PacketInterface, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error) {
	if commonProtocol.InsertRankingByPIDAndRankingScoreData == nil {
		common_globals.Logger.Warning("Ranking::UploadScore missing InsertRankingByPIDAndRankingScoreData!")
		return nil, nex.NewError(nex.ResultCodes.Core.NotImplemented, "change_error")
	}

	if err != nil {
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Ranking.InvalidArgument, "change_error")
	}

	connection := packet.Sender()
	endpoint := connection.Endpoint()

	err = commonProtocol.InsertRankingByPIDAndRankingScoreData(connection.PID(), scoreData, uniqueID)
	if err != nil {
		common_globals.Logger.Critical(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Ranking.Unknown, "change_error")
	}

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = ranking.ProtocolID
	rmcResponse.MethodID = ranking.MethodUploadScore
	rmcResponse.CallID = callID

	if commonProtocol.OnAfterUploadScore != nil {
		go commonProtocol.OnAfterUploadScore(packet, scoreData, uniqueID)
	}

	return rmcResponse, nil
}
