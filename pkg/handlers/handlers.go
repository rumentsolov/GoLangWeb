package render

import (
	"net/http"

	"github.com/rumentsolov/GoLangWeb/render"
)

// ? Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplalte(w, "home.page.tmpl")
}

// ? About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplalte(w, "about.page.tmpl")


}
