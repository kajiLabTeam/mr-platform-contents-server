package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/common"
)

func IsExistContent(contentType string, domain string) (bool, error) {
	row := db.QueryRow("SELECT id, type, domain FROM contents WHERE type = $1 AND domain = $2", contentType, domain)

	var content common.Content
	if err := row.Scan(&content.ContentId, &content.ContentType, &content.Domain); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GetContent(contentId string) (common.Content, error) {
	row := db.QueryRow("SELECT id, type, domain FROM contents WHERE id = $1", contentId)

	var content common.Content
	if err := row.Scan(&content.ContentId, &content.ContentType, &content.Domain); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return common.Content{}, nil
		}
		return common.Content{}, err
	}
	return content, nil
}

func GetContentId(contentType string, domain string) (string, error) {
	row := db.QueryRow("SELECT id FROM contents WHERE type = $1 AND domain = $2", contentType, domain)

	var content common.Content
	if err := row.Scan(&content.ContentId); err != nil {
		return "", err
	}
	return content.ContentId, nil
}

func CreateContent(content common.RequestContent) (string, error) {
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO contents (id, type, domain) VALUES ($1, $2, $3)", uuid.String(), content.ContentType, content.Domain)
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
