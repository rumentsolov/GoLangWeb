package main

//# prints Hello Web in http://localhost:8080/

import (
	//"errors"
	"fmt"
	"net/http"
)

var portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting app on port %s \n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
