package model

import (
	"encoding/json"
	"net/http"

	"github.com/uptrace/bun"
)

type Route struct {
	bun.BaseModel `bun:"table:route,alias:ro" swaggerignore:"true"`

	ID              int `bun:",pk,type:int,autoincrement"`
	ParkID          int
	Park            Park     `bun:"rel:belongs-to"`
	Places          []*Place `bun:"rel:has-many"`
	Name            string
	Description     string
	HowToGet        string
	WhatToTake      string
	InEmergency     string
	Recommendations string
	Length          int
	Duration        int
	Height          int
	Difficulty      string
	Load            json.RawMessage `bun:"type:json" swaggerignore:"true"`
	MaxLoad         json.RawMessage `bun:"type:json" swaggerignore:"true"`
	Photo           json.RawMessage `bun:"type:json" swaggerignore:"true"`
	GpxData         string
}

type RouteList struct {
	Items           []Route
	TotalItemsCount int
}

func (res *RouteList) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
