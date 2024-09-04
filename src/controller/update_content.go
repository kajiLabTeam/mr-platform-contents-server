package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
	"github.com/kajiLabTeam/mr-platform-contents-server/service"
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

	var isUpdated bool

	// type の確認
	switch req.ContentType {
	case "html2d":
		isUpdated, err = service.UpdateHtml2dContent(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !isUpdated {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update content"})
			return
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
		return
	}

	// Neo4jに関係性の更新
	err = service.UpdateDBH3Relation(req.Location.Lat, req.Location.Lon, req.ContentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// content_location の更新
	err = service.UpdateContentLocation(req.Location, req.ContentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// コンテンツの取得
	content, err := service.GetContent(req.ContentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 201を返す
	c.JSON(http.StatusCreated, content)
}
