package handlers

import (
	"net/http"

	"github.com/rumentsolov/GoLangWeb/config"
	"github.com/rumentsolov/GoLangWeb/render"
)

// TemplateData is created holds data sent from handlers to templates => to pass additional data to the be rendered with the template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

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
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) { // (m *Repository) is reciever -> links all handlers togather to the Repository so all hanglers have acces to the Repository

	// perform some logic

	//send the data to the template
	render.RenderTemplate(w, "home.page.tmpl")
}

// ? About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic

	//send the data to the template
	render.RenderTemplate(w, "about.page.tmpl")

}
