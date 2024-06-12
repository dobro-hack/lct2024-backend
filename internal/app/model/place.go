package model

import (
	"encoding/json"
	"net/http"

	"github.com/uptrace/bun"
)

type Place struct {
	bun.BaseModel `bun:"table:place,alias:pl" swaggerignore:"true"`

	ID          int `bun:",pk,type:int,autoincrement"`
	RouteID     int
	Name        string
	Description string
	Icon        string
	Location    json.RawMessage `bun:"type:json"`
}

type PlaceList struct {
	Items           []Place
	TotalItemsCount int
}

func (res *PlaceList) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
