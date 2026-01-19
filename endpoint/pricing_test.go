package endpoint

import (
	"context"
	"net/http"
	"testing"

	"github.com/laiambryant/tcgdex/client"
	"github.com/laiambryant/tcgdex/models"
)

func TestGetCardWithPricing(t *testing.T) {
	c := client.NewHTTPClient(&fakeHTTP{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.String() != "http://example/cards/swsh1-1" {
			t.Fatalf("unexpected url: %s", req.URL.String())
		}
		body := `{
			"id": "swsh1-1",
			"localId": "1",
			"name": "Test Card",
			"pricing": {
				"cardmarket": {
					"updated": "2025-08-05T00:42:15.000Z",
					"unit": "EUR",
					"avg": 0.08,
					"avg-holo": 0.27
				},
				"tcgplayer": {
					"updated": "2025-08-05T20:07:54.000Z",
					"unit": "USD",
					"normal": {
						"marketPrice": 0.09
					}
				}
			}
		}`
		return client.NewMockResponse(200, body), nil
	}}, client.WithBaseURL("http://example"))
	endpoint := New[models.Card, models.CardResume](c, "cards")
	card, err := endpoint.Get(context.Background(), "swsh1-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if card.Pricing == nil || card.Pricing.Cardmarket == nil {
		t.Fatalf("expected cardmarket pricing to be present")
	}
	if card.Pricing.Cardmarket.AvgHolo == nil || *card.Pricing.Cardmarket.AvgHolo != 0.27 {
		t.Fatalf("expected avg-holo to be 0.27, got %#v", card.Pricing.Cardmarket.AvgHolo)
	}
	if card.Pricing.TCGPlayer == nil || card.Pricing.TCGPlayer.Normal == nil {
		t.Fatalf("expected tcgplayer normal pricing to be present")
	}
	if card.Pricing.TCGPlayer.Normal.MarketPrice == nil || *card.Pricing.TCGPlayer.Normal.MarketPrice != 0.09 {
		t.Fatalf("unexpected marketPrice: %#v", card.Pricing.TCGPlayer.Normal.MarketPrice)
	}
}
