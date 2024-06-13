package model

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/uptrace/bun"
)

type Message struct {
	bun.BaseModel `bun:"table:message,alias:m"`

	ID       int `bun:",pk,type:int,autoincrement"`
	UserID   int
	SentAt   time.Time
	Type     string
	Message  string
	Location json.RawMessage `bun:"type:json"`
	Status   string
	FileUrl  string
	Phone    string
}

type SaveMessageRequest struct {
	Lat     string
	Lon     string
	Type    string
	Message string
	File    http.File
	FileUrl string
	Phone   string
}

func (a *SaveMessageRequest) Bind(r *http.Request) error {
	// just a post-process after a decode..
	// get name of file

	return nil
}

type SaveMessageResponse struct {
	Status string
}

func (res *SaveMessageResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
