package models

import "time"

// Pricing represents pricing data returned by the API for a card.
type Pricing struct {
	Cardmarket *CardmarketPricing `json:"cardmarket,omitempty"`
	TCGPlayer  *TCGPlayerPricing  `json:"tcgplayer,omitempty"`
}

// CardmarketPricing contains Cardmarket marketplace statistics.
type CardmarketPricing struct {
	Updated *time.Time `json:"updated,omitempty"`
	Unit    string     `json:"unit,omitempty"`

	Avg   *float64 `json:"avg,omitempty"`
	Low   *float64 `json:"low,omitempty"`
	Trend *float64 `json:"trend,omitempty"`
	Avg1  *float64 `json:"avg1,omitempty"`
	Avg7  *float64 `json:"avg7,omitempty"`
	Avg30 *float64 `json:"avg30,omitempty"`

	AvgHolo          *float64 `json:"avg-holo,omitempty"`
	LowHolo          *float64 `json:"low-holo,omitempty"`
	TrendHolo        *float64 `json:"trend-holo,omitempty"`
	AvgReverseHolo   *float64 `json:"avg-reverse-holo,omitempty"`
	LowReverseHolo   *float64 `json:"low-reverse-holo,omitempty"`
	TrendReverseHolo *float64 `json:"trend-reverse-holo,omitempty"`
}

// TCGPlayerPricing contains TCGPlayer marketplace data with per-variant pricing.
type TCGPlayerPricing struct {
	Updated *time.Time `json:"updated,omitempty"`
	Unit    string     `json:"unit,omitempty"`

	Normal  *TCGPlayerPriceVariant `json:"normal,omitempty"`
	Reverse *TCGPlayerPriceVariant `json:"reverse,omitempty"`
}

// TCGPlayerPriceVariant contains price points for a specific TCGPlayer variant.
type TCGPlayerPriceVariant struct {
	LowPrice       *float64 `json:"lowPrice,omitempty"`
	MidPrice       *float64 `json:"midPrice,omitempty"`
	HighPrice      *float64 `json:"highPrice,omitempty"`
	MarketPrice    *float64 `json:"marketPrice,omitempty"`
	DirectLowPrice *float64 `json:"directLowPrice,omitempty"`
}
