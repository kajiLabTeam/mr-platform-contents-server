package service

import (
	"errors"

	"github.com/go-viper/mapstructure/v2"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
)

func UpdateHtml2dContent(req common.RequestUpdateContent) (bool, error) {
	// html2dContent（common.Html2d）にreq.Content(Interface)を変換
	var html2dContent common.Html2d
	if err := mapstructure.Decode(req.Content, &html2dContent); err != nil {
		return false, err
	}

	// コンテンツがあるか確認
	isExist, err := model.IsExistHtml2dContent(req.ContentId)
	if err != nil {
		return false, err
	}
	if !isExist {
		return false, errors.New("content does not exist")
	}

	// コンテンツを更新
	isUpdated, err := model.UpdateHtml2dContent(req.ContentId, html2dContent)
	if err != nil {
		return false, err
	}
	if !isUpdated {
		return false, errors.New("failed to update content")
	}

	return true, nil
}
