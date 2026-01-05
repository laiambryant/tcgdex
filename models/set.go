package models

type Set struct {
	SetResume
	Serie SerieResume  `json:"serie"`
	Cards []CardResume `json:"cards"`
}
