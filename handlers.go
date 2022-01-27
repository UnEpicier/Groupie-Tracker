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

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	tplt := template.Must(template.ParseFiles("./static/a.html"))

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

	artist := Artist{}

	for i := 0; i < len(data.Tab); i++ {
		if data.Tab[i].Name == a {
			artist = Artist(data.Tab[i])
		}
	}

	err := tplt.Execute(w, artist)
	if err != nil {
		log.Fatal(err)
	}
}
