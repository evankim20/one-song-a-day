package main

import (
	"log"

	"github.com/evankim20/one-song-a-day/api"
	"github.com/evankim20/one-song-a-day/email"
)

func main() {
	// TODO: make new request to auth token for each call
	token, err := api.GetToken()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := api.GetSong(token)
	if err != nil {
		log.Fatal(err)
	}

	err = email.Send(resp)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: run this every day at a certain time
}
