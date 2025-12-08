package database_utility

import (
	"crypto/rand"
	"encoding/binary"
	"strconv"

	"github.com/PretendoNetwork/pokemon-gen6/database"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
)

func GenerateNEXUniqueID() uint64 {
	var uniqueID uint64

	err := binary.Read(rand.Reader, binary.BigEndian, &uniqueID)
	if err != nil {
		globals.Logger.Error(err.Error())
	}

	_, err = database.Postgres.Exec(`INSERT INTO utility.unique_ids (
		unique_id
	) VALUES (
		$1
	)`, strconv.FormatUint(uniqueID, 10))
	if err != nil {
		globals.Logger.Error(err.Error())
	}

	return uniqueID
}
