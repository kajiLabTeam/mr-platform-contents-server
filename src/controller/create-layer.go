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
	layerId, err := model.CreateLayer(req.OrganizationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// レスポンスの作成
	res := common.ResponseCreateLayer{
		LayerId: layerId,
	}

	// 201を返す
	c.JSON(http.StatusCreated, res)
}
