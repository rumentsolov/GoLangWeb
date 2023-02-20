package render

import (
	//"errors"
	"fmt"
	"net/http"
	"text/template"
)

// ? Renders the templates using html/template
func RenderTemplalte(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error parsing template %s \n", err)
	}

}
