package database

import (
	"database/sql"

	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	"github.com/PretendoNetwork/pokemon-gen6/types"
)

func GetRankingsByCategoryAndRankingOrderParam(category uint32, rankingOrderParam *ranking_types.RankingOrderParam) (error, []*ranking_types.RankingRankData) {
	rankings := make([]*ranking_types.RankingRankData, 0, rankingOrderParam.Length)

	rows, err := Postgres.Query(`
		SELECT
		owner_pid,
		score,
		groups,
		param,
		common_data
		FROM rankings WHERE category=$1 ORDER BY score DESC LIMIT $2 OFFSET $3`,
		category,
		rankingOrderParam.Length,
		rankingOrderParam.Offset,
	)
	if err != nil {
		return err, rankings
	}

	row := 1
	for rows.Next() {
		rankingRankData := ranking_types.NewRankingRankData()
		rankingRankData.UniqueID = 0
		rankingRankData.Order = uint32(row)
		rankingRankData.Category = category

		// * A custom type is needed because
		// * Postgres doesn't support scanning
		// * uint8 slices by default
		var groups types.PQUInt8Array

		err := rows.Scan(
			&rankingRankData.PrincipalID,
			&rankingRankData.Score,
			&groups,
			&rankingRankData.Param,
			&rankingRankData.CommonData,
		)

		if err != nil && err != sql.ErrNoRows {
			globals.Logger.Critical(err.Error())
		}

		if err == nil {
			rankingRankData.Groups = groups.Value
			rankings = append(rankings, rankingRankData)

			row += 1
		}
	}

	return nil, rankings
}
