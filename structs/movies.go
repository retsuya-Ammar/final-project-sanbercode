package structs

import "time"

type Movies struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Category      string    `json:"category"`
	Video_url     string    `json:"video_url"`
	Thumbnail_url string    `json:"thumbnail"`
	Rating        float64   `json:"rating"`
	Is_featured   bool      `json:"is_featured"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}
