package model

import (
	"net/http"

	"github.com/uptrace/bun"
)

type Park struct {
	bun.BaseModel `bun:"table:park,alias:p"`

	ID          int `bun:",pk,type:int,autoincrement"`
	Name        string
	Description string
	Area        []string
}

type ParkList struct {
	Items           []Park
	TotalItemsCount int
}

func (res *ParkList) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
