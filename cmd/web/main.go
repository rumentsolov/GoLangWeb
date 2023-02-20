package main

//# prints Hello Web in http://localhost:8080/

import (
	//"errors"
	"fmt"
	"net/http"

	"github.com/rumentsolov/GoLangWeb/pkg/handlers"
)

var portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting app on port %s \n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
