package database

import (
	"time"

	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
	"github.com/lib/pq"
)

func InsertRankingByPIDAndRankingScoreData(pid uint32, rankingScoreData *ranking_types.RankingScoreData) error {
	now := time.Now().UnixNano()

	_, err := Postgres.Exec(`
		INSERT INTO rankings (
			owner_pid,
			category,
			score,
			order_by,
			update_mode,
			groups,
			param,
			common_data,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		pid,
		rankingScoreData.Category,
		rankingScoreData.Score,
		rankingScoreData.OrderBy,
		rankingScoreData.UpdateMode,
		pq.Array(rankingScoreData.Groups),
		rankingScoreData.Param,
		make([]byte, 0),
		now,
	)

	return err
}
