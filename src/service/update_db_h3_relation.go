package service

import (
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
	"github.com/kajiLabTeam/mr-platform-contents-server/utils"
	"github.com/uber/h3-go/v4"
)

func UpdateDBH3Relation(lat float64, lon float64, contentId string) error {
	// 更新があるか確認
	currentLat, currentlon, err := model.GetCurrentLatLon(contentId)
	if err != nil {
		return err
	}

	if currentLat == lat && currentlon == lon {
		return nil
	}

	// 検証するコード
	latLng := h3.LatLng{Lat: lat, Lng: lon}
	cells := utils.GetH3Cells(latLng)

	// 既存の関係性を削除
	err = model.RemoveRelationH3CellToContentIdForNeo4j(contentId)
	if err != nil {
		return err
	}

	// 新しい関係性を追加
	err = model.InsertCellToContentIdRelations(cells, contentId)
	if err != nil {
		return err
	}

	return nil
}
