package groupie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var ArtistsTab []Artists
var FullArtistsTab []Artists

func APIRequest(url string) {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &ArtistsTab)
	json.Unmarshal(data, &FullArtistsTab)
}

var Rels = Relations{}

func GetRelations(id int) {
	Rels = Relations{}

	req, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(id))
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &Rels)
}

var M []Map

func GetMap(address string) {
	M = []Map{}
	address = strings.ToLower(address)
	address = strings.ReplaceAll(address, " ", "+")

	req, err := http.Get("https://nominatim.openstreetmap.org/search?q=" + address + "&format=json")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &M)
}
