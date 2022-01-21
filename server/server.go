package main

import (
	g "groupie"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", fs)
	http.HandleFunc("/artists", g.ArtistsHandler)

	err := http.ListenAndServe("localhost:80", nil)

	if err != nil {
		log.Fatal(err)
	}
}
