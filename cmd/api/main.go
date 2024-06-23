package main

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/dobro-hack/lct2024-backend/internal/app/api"
	v1 "github.com/dobro-hack/lct2024-backend/internal/app/api/v1"
	"github.com/dobro-hack/lct2024-backend/internal/app/config"
	"github.com/dobro-hack/lct2024-backend/internal/app/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"google.golang.org/api/option"
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

	opt := option.WithCredentialsFile(config.FireBaseFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	push, err := app.Messaging(ctx)
	if err != nil {
		panic(err)
	}

	api := api.NewAPI(r, db, config, push)
	v1 := v1.NewV1(api)
	v1.ConfigureRouter()

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
