package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
)

func IsExistPublicSpaces(organizationId string) (bool, error) {
	row := db.QueryRow("SELECT id, organization_id FROM public_spaces WHERE organization_id = $1", organizationId)

	var space common.PublicSpace
	if err := row.Scan(&space.PublicSpaceId, &space.OrganizationId); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreatePublicSpaces(organizationId string) (string, error) {
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO public_spaces (id, organization_id) VALUES ($1, $2)", uuid.String(), organizationId)
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

func GetPublicSpaceId(organizationId string) (string, error) {
	row := db.QueryRow("SELECT id FROM public_spaces WHERE organization_id = $1", organizationId)

	var space common.PublicSpace
	if err := row.Scan(&space.PublicSpaceId); err != nil {
		return "", err
	}
	return space.PublicSpaceId, nil
}
