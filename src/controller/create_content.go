package controller

import (
	"encoding/json"
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
	exist, err := model.ExistLayer(req.LayerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid layer id"})
		return
	}

	var contentId string

	// type の確認
	switch req.ContentType {
	case "html2d":
		contentId, err = service.CreateHtml2dContent(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
		return
	}

	// content_location に追加
	err = model.InsertContentLocation(req.Location, contentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Neo4jに関係性の追加
	err = service.InsertDBH3Relation(req.Location.Lat, req.Location.Lon, contentId)
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

	// contentをJSON文字列に変換
	contentStr, err := json.Marshal(content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 201を返す
	c.Data(http.StatusCreated, "text", contentStr)
}
