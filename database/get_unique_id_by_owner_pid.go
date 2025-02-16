package database

func GetUniqueIDByOwnerPID(pid uint32) (uint32, error) {
	var uniqueID uint32
	err := Postgres.QueryRow(`SELECT unique_id FROM common_data WHERE owner_pid=$1`, pid).Scan(&uniqueID)
	if err != nil {
		return 0, err
	}

	return uniqueID, nil
}
