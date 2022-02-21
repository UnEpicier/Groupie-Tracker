package main

import (
	"fmt"
	g "groupie"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Serve CSS and JavaScript
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))

	// Handle Pages Templates
	http.HandleFunc("/", g.ErrorHandler)
	http.HandleFunc("/index", g.IndexHandler)
	http.HandleFunc("/artists", g.ArtistsHandler)
	http.HandleFunc("/artist", g.ArtistHandler)

	ip := "localhost" // Allow every devices
	port := 80

	// Start the server
	fmt.Println("Starting local server...")
	fmt.Printf("Server available at:\n\t- http://localhost:%d\n\t- http://%s:%d\n", port, g.GetLocalIP(), port)
	err := http.ListenAndServe(ip+":"+strconv.Itoa(port), nil)

	if err != nil {
		log.Fatal(err)
	}
}
