package utils

import (
	"github.com/uber/h3-go/v4"
)

func GetH3Cells(latLng h3.LatLng) []h3.Cell {
	var cells []h3.Cell
	maxResolution := 12
	for i := 0; i <= maxResolution; i++ {
		cell := h3.LatLngToCell(latLng, i)
		cells = append(cells, cell)
	}
	return cells
}
