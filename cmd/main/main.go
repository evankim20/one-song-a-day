package main

import (
	"log"
	"net/http"

	"github.com/evankim20/one-song-a-day/server"
)

func main() {
	// server set up
	addr, err := server.DetermineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", server.RootHandler)
	log.Printf("Listening on %s...\n", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
