package model

import (
	"github.com/google/uuid"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
)

func InsertContentLocation(location common.Location, contentId string) error {
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO content_location (id, content_id, lat, lon, height, scale) VALUES ($1, $2, $3, $4, $5, $6)", uuid.String(), contentId, location.Lat, location.Lon, location.Height, location.Scale)
	if err != nil {
		return err
	}
	return nil
}

func GetCurrentContentLocation(contentId string) (common.Location, error) {
	row := db.QueryRow("SELECT lat, lon, height, scale FROM content_location WHERE content_id = $1", contentId)

	var location common.Location
	if err := row.Scan(&location.Lat, &location.Lon, &location.Height, &location.Scale); err != nil {
		return common.Location{}, err
	}
	return location, nil
}

func UpdateContentLocation(location common.Location, contentId string) error {
	_, err := db.Exec("UPDATE content_location SET lat = $1, lon = $2, height = $3, scale = $4 WHERE content_id = $5", location.Lat, location.Lon, location.Height, location.Scale, contentId)
	if err != nil {
		return err
	}
	return nil
}

func GetCurrentLatLon(contentId string) (float64, float64, error) {
	row := db.QueryRow("SELECT lat, lon FROM content_location WHERE content_id = $1", contentId)

	var lat, lon float64
	if err := row.Scan(&lat, &lon); err != nil {
		return 0, 0, err
	}
	return lat, lon, nil
}
