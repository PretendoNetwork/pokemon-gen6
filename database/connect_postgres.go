package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"

	"github.com/PretendoNetwork/pokemon-gen6/globals"
)

var Postgres *sql.DB

func ConnectPostgres() {
	var err error

	Postgres, err = sql.Open("postgres", os.Getenv("PN_POKEGEN6_POSTGRES_URI"))
	if err != nil {
		globals.Logger.Critical(err.Error())
	}

	globals.Logger.Success("Connected to Postgres!")
}

func InitUtilityDatabase() {
	_, err := Postgres.Exec(`CREATE SCHEMA IF NOT EXISTS utility`)
	if err != nil {
		globals.Logger.Error(err.Error())
		return
	}

	_, err = Postgres.Exec(`CREATE TABLE IF NOT EXISTS utility.unique_ids (
		unique_id numeric(20) PRIMARY KEY,
		password numeric(20) DEFAULT 0,
		associated_pid numeric(20) DEFAULT 0
	)`)
	if err != nil {
		globals.Logger.Error(err.Error())
		return
	}
}
