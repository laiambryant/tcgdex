package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestPricingUnmarshal(t *testing.T) {
	payload := `{
		"cardmarket": {
			"updated": "2025-08-05T00:42:15.000Z",
			"unit": "EUR",
			"avg": 0.08,
			"low": 0.02,
			"trend": 0.08,
			"avg1": 0.03,
			"avg7": 0.08,
			"avg30": 0.08,
			"avg-holo": 0.27,
			"low-holo": 0.03,
			"trend-holo": 0.21
		},
		"tcgplayer": {
			"updated": "2025-08-05T20:07:54.000Z",
			"unit": "USD",
			"normal": {
				"lowPrice": 0.02,
				"midPrice": 0.17,
				"highPrice": 25.09,
				"marketPrice": 0.09,
				"directLowPrice": 0.04
			},
			"reverse": {
				"lowPrice": 0.09,
				"midPrice": 0.26,
				"highPrice": 5.17,
				"marketPrice": 0.23,
				"directLowPrice": 0.23
			}
		}
	}`

	var p Pricing
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		t.Fatalf("unexpected error unmarshaling pricing: %v", err)
	}
	if p.Cardmarket == nil {
		t.Fatalf("expected cardmarket pricing to be present")
	}
	if p.TCGPlayer == nil {
		t.Fatalf("expected tcgplayer pricing to be present")
	}
	if p.Cardmarket.AvgHolo == nil || *p.Cardmarket.AvgHolo != 0.27 {
		t.Fatalf("expected avg-holo to be 0.27, got %#v", p.Cardmarket.AvgHolo)
	}
	if p.TCGPlayer.Normal == nil || p.TCGPlayer.Normal.MarketPrice == nil || *p.TCGPlayer.Normal.MarketPrice != 0.09 {
		t.Fatalf("expected normal marketPrice to be 0.09, got %#v", p.TCGPlayer.Normal)
	}
	wantCM, err := time.Parse(time.RFC3339Nano, "2025-08-05T00:42:15.000Z")
	if err != nil {
		t.Fatalf("unexpected time parse error: %v", err)
	}
	if p.Cardmarket.Updated == nil || !p.Cardmarket.Updated.Equal(wantCM) {
		t.Fatalf("unexpected cardmarket updated time: %#v", p.Cardmarket.Updated)
	}
}

func TestPricingUnmarshalMissingProvider(t *testing.T) {
	payload := `{
		"cardmarket": {
			"updated": "2025-08-05T00:42:15.000Z",
			"unit": "EUR",
			"avg": 0.08
		}
	}`

	var p Pricing
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		t.Fatalf("unexpected error unmarshaling pricing: %v", err)
	}
	if p.Cardmarket == nil {
		t.Fatalf("expected cardmarket pricing to be present")
	}
	if p.TCGPlayer != nil {
		t.Fatalf("expected tcgplayer pricing to be nil when missing")
	}
}
