package util

import (
	"errors"

	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
)

func GetContent(contentId string) (common.Content, error) {
	contentType, err := getDataType(contentId)
	if err != nil {
		return common.Content{}, err
	}
	switch contentType {
	case "html2d":
		content, err := getDataFromHtml2d(contentId)
		if err != nil {
			return common.Content{}, err
		}
		return common.Content{
			ContentId:   contentId,
			ContentType: contentType,
			Content:     content,
		}, nil

	default:
		return common.Content{}, errors.New("invalid content type")
	}
}

func getDataType(contentId string) (string, error) {
	// データが存在するか確認
	isExist, err := model.IsExistContentId(contentId)
	if err != nil {
		return "", err
	}
	if !isExist {
		return "", errors.New("contentId does not exist")
	}

	contentType, err := model.GetContentType(contentId)
	if err != nil {
		return "", err
	}

	return contentType, nil
}

func getDataFromHtml2d(contentId string) (common.Html2d, error) {
	// データが存在するか確認
	isExist, err := model.IsExistContentId(contentId)
	if err != nil {
		return common.Html2d{}, err
	}
	if !isExist {
		return common.Html2d{}, nil
	}

	// SQLから取得 GetHtml2dContent
	content, err := model.GetHtml2dContent(contentId)
	if err != nil {
		return common.Html2d{}, err
	}
	// 構造体が空の場合はエラーを返す
	if content == (common.Html2d{}) {
		return common.Html2d{}, nil
	}

	return content, nil
}
