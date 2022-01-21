package groupie

import (
	"net/http"
	"text/template"
)

type Index struct {
	Test string
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	tplt := template.Must(template.ParseFiles("./static/artists.html"))

	APIRequest("https://groupietrackers.herokuapp.com/api/artists")
	data := ArtistsStruct{
		Tab: ArtistsTab,
	}

	tplt.Execute(w, data)
}
