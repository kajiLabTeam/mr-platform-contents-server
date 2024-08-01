package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
)

func IsExistLayer(layerId string) (bool, error) {
	row := db.QueryRow("SELECT id, organization_id FROM layer WHERE id = $1", layerId)

	var space common.PublicSpace
	if err := row.Scan(&space.LayerId, &space.OrganizationId); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreateLayer(organizationId string) (string, error) {
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO layer (id, organization_id) VALUES ($1, $2)", uuid.String(), organizationId)
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

func GetLayerId(organizationId string) (string, error) {
	row := db.QueryRow("SELECT id FROM layer WHERE organization_id = $1", organizationId)

	var space common.PublicSpace
	if err := row.Scan(&space.LayerId); err != nil {
		return "", err
	}
	return space.LayerId, nil
}
