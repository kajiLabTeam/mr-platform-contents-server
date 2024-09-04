package service

import (
	"errors"

	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
)

func GetContent(contentId string) (content common.Content, err error) {
	contentType, err := getDataType(contentId)
	if err != nil {
		return common.Content{}, err
	}
	content = common.Content{
		ContentId:   contentId,
		ContentType: contentType,
	}
	switch contentType {
	case "html2d":
		html2dContent, err := getDataFromHtml2d(contentId)
		if err != nil {
			return common.Content{}, err
		}
		content.Content = html2dContent

	default:
		return common.Content{}, errors.New("invalid content type")
	}

	location, err := model.GetCurrentContentLocation(contentId)
	if err != nil {
		return common.Content{}, err
	}

	content.Location = location

	return content, nil
}

func getDataType(contentId string) (contentType string, err error) {
	// データが存在するか確認
	isExist, err := model.IsExistContentId(contentId)
	if err != nil {
		return "", err
	}
	if !isExist {
		return "", errors.New("contentId does not exist")
	}

	contentType, err = model.GetContentType(contentId)
	if err != nil {
		return "", err
	}

	return contentType, nil
}

func getDataFromHtml2d(contentId string) (content common.Html2d, err error) {
	// データが存在するか確認
	isExist, err := model.IsExistContentId(contentId)
	if err != nil {
		return common.Html2d{}, err
	}
	if !isExist {
		return common.Html2d{}, nil
	}

	// SQLから取得 GetHtml2dContent
	content, err = model.GetHtml2dContent(contentId)
	if err != nil {
		return common.Html2d{}, err
	}

	return content, nil
}
