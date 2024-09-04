package service

import (
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
	"github.com/kajiLabTeam/mr-platform-contents-server/utils"
	"github.com/uber/h3-go/v4"
)

func InsertDBH3Relation(lat float64, lon float64, contentId string) error {
	latLng := h3.LatLng{Lat: lat, Lng: lon}
	cells := utils.GetH3Cells(latLng)

	err := model.InsertContentForNeo4j(contentId)
	if err != nil {
		return err
	}

	err = model.InsertCellToContentIdRelations(cells, contentId)
	if err != nil {
		return err
	}

	return nil
}
