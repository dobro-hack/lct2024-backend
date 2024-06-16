package model

import (
	"errors"
	"net/http"

	"github.com/uptrace/bun"
)

type Org struct {
	bun.BaseModel `bun:"table:org,alias:o" swaggerignore:"true"`

	ID        int    `bun:",pk,type:int,autoincrement"`
	RequestID string `json:"request_id"`
	Name      string
	INN       string
	KPP       string
	Address   string
	Phone     string
	Email     string
}

func (a *Org) Bind(r *http.Request) error {
	if len(a.RequestID) == 0 {
		return errors.New("empty RequestID")
	}
	if len(a.Name) == 0 {
		return errors.New("empty Name")
	}
	if len(a.INN) == 0 {
		return errors.New("empty INN")
	}
	if len(a.KPP) == 0 {
		return errors.New("empty KPP")
	}
	if len(a.Address) == 0 {
		return errors.New("empty Address")
	}
	if len(a.Phone) == 0 {
		return errors.New("empty Phone")
	}

	// just a post-process after a decode..
	return nil
}
