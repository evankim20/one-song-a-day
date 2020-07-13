// Package api interacts with the Spotify API and grabs and formats responses
package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/evankim20/one-song-a-day/email"
	"github.com/evankim20/one-song-a-day/spotify"
)

// getSong sends a GET request to the Spotify API to get a random song with the genre r&b and returns the full
// response as well as any errors encountered
func getSong(token string) (spotify.Response, error) {
	rand.Seed(time.Now().UnixNano())
	offset := rand.Intn(1999) + 1 // handling random offset from 1 to 2000
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=genre%%3Ar%%26b&type=track&market=US&limit=1&offset=%d", offset)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return spotify.Response{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return spotify.Response{}, err
	}
	defer res.Body.Close()

	r, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return spotify.Response{}, err
	}
	response := spotify.Response{}
	err = json.Unmarshal(r, &response)
	if err != nil {
		return spotify.Response{}, err
	}

	return response, nil
}

// MainTask will make the required calls to query a random song and send the formatted email
func MainTask() error {
	log.Printf("Sent at %v\n", time.Now())
	token, err := getToken()
	if err != nil {
		return err
	}

	resp, err := getSong(token)
	if err != nil {
		return err
	}

	err = email.Send(resp)
	if err != nil {
		return err
	}
	return nil
}
