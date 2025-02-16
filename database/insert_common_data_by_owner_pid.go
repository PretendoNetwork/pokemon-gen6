package database

func InsertCommonDataByOwnerPID(pid uint32) (uint32, error) {
	var uniqueID uint32
	err := Postgres.QueryRow(`INSERT INTO common_data (owner_pid) VALUES ($1) RETURNING unique_id`, pid).Scan(&uniqueID)
	if err != nil {
		return 0, err
	}

	return uniqueID, nil
}
