package main

import (
	g "groupie"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", g.MainHandler)

	err := http.ListenAndServe("localhost:80", nil)

	if err != nil {
		log.Fatal(err)
	}
}
