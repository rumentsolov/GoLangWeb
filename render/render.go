package render

import (
	//"errors"
	"fmt"
	"net/http"
	"text/template"
)

//? Renders the templates using html/template
func RenderTemplalte(w http.ResponseWriter, html5 string) {
	parsedTemplate, _ := template.ParseFiles("./html/" + html5)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error parsing template %s \n", err)
	}

}
