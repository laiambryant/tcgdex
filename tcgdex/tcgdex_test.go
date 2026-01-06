package tcgdex

import (
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
