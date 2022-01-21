package groupie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ArtistsStruct struct {
	Tab []Artists
}

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

var ArtistsTab []Artists

func APIRequest() {
	req, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &ArtistsTab)

}
