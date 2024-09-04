package service

import (
	"fmt"

	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
)

func UpdateContentLocation(location common.Location, contentId string) error {
	// 更新があるか確認
	currentLocation, err := model.GetCurrentContentLocation(contentId)
	if err != nil {
		return err
	}
	fmt.Println("currentLocation", currentLocation)
	fmt.Println("location", location)
	if currentLocation == location {
		return nil
	}

	fmt.Println("UpdateContentLocation")

	// 更新
	err = model.UpdateContentLocation(location, contentId)
	if err != nil {
		return err
	}
	return nil
}
