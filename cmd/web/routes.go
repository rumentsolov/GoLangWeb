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

	// just a variable name
	fileServer := http.FileServer(http.Dir("./static/")) // if doesnt show the route to the  website doesnt show the images
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer)) // I am using it to show where are the images

	

	return mux
}
