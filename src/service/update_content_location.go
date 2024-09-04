package service

import (
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
)

func UpdateContentLocation(location common.Location, contentId string) error {
	// 更新があるか確認
	currentLocation, err := model.GetCurrentContentLocation(contentId)
	if err != nil {
		return err
	}
	if currentLocation == location {
		return nil
	}

	// 更新
	err = model.UpdateContentLocation(location, contentId)
	if err != nil {
		return err
	}
	return nil
}
