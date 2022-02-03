package groupie

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Albums struct {
	Urls []string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tplt := template.Must(template.ParseFiles("./static/index.html"))

	APIRequest("https://groupietrackers.herokuapp.com/api/artists")
	data := ArtistsStruct{
		Tab: ArtistsTab,
	}

	CollectionImages := Albums{
		Urls: getRandomAlbum(data),
	}

	err := tplt.Execute(w, CollectionImages)
	if err != nil {
		log.Fatal(err)
	}
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	tplt := template.Must(template.ParseFiles("./static/artists.html"))

	APIRequest("https://groupietrackers.herokuapp.com/api/artists")
	data := ArtistsStruct{
		Tab: ArtistsTab,
	}

	err := tplt.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	tplt := template.Must(template.ParseFiles("./static/a.html"))

	F := Final{}

	APIRequest("https://groupietrackers.herokuapp.com/api/artists")
	data := ArtistsStruct{
		Tab: ArtistsTab,
	}

	a := r.URL.RawQuery
	a = strings.ReplaceAll(a, "a=", "")
	a = strings.ReplaceAll(a, "%20", " ")
	if r.URL.RawQuery == "" {
		http.Redirect(w, r, "/artists", http.StatusMovedPermanently)
	}

	for i := 0; i < len(data.Tab); i++ {
		if data.Tab[i].Name == a {
			F.A = Artist(data.Tab[i])
		}
	}

	GetRelations(F.A.Id)
	dates := []string{}
	locs := []string{}
	for k, v := range Rels.DatesLocations {
		for _, j := range v {
			k = strings.ReplaceAll(k, "_", " ")
			k = strings.Title(k)
			locs = append(locs, k)
			dates = append(dates, j)
		}
	}
	F.Dates = dates
	F.Locations = locs

	err := tplt.Execute(w, F)
	if err != nil {
		log.Fatal(err)
	}
}
