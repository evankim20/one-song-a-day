package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/evankim20/one-song-a-day/api"
)

// DetermineListenAddress will try to get the assigned Port or throw an error
func DetermineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

// RootHandler handles the path /
func RootHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	lower := time.Date(now.Year(), now.Month(), now.Day(), 15, 0, 0, 0, time.UTC)
	upper := time.Date(now.Year(), now.Month(), now.Day(), 15, 4, 0, 0, time.UTC)
	if now.After(lower) && now.Before(upper) {
		api.MainTask()
		fmt.Fprintln(w, "Just sent a request!")
		return
	}
	fmt.Fprintln(w, "Not the correct time")
}
