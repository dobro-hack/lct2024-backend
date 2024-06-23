package model

import (
	"errors"
	"net/http"
)

type Push struct {
	Message string
	Title   string
}

func (a *Push) Bind(r *http.Request) error {
	if len(a.Message) == 0 {
		return errors.New("empty Message")
	}
	if len(a.Title) == 0 {
		return errors.New("empty Title")
	}
	return nil
}
