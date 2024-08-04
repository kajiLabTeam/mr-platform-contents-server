package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
)

func GetHtml2dContent(contentId string) (common.Html2d, error) {
	roll := db.QueryRow("SELECT location_x, location_y, location_z, rotation_roll, rotation_pitch, rotation_yaw, size_width, size_height, text_type, text_url, style_url FROM html2d WHERE content_id = $1", contentId)

	var content common.Html2d
	if err := roll.Scan(&content.Location.X, &content.Location.Y, &content.Location.Z, &content.Rotation.Roll, &content.Rotation.Pitch, &content.Rotation.Yaw, &content.Size.Width, &content.Size.Height, &content.TextType, &content.TextURL, &content.StyleURL); err != nil {
		return common.Html2d{}, err
	}
	return content, nil
}

func CreateHtml2dContent(contentId string, content common.Html2d) error {
	// uuidを生成
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO html2d (id, content_id, location_x, location_y, location_z, rotation_roll, rotation_pitch, rotation_yaw, size_width, size_height, text_type, text_url, style_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)", uuid.String(), contentId, content.Location.X, content.Location.Y, content.Location.Z, content.Rotation.Roll, content.Rotation.Pitch, content.Rotation.Yaw, content.Size.Width, content.Size.Height, content.TextType, content.TextURL, content.StyleURL)
	if err != nil {
		return err
	}
	return nil
}

func IsExistHtml2dContent(contentId string) (bool, error) {
	roll := db.QueryRow("SELECT location_x, location_y, location_z, rotation_roll, rotation_pitch, rotation_yaw, size_width, size_height,text_type, text_url, style_url FROM html2d WHERE content_id = $1", contentId)

	var content common.SQLHtml2d
	if err := roll.Scan(&content.Location.X, &content.Location.Y, &content.Location.Z, &content.Rotation.Roll, &content.Rotation.Pitch, &content.Rotation.Yaw, &content.Size.Width, &content.Size.Height, &content.TextType, &content.TextURL, &content.StyleURL); err != nil {
		if err == sql.ErrNoRows {
			// No rolls were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// id以外を比較して同一のコンテンツが存在するか確認
func IsExistHtml2dContentExceptId(checkContent common.Html2d) (bool, error) {
	roll := db.QueryRow("SELECT location_x, location_y, location_z, rotation_roll, rotation_pitch, rotation_yaw, size_width, size_height,text_type, text_url, style_url FROM html2d WHERE location_x = $1 AND location_y = $2 AND location_z = $3 AND rotation_roll = $4 AND rotation_pitch = $5 AND rotation_yaw = $6 AND size_width = $7 AND size_height = $8 AND text_type = $9 AND text_url = $10 AND style_url = $11", checkContent.Location.X, checkContent.Location.Y, checkContent.Location.Z, checkContent.Rotation.Roll, checkContent.Rotation.Pitch, checkContent.Rotation.Yaw, checkContent.Size.Width, checkContent.Size.Height, checkContent.TextType, checkContent.TextURL, checkContent.StyleURL)

	var content common.SQLHtml2d
	if err := roll.Scan(&content.Location.X, &content.Location.Y, &content.Location.Z, &content.Rotation.Roll, &content.Rotation.Pitch, &content.Rotation.Yaw, &content.Size.Width, &content.Size.Height, &content.TextType, &content.TextURL, &content.StyleURL); err != nil {
		if err == sql.ErrNoRows {
			// No rolls were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func UpdateHtml2dContent(contentId string, content common.Html2d) (bool, error) {
	_, err := db.Exec("UPDATE html2d SET location_x = $1, location_y = $2, location_z = $3, rotation_roll = $4, rotation_pitch = $5, rotation_yaw = $6, size_width = $7, size_height = $8, text_type = $9, text_url = $10, style_url = $11 WHERE content_id = $12", content.Location.X, content.Location.Y, content.Location.Z, content.Rotation.Roll, content.Rotation.Pitch, content.Rotation.Yaw, content.Size.Width, content.Size.Height, content.TextType, content.TextURL, content.StyleURL, contentId)
	if err != nil {
		return false, err
	}
	return true, nil
}
