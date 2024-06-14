package model

import (
	"errors"
	"net/http"

	"github.com/uptrace/bun"
)

type Person struct {
	bun.BaseModel `bun:"table:person,alias:pe"`

	ID         int    `bun:",pk,type:int,autoincrement"`
	RequestID  string `json:"request_id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Sitizen    string
	Region     string
	Gender     string
	Passport   string
	Birthday   string
	Phone      string
	Email      string
	IsLeader   bool `json:"is_leader"`
}

func (a *Person) Bind(r *http.Request) error {
	if len(a.RequestID) == 0 {
		return errors.New("empty RequestID")
	}
	if len(a.FirstName) == 0 {
		return errors.New("empty FirstName")
	}
	if len(a.MiddleName) == 0 {
		return errors.New("empty MiddleName")
	}
	if len(a.LastName) == 0 {
		return errors.New("empty LastName")
	}
	if len(a.Sitizen) == 0 {
		return errors.New("empty Sitizen")
	}
	if len(a.Region) == 0 {
		return errors.New("empty Region")
	}
	if len(a.Gender) == 0 {
		return errors.New("empty Gender")
	}
	if len(a.Passport) == 0 {
		return errors.New("empty Passport")
	}
	if len(a.Birthday) == 0 {
		return errors.New("empty Birthday")
	}
	if len(a.Phone) == 0 {
		return errors.New("empty Phone")
	}
	if len(a.Email) == 0 {
		return errors.New("empty Email")
	}

	// just a post-process after a decode..
	return nil
}
