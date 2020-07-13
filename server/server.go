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
	// TODO: CHANGE THIS TO 8 AM --> 15:00
	lower := time.Date(now.Year(), now.Month(), now.Day(), 19, 0, 0, 0, time.UTC)
	upper := time.Date(now.Year(), now.Month(), now.Day(), 19, 30, 0, 0, time.UTC)
	if now.After(lower) && now.Before(upper) {
		api.MainTask()
	}
	fmt.Fprintln(w, "One Song a Day")
}
