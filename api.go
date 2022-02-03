package groupie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var ArtistsTab []Artists

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
