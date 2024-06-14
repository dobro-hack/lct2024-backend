package model

import (
	"errors"
	"net/http"

	"github.com/uptrace/bun"
)

type Device struct {
	bun.BaseModel `bun:"table:device,alias:d"`

	DeviceID string `bun:",pk,type:string"`
	UserID   int
}

type SaveUserRequest struct {
	DeviceID string `json:"device_id"`
}

func (a *SaveUserRequest) Bind(r *http.Request) error {
	if len(a.DeviceID) == 0 {
		return errors.New("empty device ID")
	}

	// just a post-process after a decode..
	return nil
}
