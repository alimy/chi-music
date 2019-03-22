package models

import "net/http"

// Album indicate album info
type Album struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseYear string `json:"releaseYear"`
	Genre       string `json:"genre"`
	TrackCount  int    `json:"-"`
	AlbumId     string `json:"albumId,omitempty"`
}

// Albums indicate album slice
type Albums []*Album

// Render implement render.Render(...)
func (*Album) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render implement render.Render(...)
func (Albums) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
