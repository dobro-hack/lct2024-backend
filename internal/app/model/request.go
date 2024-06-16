package model

import (
	"errors"
	"net/http"

	"github.com/uptrace/bun"
)

type Request struct {
	bun.BaseModel `bun:"table:request,alias:r" swaggerignore:"true"`

	RequestID string `bun:",pk,type:string" json:"request_id"`
	UserID    int    `json:"user_id"`
	Quantity  int
	RouteID   int    `json:"route_id"`
	DateStart string `json:"date_start"`

	Person []Person `bun:"rel:has-many,join:request_id=request_id"`
	Org    *Org     `bun:"rel:has-one,join:request_id=request_id"`
}

func (a *Request) Bind(r *http.Request) error {
	if len(a.RequestID) == 0 {
		return errors.New("empty RequestID")
	}
	if a.RouteID == 0 {
		return errors.New("empty RouteID")
	}
	if a.Quantity == 0 {
		return errors.New("empty Quantity")
	}

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

type RequestList struct {
	Items           []Request
	TotalItemsCount int
}

func (res *RequestList) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
