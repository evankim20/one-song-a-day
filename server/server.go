package server

import (
	"fmt"
	"net/http"
	"os"
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
	fmt.Fprintln(w, "One Song a Day")
}
