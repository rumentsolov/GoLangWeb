package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// ? Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplalte(w, "home.page.html")
}

// ? About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {

}

func renderTemplalte(w http.ResponseWriter, html5 string) {
	parsedTemplate, _ := template.ParseFiles("./html/" + html5)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error parsing template %s \n", err)
	}

}
