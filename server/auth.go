package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GetToken will take the client ID and client secret environmental variables and request a token from the
// Spotify API to authenticate further requests to the API
func GetToken() (string, error) {
	// setup as shown in: https://developer.spotify.com/documentation/general/guides/authorization-guide/
	codes := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))))
	url := "https://accounts.spotify.com/api/token"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString("grant_type=client_credentials"))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", codes))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}
	token := fmt.Sprintf("%v", res["access_token"])
	if token == "" {
		return "", errors.New("the access token is nil")
	}
	return token, nil
}
