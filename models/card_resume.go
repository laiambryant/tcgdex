package models

import (
	"fmt"

	"github.com/laiambryant/tcgdex/enums"
)

type CardResume struct {
	ID      string  `json:"id"`
	LocalID string  `json:"localId"`
	Name    string  `json:"name"`
	Image   *string `json:"image,omitempty"`
}

func (c *CardResume) GetImageURL(quality enums.Quality, extension enums.Extension) *string {
	if c.Image == nil {
		return nil
	}
	url := fmt.Sprintf("%s/%s.%s", *c.Image, quality, extension)
	return &url
}
