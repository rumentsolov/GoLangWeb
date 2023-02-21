package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/rumentsolov/GoLangWeb/config"
	"github.com/rumentsolov/GoLangWeb/pkg/handlers"
)

// here I create a new http.Handler that is called MAX /multiplexer/
// I use external package from go get github.com/bmizerany/pat
func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.Home))

	return mux
}
