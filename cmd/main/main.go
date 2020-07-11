package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/evankim20/one-song-a-day/api"
	"github.com/evankim20/one-song-a-day/email"
	"github.com/evankim20/one-song-a-day/server"
)

var ticker *time.Ticker = nil

func main() {
	// server set up
	addr, err := server.DetermineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", server.RootHandler)
	log.Printf("Listening on %s...\n", addr)
	go http.ListenAndServe(addr, nil)

	time.AfterFunc(duration(), mainTask)
	wg.Add(1)
	// do normal task here
	wg.Wait()
}

func mainTask() {
	if ticker == nil {
		ticker = time.NewTicker(24 * time.Second)
	}
	for {
		log.Printf("Sent at %v\n", time.Now())
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
		<-ticker.C
	}
}

func duration() time.Duration {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), 21, 59, 0, 0, t.Location())
	if t.After(n) {
		n = n.Add(24 * time.Hour)
	}
	d := n.Sub(t)
	return d
}

var wg sync.WaitGroup
