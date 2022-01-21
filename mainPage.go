package groupie

import (
	"net/http"
	"text/template"
)

type Index struct {
	Test string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tplt := template.Must(template.ParseFiles("static/index.html"))

	APIRequest()
	data := ArtistsStruct{
		Tab: ArtistsTab,
	}

	tplt.Execute(w, data)

}
