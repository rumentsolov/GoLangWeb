package handlers

import (
	"net/http"

	"github.com/rumentsolov/GoLangWeb/config"
	"github.com/rumentsolov/GoLangWeb/render"
)

// we use reposatory patthern

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// ? Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplalte(w, "home.page.tmpl")
}

// ? About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplalte(w, "about.page.tmpl")

}
