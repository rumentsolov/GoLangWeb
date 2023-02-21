package handlers

import (
	"net/http"

	"github.com/rumentsolov/GoLangWeb/config"
	"github.com/rumentsolov/GoLangWeb/models"
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
// (m *Repository) is reciever -> links all handlers togather to the Repository so all hanglers have acces to the Repository
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	//* getting the user's ip address {who is visiting my website}
	remoteIP := r.RemoteAddr
	//* now I put the user's ip in a session
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	//* everytime somebody hits the page I store his ip addr

	//send the data to the template
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// ? About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP // pulling remote_ip out of the session

	//send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})

}
