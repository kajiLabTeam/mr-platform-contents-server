package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/common"
)

func GetModel3dContent(contentId string) (common.Model3d, error) {
	row := db.QueryRow("SELECT * FROM model3d WHERE content_id = $1", contentId)

	var content common.Model3d
	if err := row.Scan(&content.Id, &content.ContentId, &content.Location.X, &content.Location.Y, &content.Location.Z, &content.Rotation.Row, &content.Rotation.Pitch, &content.Rotation.Yaw, &content.PresignedURL); err != nil {
		return common.Model3d{}, err
	}
	return content, nil
}

func CreateModel3dContent(content common.Model3d) error {
	// uuidを生成
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO model3d (id, content_id,location_x, location_y, location_z, rotation_row, rotation_pitch, rotation_yaw, file_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", uuid.String(), content.ContentId, content.Location.X, content.Location.Y, content.Location.Z, content.Rotation.Row, content.Rotation.Pitch, content.Rotation.Yaw, content.PresignedURL)
	if err != nil {
		return err
	}
	return nil
}

func IsExistModel3dContent(contentId string) (bool, error) {
	row := db.QueryRow("SELECT id, content_id,location_x, location_y, location_z, rotation_row, rotation_pitch, rotation_yaw, file_url FROM model3d WHERE content_id = $1", contentId)

	var content common.Model3d
	if err := row.Scan(&content.Id, &content.ContentId, &content.Location.X, &content.Location.Y, &content.Location.Z, &content.Rotation.Row, &content.Rotation.Pitch, &content.Rotation.Yaw, &content.PresignedURL); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}
