package models

type SetResume struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Logo      *string      `json:"logo,omitempty"`
	Symbol    *string      `json:"symbol,omitempty"`
	CardCount SetCardCount `json:"cardCount"`
}
