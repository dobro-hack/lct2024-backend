package v1

import "github.com/dobro-hack/lct2024-backend/internal/app/api"

type v1handler struct {
	api *api.Api
}

func NewV1(a *api.Api) *v1handler {
	return &v1handler{
		api: a,
	}
}
