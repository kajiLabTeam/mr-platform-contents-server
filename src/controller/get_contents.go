package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/service"
)

func GetContents(c *gin.Context) {
	// RequestGetContents型のデータを取得
	var req common.RequestGetContents
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 返すデータの作成
	var res []common.Content

	for index := range req.ContentIds {
		content, err := service.GetContent(req.ContentIds[index])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res = append(res, content)
	}

	c.JSON(http.StatusOK, res)
}
