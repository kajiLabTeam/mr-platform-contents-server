package model

import (
	"database/sql"

	"github.com/kajiLabTeam/mr-platform-contents-server/common"
)

func IsExistLayer(layerId string) (bool, error) {
	row := db.QueryRow("SELECT id FROM layer WHERE id = $1", layerId)

	var space common.PublicSpace
	if err := row.Scan(&space.LayerId); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreateLayer(layerId string) (bool, error) {
	_, err := db.Exec("INSERT INTO layer (id) VALUES ($1)", layerId)
	if err != nil {
		return false, err
	}
	return true, nil
}
