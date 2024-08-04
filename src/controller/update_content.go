package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-viper/mapstructure/v2"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
	"github.com/kajiLabTeam/mr-platform-contents-server/util"
)

func UpdateContent(c *gin.Context) {
	// RequestUpdateContents型のデータを取得
	var req common.RequestUpdateContent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// コンテンツがあるか確認
	isExist, err := model.IsExistContentId(req.ContentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content id"})
		return
	}

	// type の確認
	switch req.ContentType {
	case "html2d":
		if _, err := updateHtml2dContent(req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
		return
	}

	// コンテンツの取得
	content, err := util.GetContent(req.ContentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 201を返す
	c.JSON(http.StatusCreated, content)
}

func updateHtml2dContent(req common.RequestUpdateContent) (bool, error) {
	// 同一コンテンツが存在するか確認
	// html2dContent（common.Html2d）にreq.Content(Interface)を変換
	var html2dContent common.Html2d
	if err := mapstructure.Decode(req.Content, &html2dContent); err != nil {
		return false, err
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
