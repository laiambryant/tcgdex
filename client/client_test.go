package client

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestClient_Get(t *testing.T) {
	mockRT := &MockRoundTripper{
		RoundTripFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.Path == "/test" {
				return NewMockResponse(200, `{"ok": true}`), nil
			}
			if req.URL.Path == "/404" {
				return NewMockResponse(404, "Not Found"), nil
			}
			return NewMockResponse(500, "Error"), nil
		},
	}

	c := NewHTTPClient(&http.Client{Transport: mockRT}, WithBaseURL("http://example.com"))

	t.Run("Success", func(t *testing.T) {
		data, err := c.Get(context.Background(), "/test")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if string(data) != `{"ok": true}` {
			t.Errorf("expected body %s, got %s", `{"ok": true}`, string(data))
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		_, err := c.Get(context.Background(), "/404")
		if err != ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", err)
		}
	})

	t.Run("Cache", func(t *testing.T) {
		cWithCache := NewHTTPClient(&http.Client{Transport: mockRT}, WithBaseURL("http://example.com"), WithCache(time.Minute))

		_, err := cWithCache.Get(context.Background(), "/test")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		mockRT.RoundTripFunc = func(req *http.Request) (*http.Response, error) {
			return NewMockResponse(500, "Should not be called"), nil
		}

		data, err := cWithCache.Get(context.Background(), "/test")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if string(data) != `{"ok": true}` {
			t.Errorf("expected cached body, got %s", string(data))
		}
	})
}
