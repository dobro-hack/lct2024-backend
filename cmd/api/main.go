package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dobro-hack/lct2024-backend/internal/app/api"
	v1 "github.com/dobro-hack/lct2024-backend/internal/app/api/v1"
	"github.com/dobro-hack/lct2024-backend/internal/app/config"
	"github.com/dobro-hack/lct2024-backend/internal/app/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// @title lct2024 API
// @version 1.0
// @description API Server for app
// @host localhost:8080
// @BasePath /
func main() {
	config := config.Get()

	ctx := context.Background()

	// Initialize database
	db, err := storage.Open(ctx, config)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	api := api.NewAPI(r, db, config)
	v1 := v1.NewV1(api)
	v1.ConfigureRouter()

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
	}
}
