package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

//! AUTOMATIC TEMPLATE CACH SYSTEM
//* its not necessary to track all my files and include them for rendering all the time
//? avoiding loading the entire template cache everytime we display the page in website => this is problem due to the information is parsing all the time from the hard drive to the screen every time

// Renders the templates using html/template
func RenderTemplalteA(w http.ResponseWriter, tmpl string) {
	//create template cache variable
	tc, err := CreateTemplateCacheA()
	ErrorCheck(err)

	//get requested template from cach
	t, ok := tc[tmpl] // t is the value of the template I want to render, ok is true of false
	if !ok {
		log.Fatal(err) // if I dont have the template I want to render just DIE MOTHERFCKER ^^
	}

	buf := new(bytes.Buffer)  // this arbiter step that usually doesn needed. I make buffer that hold bytes and trying to execute the buffer directly and then to write it out. Final grinf error checking!
	err = t.Execute(buf, nil) // this gives me clear indication of if I there is something wrong with my valye tooked from the map
	ErrorCheck(err)

	log.Println("Beggining AUTO rendering!")
	//renderer template
	_, err = buf.WriteTo(w)
	ErrorCheck(err)
}

func CreateTemplateCacheA() (map[string]*template.Template, error) {

	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{} // exactly the same as bellow statement

	// the idea is to fill all the map with the info from all FileSystem in root folder of the projects! When you rendering template that uses layout then you must have as a first thing to Parse the template you want to render and then the associated layot and parials and so on... that means that everything ending with .page , .tmpl and .html should go first. Thats why I need to check the whole file system and find those files
	//? get all the files *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl") // filepath is the package name
	ErrorCheck(err)

	// range trough all the files ending with .page.tmpl
	for _, page := range pages {
		// to extract the name of the template
		name := filepath.Base(page) // strips the path from the filename
		ts, err := template.New(name).ParseFiles(page)
		ErrorCheck(err)
		// now when we did that we have to look for any layouts in our directories
		matches, err := filepath.Glob("./templates/*layout.tmpl")
		ErrorCheck(err)

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl") // ParseGlob parses the template definitions in the files identified by the pattern and associates the resulting templates with ts.
			ErrorCheck(err)
		}
		myCache[name] = ts // adding to the map
	}

	return myCache, nil

}
