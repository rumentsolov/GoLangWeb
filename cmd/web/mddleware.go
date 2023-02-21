package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{ //it uses cookies that token it generates is avalable on the server
		HttpOnly: true,
		Path:     "/",                  //refers to entier website for a cookie path
		Secure:   false,                // we are not running in SSL mode
		SameSite: http.SameSiteLaxMode, // built in by standard

	})

	return csrfHandler
}
