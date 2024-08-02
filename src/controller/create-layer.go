package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
)

func CreateLayer(c *gin.Context) {
	// RequestCreateLayer型のデータを取得
	var req common.RequestCreateLayer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// レイヤーを作成
	isCreated, err := model.CreateLayer(req.LayerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isCreated {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create layer"})
		return
	}

	// 201を返す
	c.JSON(http.StatusCreated, req)
}
