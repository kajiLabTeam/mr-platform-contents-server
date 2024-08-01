package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
)

func IsExistContentId(contentId string) (bool, error) {
	row := db.QueryRow("SELECT id FROM contents WHERE id = $1", contentId)

	var content common.Content
	if err := row.Scan(&content.ContentId); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GetContentType(contentId string) (string, error) {
	row := db.QueryRow("SELECT type FROM contents WHERE id = $1", contentId)

	var contentType string
	if err := row.Scan(&contentType); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return "", nil
		}
		return "", err
	}
	return contentType, nil
}

func GetContent(contentId string) (common.Content, error) {
	row := db.QueryRow("SELECT id, type FROM contents WHERE id = $1", contentId)

	var content common.Content
	if err := row.Scan(&content.ContentId, &content.ContentType); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return common.Content{}, nil
		}
		return common.Content{}, err
	}
	return content, nil
}

func CreateContent(contentType string) (string, error) {
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO contents (id, type) VALUES ($1, $2)", uuid.String(), contentType)
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
