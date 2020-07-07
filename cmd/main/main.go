package main

import (
	"log"

	"github.com/evankim20/one-song-a-day/server"
)

func main() {
	log.Println("hello world")
	// TODO: make new request to auth token for each call
	token, err := server.GetToken()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := server.GetSong(token)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

	// TODO: get desired fields and send email (in correct format)
	// TODO: run this every day at a certain time
}
