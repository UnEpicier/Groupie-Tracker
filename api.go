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

func APIRequest(url string) int {
	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return 1
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
		return 1
	}

	json.Unmarshal(data, &ArtistsTab)
	return 0
}
