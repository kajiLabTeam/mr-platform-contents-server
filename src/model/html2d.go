package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
)

func GetHtml2dContent(contentId string) (common.Html2d, error) {
	row := db.QueryRow("SELECT location_lat, location_lon, location_height,  rotation_roll, rotation_pitch, rotation_yaw,  size_width, size_height,  text_type, text_url, style_url FROM html2d WHERE content_id = $1", contentId)

	var content common.Html2d
	if err := row.Scan(&content.Location.Lat, &content.Location.Lon, &content.Location.Height,&content.Rotation.Roll, &content.Rotation.Pitch, &content.Rotation.Yaw,&content.Size.Width, &content.Size.Height,&content.TextType, &content.TextURL, &content.StyleURL); err != nil {
		return common.Html2d{}, err
	}
	return content, nil
}

func CreateHtml2dContent(contentId string, content common.Html2d) error {
	// uuidを生成
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO html2d (id, content_id, location_lat, location_lon, location_height, rotation_roll, rotation_pitch, rotation_yaw, size_width, size_height, text_type, text_url, style_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",uuid.String(), contentId,content.Location.Lat, content.Location.Lon, content.Location.Height,content.Rotation.Roll, content.Rotation.Pitch, content.Rotation.Yaw,content.Size.Width, content.Size.Height,content.TextType, content.TextURL, content.StyleURL)
	if err != nil {
		return err
	}
	
	// ログテーブルにも追加
	if err := insertLogHtml2dContent(uuid.String(), contentId, content); err != nil {
		return err
	}
	
	return nil
}

func IsExistHtml2dContent(contentId string) (bool, error) {
	row := db.QueryRow("SELECT location_lat, location_lon, location_height, rotation_roll, rotation_pitch, rotation_yaw, size_width, size_height, text_type, text_url, style_url FROM html2d WHERE content_id = $1", contentId)

	var content common.SQLHtml2d
	if err := row.Scan(&content.Location.Lat, &content.Location.Lon, &content.Location.Height,&content.Rotation.Roll, &content.Rotation.Pitch, &content.Rotation.Yaw,&content.Size.Width, &content.Size.Height,&content.TextType, &content.TextURL, &content.StyleURL,); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// id以外を比較して同一のコンテンツが存在するか確認
func IsExistHtml2dContentExceptId(checkContent common.Html2d) (bool, error) {
	row := db.QueryRow("SELECT location_lat, location_lon, location_height, rotation_roll, rotation_pitch, rotation_yaw, size_width, size_height, text_type, text_url, style_url FROM html2d WHERE location_lat = $1 AND location_lon = $2 AND location_height = $3 AND rotation_roll = $4 AND rotation_pitch = $5 AND rotation_yaw = $6 AND size_width = $7 AND size_height = $8 AND text_type = $9 AND text_url = $10 AND style_url = $11",checkContent.Location.Lat, checkContent.Location.Lon, checkContent.Location.Height,checkContent.Rotation.Roll, checkContent.Rotation.Pitch, checkContent.Rotation.Yaw,checkContent.Size.Width, checkContent.Size.Height,checkContent.TextType, checkContent.TextURL, checkContent.StyleURL)

	var content common.SQLHtml2d
	if err := row.Scan(&content.Location.Lat, &content.Location.Lon, &content.Location.Height,&content.Rotation.Roll, &content.Rotation.Pitch, &content.Rotation.Yaw,&content.Size.Width, &content.Size.Height,&content.TextType, &content.TextURL, &content.StyleURL,); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func UpdateHtml2dContent(contentId string, content common.Html2d) (bool, error) {
	_, err := db.Exec("UPDATE html2d SET location_lat = $1, location_lon = $2, location_height = $3, rotation_roll = $4, rotation_pitch = $5, rotation_yaw = $6, size_width = $7, size_height = $8, text_type = $9, text_url = $10, style_url = $11 WHERE content_id = $12",content.Location.Lat, content.Location.Lon, content.Location.Height,content.Rotation.Roll, content.Rotation.Pitch, content.Rotation.Yaw,content.Size.Width, content.Size.Height,content.TextType, content.TextURL, content.StyleURL, contentId)
	if err != nil {
		return false, err
	}

	// ログテーブルにも追加
	html2dId, err := getHtml2dId(contentId)
	if err != nil {
		return false, err
	}

	if err := insertLogHtml2dContent(html2dId, contentId, content); err != nil {
		return false, err
	}
	return true, nil
}

func getHtml2dId(contentId string) (string, error) {
	row := db.QueryRow("SELECT id FROM html2d WHERE content_id = $1", contentId)

	var html2dId string
	if err := row.Scan(&html2dId); err != nil {
		return "", err
	}
	return html2dId, nil
}

func insertLogHtml2dContent(html2dId string, contentId string, content common.Html2d) error {
	// uuidを生成
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO log_html2d (id, html2d_id, content_id, location_lat, location_lon, location_height, rotation_roll, rotation_pitch, rotation_yaw, size_width, size_height, text_type, text_url, style_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",uuid.String(), html2dId, contentId,content.Location.Lat, content.Location.Lon, content.Location.Height,content.Rotation.Roll, content.Rotation.Pitch, content.Rotation.Yaw,content.Size.Width, content.Size.Height,content.TextType, content.TextURL, content.StyleURL)
	if err != nil {
		return err
	}
	return nil
}
