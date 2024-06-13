package model

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/uptrace/bun"
)

type Request struct {
	bun.BaseModel `bun:"table:request,alias:r"`

	ID          int `bun:",pk,type:int,autoincrement"`
	Name        string
	Description string
	Area        json.RawMessage `bun:"type:json"`
}

type SaveRequest struct {
	Quantity  int
	Type      string
	PersonID  int
	OrgID     int
	RouteID   int
	DateStart time.Time
}

func (a *SaveRequest) Bind(r *http.Request) error {
	// just a post-process after a decode..
	return nil
}

type SaveRequestResponse struct {
	Status string
}

func (res *SaveRequestResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
