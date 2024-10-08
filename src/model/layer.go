package model

import (
	"database/sql"
)

func ExistLayer(layerId string) (bool, error) {
	row := db.QueryRow("SELECT id FROM layer WHERE id = $1", layerId)

	var layer string
	if err := row.Scan(&layer); err != nil {
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
