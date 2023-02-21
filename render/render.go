package render

import (
	"net/http"
)

func RenderTemplalte(w http.ResponseWriter, tmpl string) {

	var auto = false

	if auto {
		RenderTemplalteA(*w, *tmpl)
	} else {
		RenderTemplalteM(*w, *tmpl)
	}
}
