package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-viper/mapstructure/v2"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
)

func CreateContent(c *gin.Context) {
	// RequestGetContents型のデータを取得
	var req common.RequestCreateContent
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// レイヤーがあるか確認
	isExist, err := model.IsExistLayer(req.LayerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid layer id"})
		return
	}

	var res common.ResponseCreateContent

	// type の確認
	switch req.ContentType {
	case "html2d":
		contentId, err := createHtml2dContent(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res.ContentId = contentId

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
		return
	}

	// 201を返す
	c.JSON(http.StatusCreated, res)
}

func createHtml2dContent(req common.RequestCreateContent) (string, error) {
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
	contentId, err := model.CreateContent(req.ContentType)
	if err != nil {
		return "", err
	}

	// コンテンツを作成
	if err := model.CreateHtml2dContent(contentId, html2dContent); err != nil {
		return "", err
	}

	return contentId, nil
}
