package service

import (
	"github.com/kajiLabTeam/mr-platform-contents-server/common"
	"github.com/kajiLabTeam/mr-platform-contents-server/model"
	"github.com/uber/h3-go/v4"
)

func InsertDBH3Relation(lat, lon float64, contentId string) error {
	latLng := h3.NewLatLng(lat, lon)
	var CellAndReses []common.CellAndRes

	for i := 0; i <= 12; i++ {
		cell := h3.LatLngToCell(latLng, i)
		CellAndReses = append(CellAndReses, common.CellAndRes{
			Cell: cell,
			Res:  i,
		})
	}

	err := model.InsertContentForNeo4j(contentId)
	if err != nil {
		return err
	}

	err = model.InsertCellToContentRelations(CellAndReses, contentId)
	if err != nil {
		return err
	}

	return nil
}
