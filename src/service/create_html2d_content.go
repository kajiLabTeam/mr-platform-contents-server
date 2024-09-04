package service

import (
	"errors"

	"github.com/go-viper/mapstructure/v2"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
)

func CreateHtml2dContent(req common.RequestCreateContent) (contentId string, err error) {
	// 同一コンテンツが存在するか確認
	// html2dContent（common.Html2d）にreq.Content(Interface)を変換
	var html2dContent common.Html2d
	if err := mapstructure.Decode(req.Content, &html2dContent); err != nil {
		return "", err
	}

	isExist, err := model.IsExistHtml2dContentExceptId(html2dContent)
	if err != nil {
		return "", err
	}
	if isExist {
		return "", errors.New("the same content already exists")
	}

	// コンテンツを作成
	contentId, err = model.CreateContent(req.ContentType)
	if err != nil {
		return "", err
	}

	// コンテンツを作成
	if err := model.CreateHtml2dContent(contentId, html2dContent); err != nil {
		return "", err
	}

	return contentId, nil
}
