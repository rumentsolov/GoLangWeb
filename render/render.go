package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

/*
// Renders the templates using html/template
func RenderTemplalte(w http.ResponseWriter, tmpl string) {
	log.Println("beggining rendering")
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	//! add one entry for any template file I create from now on
	err := parsedTemplate.Execute(w, nil)
	ErrorCheck(err)
}
*/
//! SIMPLE TEMPLATE CACH SYSTEM

// Template Cach
var tc = make(map[string]*template.Template)

// instead of reading the template from the disk we are storing into the variable. The variable has to be avalable every time I call this function
func RenderTemplalte(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//* check to see if we have already the template in our cache [if we have string t in our map currently]
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("creating template& adding to cache")
		err = createTemplateCache(t)
		ErrorCheck(err)
	} else {
		// we have template in the cach
		log.Println("using template cache")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	ErrorCheck(err)
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	//* parse the template -> it takes each entry from t string and puts it in individual strings
	tmpll, err := template.ParseFiles(templates...)
	ErrorCheck(err)

	//* add the template to cache
	tc[t] = tmpll
	return nil
}
