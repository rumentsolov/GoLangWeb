package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rumentsolov/GoLangWeb/config"
	"github.com/rumentsolov/GoLangWeb/pkg/handlers"
	"github.com/rumentsolov/GoLangWeb/render"
)

var portNumber = ":8080"
var app config.AppConfig // now its avalable for middleware
var session *scs.SessionManager // this should be transfered to middleware

func main() {
	
	// get the template cache from app config


	//change this to true when we are in production
	app.InProduction = false


	session = scs.New()
	session.Lifetime = 24 * time.Hour // how much time will live the session ( the cookie for the session)
	session.Cookie.Persist = true // says if the cookie should be persisted after someone closes the browser session
	session.Cookie.SameSite = http.SameSiteLaxMode // parameter that says to the cookie how strict the cookie should be for all it apples to
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction


	app.Session = session

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
