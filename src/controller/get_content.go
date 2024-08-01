package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
)

func GetContent(c *gin.Context) {
	// RequestGetContents型のデータを取得
	var req common.RequestGetContents
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 返すデータの作成
	var res []common.Content

	for index := range req.ContentIds {
		contentType, err := getDataType(req.ContentIds[index])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var tmp interface{}
		switch contentType {
		case "html2d":
			content, err := getDataFromHtml2d(req.ContentIds[index])
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if content == (common.Html2d{}) {
				continue
			}
			tmp = content

		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
			return
		}

		res = append(res, common.Content{
			ContentId:   req.ContentIds[index],
			ContentType: contentType,
			Content:     tmp,
		})
	}

	c.JSON(http.StatusOK, res)
}

func getDataType(contentId string) (string, error) {
	// データが存在するか確認
	isExist, err := model.IsExistContentId(contentId)
	if err != nil {
		return "", err
	}
	if !isExist {
		return "", errors.New("contentId does not exist")
	}

	contentType, err := model.GetContentType(contentId)
	if err != nil {
		return "", err
	}

	return contentType, nil
}

func getDataFromHtml2d(contentId string) (common.Html2d, error) {
	// データが存在するか確認
	isExist, err := model.IsExistContentId(contentId)
	if err != nil {
		return common.Html2d{}, err
	}
	if !isExist {
		return common.Html2d{}, nil
	}

	// SQLから取得 GetHtml2dContent
	content, err := model.GetHtml2dContent(contentId)
	if err != nil {
		return common.Html2d{}, err
	}
	// 構造体が空の場合はエラーを返す
	if content == (common.Html2d{}) {
		return common.Html2d{}, nil
	}

	return content, nil
}
