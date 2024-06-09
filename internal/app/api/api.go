package api

import (
	"github.com/dobro-hack/lct2024-backend/internal/app/config"
	"github.com/dobro-hack/lct2024-backend/internal/app/storage"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router  *chi.Mux
	Storage *storage.Database
	Config  config.Config
}

// NewAPI инициализируем сервер
func NewAPI(r *chi.Mux, s *storage.Database, c config.Config) *Api {
	return &Api{
		Router:  r,
		Storage: s,
		Config:  c,
	}
}
