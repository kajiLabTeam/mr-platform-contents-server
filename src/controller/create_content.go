package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
	"github.com/kajiLabTeam/mr-platform-contents-server/service"
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

	var contentId string
	var lat, lon float64

	// type の確認
	switch req.ContentType {
	case "html2d":
		contentId, lat, lon, err = service.CreateHtml2dContent(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
		return
	}

	err = service.InsertDBH3Relation(lat, lon, contentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// コンテンツの取得
	content, err := service.GetContent(contentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 201を返す
	c.JSON(http.StatusCreated, content)
}
