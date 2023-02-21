package main

import (
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
	// get the template cache from app config

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cach")
	}

	app.TemplateCache = tc
	app.UseCache = false

	// passing repository from handlers and back to handlers
	repo := handlers.NewRepo(&app) // creates the repository variables
	// I pass it back to the handlers
	handlers.NewHandler(repo)

	render.NewTemplate(&app) // this gives our application access to the the app config

	// .Repo. because I have reciever from type REPO
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	
	fmt.Printf("Starting app on port %s \n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

	//_ = http.ListenAndServe(portNumber, nil)
}
