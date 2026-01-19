package tcgdex

import (
	"context"
	"net/http"
	"testing"

	"github.com/laiambryant/tcgdex/client"
)

type fakeHTTPClient struct{}

func (f *fakeHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return client.NewMockResponse(200, `[]`), nil
}

func TestNewDefaultsAndEndpoints(t *testing.T) {
	sdk := New()
	if sdk.Client == nil {
		t.Fatalf("expected Client to be non-nil")
	}
	if sdk.Client.BaseURL != "https://api.tcgdex.net/v2/en" {
		t.Fatalf("unexpected default BaseURL: %s", sdk.Client.BaseURL)
	}
	if sdk.Card == nil || sdk.Card.Path != "cards" {
		t.Fatalf("Card endpoint not initialized correctly: %#v", sdk.Card)
	}
	if sdk.Set == nil || sdk.Set.Path != "sets" {
		t.Fatalf("Set endpoint not initialized correctly: %#v", sdk.Set)
	}
	if sdk.Serie == nil || sdk.Serie.Path != "series" {
		t.Fatalf("Serie endpoint not initialized correctly: %#v", sdk.Serie)
	}
	if sdk.Card.Client != sdk.Client || sdk.Set.Client != sdk.Client || sdk.Serie.Client != sdk.Client {
		t.Fatalf("endpoints should share the same client instance")
	}
}

func TestNewWithOptionsOverrides(t *testing.T) {
	fake := &fakeHTTPClient{}
	sdk := New(client.WithBaseURL("http://example"), client.WithUserAgent("custom-agent"), client.WithHTTPClient(fake))
	if sdk.Client.BaseURL != "http://example" {
		t.Fatalf("expected base url override, got %s", sdk.Client.BaseURL)
	}
	if sdk.Client.UserAgent != "custom-agent" {
		t.Fatalf("expected user agent override, got %s", sdk.Client.UserAgent)
	}
	if sdk.Client.HTTP != fake {
		t.Fatalf("expected provided HTTP client to be used")
	}
}

type fakeHTTPPricingClient struct{}

func (f *fakeHTTPPricingClient) Do(req *http.Request) (*http.Response, error) {
	if req.URL.String() != "http://example/cards/swsh1-1" {
		return client.NewMockResponse(404, `{"error":"not found"}`), nil
	}
	body := `{
		"id": "swsh1-1",
		"localId": "1",
		"name": "Test Card",
		"pricing": {
			"cardmarket": {
				"updated": "2025-08-05T00:42:15.000Z",
				"unit": "EUR",
				"avg": 0.08
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
}

func TestGetCardWithPricing(t *testing.T) {
	sdk := New(client.WithBaseURL("http://example"), client.WithHTTPClient(&fakeHTTPPricingClient{}))
	card, err := sdk.GetCardWithPricing(context.Background(), "swsh1-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if card.Pricing == nil || card.Pricing.Cardmarket == nil {
		t.Fatalf("expected cardmarket pricing to be present")
	}
	if card.Pricing.TCGPlayer == nil || card.Pricing.TCGPlayer.Normal == nil {
		t.Fatalf("expected tcgplayer normal pricing to be present")
	}
}
