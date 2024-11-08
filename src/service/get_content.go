package service

import (
	"errors"
	"log"

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

		log.Println("GetContent html2dContent: ", html2dContent)

		content.Content = html2dContent

	default:
		return common.Content{}, errors.New("invalid content type")
	}

	location, err := model.GetCurrentContentLocation(contentId)
	if err != nil {
		return common.Content{}, err
	}

	content.Location = location

	log.Println("GetContent content: ", content)

	return content, nil
}

func getDataType(contentId string) (contentType string, err error) {
	// データが存在するか確認
	exist, err := model.ExistContentId(contentId)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", errors.New("contentId does not exist")
	}

	contentType, err = model.GetContentType(contentId)
	if err != nil {
		return "", err
	}

	return contentType, nil
}

func getDataFromHtml2d(contentId string) (content common.ReturnHtml2d, err error) {
	// データが存在するか確認
	exist, err := model.ExistContentId(contentId)
	if err != nil {
		return common.ReturnHtml2d{}, err
	}
	if !exist {
		return common.ReturnHtml2d{}, nil
	}

	// SQLから取得 GetHtml2dContent
	contentBySQL, err := model.GetHtml2dContent(contentId)
	if err != nil {
		return common.ReturnHtml2d{}, err
	}

	// MinIOから取得
	imgURL, err := model.MinioGetPng("html2d", contentId)
	if err != nil {
		return common.ReturnHtml2d{}, err
	}

	log.Println("getDataFromHtml2d imgURL: ", imgURL)

	content = common.ReturnHtml2d{
		Size:     contentBySQL.Size,
		TextType: contentBySQL.TextType,
		TextURL:  contentBySQL.TextURL,
		ImgURL:   imgURL,
	}

	return content, nil
}
