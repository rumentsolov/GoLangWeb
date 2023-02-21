package render

import (
	"net/http"
)

func RenderTemplalte(w http.ResponseWriter, tmpl string) {

	var auto = true

	if auto {
		RenderTemplalteA(w, tmpl)
	} else {
		RenderTemplalteM(w, tmpl)
	}
}
