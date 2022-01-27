package main

import (
	g "groupie"
	"log"
	"net/http"
)

func main() {
	// Serve CSS and JavaScript
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))

	// Handle Pages Templates
	http.HandleFunc("/", g.IndexHandler)
	http.HandleFunc("/artists", g.ArtistsHandler)

	// Start the server
	err := http.ListenAndServe("localhost:80", nil)

	if err != nil {
		log.Fatal(err)
	}
}
