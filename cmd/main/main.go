package main

import (
	"fmt"
	"log"

	"github.com/evankim20/one-song-a-day/server"
)

func main() {
	log.Println("hello world")
	token, err := server.GetToken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
}
