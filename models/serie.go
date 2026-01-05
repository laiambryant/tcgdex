package models

type Serie struct {
	SerieResume
	Sets []SetResume `json:"sets"`
}
