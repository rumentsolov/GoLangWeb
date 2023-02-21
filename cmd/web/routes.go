package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rumentsolov/GoLangWeb/config"
	"github.com/rumentsolov/GoLangWeb/pkg/handlers"
)

// routes is a new http.Handler that is called MAX /multiplexer/

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
