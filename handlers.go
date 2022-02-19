package groupie

import (
	"log"
	"net/http"
	"sort"
	"strconv"
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
		Tab:       ArtistsTab,
		CreaDates: []int{},
		AlbDates:  []int{},
		Locations: []string{},
	}

	for _, v := range data.Tab {
		data.CreaDates = append(data.CreaDates, v.CreationDate)

		tempAl, _ := strconv.Atoi(strings.Split(v.FirstAlbum, "-")[2])
		data.AlbDates = append(data.AlbDates, tempAl)
	}
	sort.Ints(data.CreaDates)
	sort.Ints(data.AlbDates)

	data.Locations = getLocations(data)

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		minDateC, _ := strconv.Atoi(r.FormValue("minDateC"))
		maxDateC, _ := strconv.Atoi(r.FormValue("maxDateC"))

		minDateA, _ := strconv.Atoi(r.FormValue("minDateA"))
		maxDateA, _ := strconv.Atoi(r.FormValue("maxDateA"))

		//location := r.FormValue("location")

		for i := 0; i < len(data.Tab); i++ {
			dA, _ := strconv.Atoi(strings.Split(data.Tab[i].FirstAlbum, "-")[2])

			if !((data.Tab[i].CreationDate >= minDateC && data.Tab[i].CreationDate <= maxDateC) && (dA >= minDateA && dA <= maxDateA)) {
				data.Tab = removeArtist(data.Tab, i)
				i--
			}
		}
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

	for i := 0; i < len(locs); i++ {
		if contains(locs[i+1:], locs[i]) {
			locs = remove(locs, i)
			i--
		}
	}

	coords := map[string][]string{}
	for _, v := range locs {
		GetMap(v)
		coords[v] = []string{M[0].Lat, M[0].Lon}
	}

	F.Coords = coords

	err := tplt.Execute(w, F)
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		IndexHandler(w, r)
		return
	}

	tplt := template.Must(template.ParseFiles("./static/error.html"))

	err := tplt.Execute(w, tplt)
	if err != nil {
		log.Fatal(err)
	}
}
