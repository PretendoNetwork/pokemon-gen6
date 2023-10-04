package database

import "github.com/PretendoNetwork/pokemon-gen6/globals"

func initPostgres() {
	var err error

	_, err = Postgres.Exec(`CREATE TABLE IF NOT EXISTS rankings (
		id bigserial PRIMARY KEY,
		owner_pid integer,
		category integer,
		score integer,
		order_by integer,
		update_mode integer,
		groups integer[],
		param bigint,
		common_data bytea,
		created_at bigint
	)`)
	if err != nil {
		globals.Logger.Critical(err.Error())
		return
	}

	globals.Logger.Success("Postgres tables created")
}
