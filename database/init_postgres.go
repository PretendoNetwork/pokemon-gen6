package database

import "github.com/PretendoNetwork/pokemon-gen6/globals"

func initPostgres() {
	var err error

	_, err = Postgres.Exec(`CREATE TABLE IF NOT EXISTS common_data (
		unique_id serial PRIMARY KEY,
		owner_pid integer,
		common_data bytea
	)`)
	if err != nil {
		globals.Logger.Critical(err.Error())
		return
	}

	globals.Logger.Success("Postgres tables created")
}
