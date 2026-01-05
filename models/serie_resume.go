package models

type SerieResume struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Logo *string `json:"logo,omitempty"`
}
