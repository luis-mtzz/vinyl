package models

type Release struct {
	ArtistName string   `json:"artist_name"`
	Genres     []string `json:"genres"`
	Title      string   `json:"title"`
}
