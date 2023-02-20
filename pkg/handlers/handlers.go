package handlers

import (
	"github.com/rumentsolov/GoLangWeb/render"
	"net/http"
)

// ? Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplalte(w, "home.page.html") 
}

// ? About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplalte(w, "about.page.html")
}
