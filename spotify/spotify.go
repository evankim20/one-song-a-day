// Package spotify defines the correct structures that correspond to the JSON response
// from the /search endpoint from the Spotify API
package spotify

// Response is the root structure of the json response and contains information about the request
type Response struct {
	Tracks struct {
		Href     string      `json:"href"`
		Items    []TrackInfo `json:"items"`
		Limit    int         `json:"limit"`
		Next     string      `json:"next"`
		Offset   int         `json:"offset"`
		Previous string      `json:"previous"`
		Total    int         `json:"total"`
	} `json:"tracks"`
}

// TrackInfo is the structure of the response that contains information about the current track
type TrackInfo struct {
	Album       AlbumInfo `json:"album"`
	Artists     []Artist  `json:"artists"`
	DiscNumber  int       `json:"disc_number"`
	DurationMs  int       `json:"duration_ms"`
	Explicit    bool      `json:"explicit"`
	ExternalIds struct {
		Isrc string `json:"isrc"`
	} `json:"external_ids"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href        string `json:"href"`
	ID          string `json:"id"`
	IsLocal     bool   `json:"is_local"`
	IsPlayable  bool   `json:"is_playable"`
	Name        string `json:"name"`
	Popularity  int    `json:"popularity"`
	PreviewURL  string `json:"preview_url"`
	TrackNumber int    `json:"track_number"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
}

// AlbumInfo is the structure of the response that holds info about the Album that the
// current Track belongs to
type AlbumInfo struct {
	AlbumType    string   `json:"album_type"`
	Artists      []Artist `json:"artists"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href                 string  `json:"href"`
	ID                   string  `json:"id"`
	Images               []Image `json:"images"`
	Name                 string  `json:"name"`
	ReleaseDate          string  `json:"release_date"`
	ReleaseDatePrecision string  `json:"release_date_precision"`
	TotalTracks          int     `json:"total_tracks"`
	Type                 string  `json:"type"`
	URI                  string  `json:"uri"`
}

// Artist is the structure of the rsponse that holds information about a particular artist
// for the current Track
type Artist struct {
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href string `json:"href"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

// Image is the structure that holds the attributes for images associated with the Track
type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}
