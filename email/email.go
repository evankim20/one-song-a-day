// Package email contains funcitonality to interact and send emails to clients
package email

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/evankim20/one-song-a-day/spotify"
	"gopkg.in/gomail.v2"
)

// Send will take in a response from the Spotify API for a track and send an email with
// that track to the pre-set email from the pre-set song email
func Send(r spotify.Response) error {
	type TemplateData struct {
		Name   string
		Image  string
		Artist string
		Track  string
		URI    string
	}
	track := r.Tracks.Items[0]
	from := os.Getenv("SONG_EMAIL")
	pass := os.Getenv("SONG_EMAIL_PASS")
	to := os.Getenv("PERSONAL_EMAIL")

	data := TemplateData{
		Name:   os.Getenv("CURRENT_USER"),
		Image:  track.Album.Images[1].URL,
		Artist: track.Artists[0].Name,
		Track:  track.Name,
		URI:    strings.Split(track.URI, ":")[2],
	}

	t, err := template.ParseFiles("templates/template.html")
	if err != nil {
		return err
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return err
	}
	result := tpl.String()

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", fmt.Sprintf("%v:%s", time.Now().Format("01-02-2006"), track.Name))
	m.SetBody("text/html", result)

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, from, pass)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
