package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{ //it uses cookies that token it generates is avalable on the server
		HttpOnly: true,
		Path:     "/",                  //refers to entier website for a cookie path
		Secure:   app.InProduction,     // we are not running in SSL mode
		SameSite: http.SameSiteLaxMode, // built in by standard

	})

	return csrfHandler
}

// SessionLoad loads nad saves the session on every request / all its doing is to load the session
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next) // provide middleware for session
}
