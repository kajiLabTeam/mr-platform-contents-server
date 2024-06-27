package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/common"
)

func GetHtml2dContent(contentId string) (common.Html2d, error) {
	row := db.QueryRow("SELECT id, content_id,   location_x, location_y, location_z, rotation_row, rotation_pitch, rotation_yaw, text_type FROM html2d WHERE content_id = $1", contentId)

	var content common.Html2d
	if err := row.Scan(&content.Id, &content.ContentId, &content.Location.X, &content.Location.Y, &content.Location.Z, &content.Rotation.Row, &content.Rotation.Pitch, &content.Rotation.Yaw, &content.TextType); err != nil {
		return common.Html2d{}, err
	}
	return content, nil
}

func CreateHtml2dContent(content common.Html2d) error {
	// uuidを生成
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO html2d (id, content_id,  location_x, location_y, location_z, rotation_row, rotation_pitch, rotation_yaw, text_type, text_url, style_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", uuid.String(), content.ContentId, content.Location.X, content.Location.Y, content.Location.Z, content.Rotation.Row, content.Rotation.Pitch, content.Rotation.Yaw, content.TextType, content.TextURL, content.StyleURL)
	if err != nil {
		return err
	}
	return nil
}

func IsExistHtml2dContent(contentId string) (bool, error) {
	row := db.QueryRow("SELECT id, content_id,  location_x, location_y, location_z, rotation_row, rotation_pitch, rotation_yaw,text_type, text_url, style_url FROM html2d WHERE content_id = $1", contentId)

	var content common.Html2d
	if err := row.Scan(&content.Id, &content.ContentId, &content.Location.X, &content.Location.Y, &content.Location.Z, &content.Rotation.Row, &content.Rotation.Pitch, &content.Rotation.Yaw, &content.TextType, &content.TextURL, &content.StyleURL); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}
