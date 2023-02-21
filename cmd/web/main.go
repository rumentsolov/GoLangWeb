package main

//# prints Hello Web in http://localhost:8080/

import (
	//"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/rumentsolov/GoLangWeb/config"
	"github.com/rumentsolov/GoLangWeb/pkg/handlers"
	"github.com/rumentsolov/GoLangWeb/render"
)

var portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCacheA()
	if err != nil {
		log.Fatal("Cannot create template cach")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting app on port %s \n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
