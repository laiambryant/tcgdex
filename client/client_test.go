package client

import (
	"context"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"
)

type errReadCloser struct{}

func (errReadCloser) Read([]byte) (int, error) { return 0, errors.New("boom") }
func (errReadCloser) Close() error             { return nil }

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

	t.Run("CacheExpiry", func(t *testing.T) {
		cWithCache := NewHTTPClient(&http.Client{Transport: mockRT}, WithBaseURL("http://example.com"), WithCache(10*time.Millisecond))
		prev := mockRT.RoundTripFunc
		defer func() { mockRT.RoundTripFunc = prev }()
		mockRT.RoundTripFunc = func(req *http.Request) (*http.Response, error) {
			if req.URL.Path == "/expire" {
				return NewMockResponse(200, `{"ok": true}`), nil
			}
			return NewMockResponse(500, "Error"), nil
		}

		data, err := cWithCache.Get(context.Background(), "/expire")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if string(data) != `{"ok": true}` {
			t.Fatalf("expected body %s, got %s", `{"ok": true}`, string(data))
		}
		mockRT.RoundTripFunc = func(req *http.Request) (*http.Response, error) {
			if req.URL.Path == "/expire" {
				return NewMockResponse(200, `{"ok": false}`), nil
			}
			return NewMockResponse(500, "Error"), nil
		}
		time.Sleep(20 * time.Millisecond)
		data, err = cWithCache.Get(context.Background(), "/expire")
		if err != nil {
			t.Fatalf("expected no error after expiry, got %v", err)
		}
		if string(data) != `{"ok": false}` {
			t.Fatalf("expected refreshed body after expiry, got %s", string(data))
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

func TestClient_ErrorsAndDownload(t *testing.T) {
	mockRT := &MockRoundTripper{}

	c := NewHTTPClient(&http.Client{Transport: mockRT}, WithBaseURL("http://example.com"))

	t.Run("DoRequestError", func(t *testing.T) {
		mockRT.RoundTripFunc = func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("net")
		}
		_, err := c.Get(context.Background(), "/x")
		var re *RequestError
		if !errors.As(err, &re) {
			t.Fatalf("expected RequestError, got %v", err)
		}
	})

	t.Run("GetCreateRequestError", func(t *testing.T) {
		c2 := NewHTTPClient(nil, WithBaseURL("\x00"))
		_, err := c2.Get(context.Background(), "/x")
		var re *RequestError
		if !errors.As(err, &re) || re.Op != "create request" {
			t.Fatalf("expected create RequestError, got %v", err)
		}
	})

	t.Run("GetReadBodyError", func(t *testing.T) {
		mockRT.RoundTripFunc = func(req *http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: errReadCloser{}, Header: make(http.Header)}, nil
		}
		_, err := c.Get(context.Background(), "/x")
		var re *RequestError
		if !errors.As(err, &re) || re.Op != "read body" {
			t.Fatalf("expected read body RequestError, got %v", err)
		}
	})

	t.Run("GetHTTPError", func(t *testing.T) {
		mockRT.RoundTripFunc = func(req *http.Request) (*http.Response, error) {
			return NewMockResponse(500, "bad"), nil
		}
		_, err := c.Get(context.Background(), "/x")
		var he *HTTPError
		if !errors.As(err, &he) || he.Status != 500 || he.Body != "bad" {
			t.Fatalf("expected HTTPError 500 bad, got %v", err)
		}
	})

	t.Run("DownloadSuccess", func(t *testing.T) {
		mockRT.RoundTripFunc = func(req *http.Request) (*http.Response, error) {
			return NewMockResponse(200, "payload"), nil
		}
		rc, err := c.Download(context.Background(), "http://ex/p")
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		b, _ := io.ReadAll(rc)
		rc.Close()
		if string(b) != "payload" {
			t.Fatalf("bad body %s", string(b))
		}
	})

	t.Run("DownloadNotFound", func(t *testing.T) {
		mockRT.RoundTripFunc = func(req *http.Request) (*http.Response, error) {
			return NewMockResponse(404, "no"), nil
		}
		_, err := c.Download(context.Background(), "u")
		if err != ErrNotFound {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})

	t.Run("DownloadDoRequestError", func(t *testing.T) {
		mockRT.RoundTripFunc = func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("net")
		}
		_, err := c.Download(context.Background(), "http://ex/p")
		var re *RequestError
		if !errors.As(err, &re) || re.Op != "do request" {
			t.Fatalf("expected do request RequestError, got %v", err)
		}
	})

	t.Run("DownloadCreateRequestError", func(t *testing.T) {
		_, err := c.Download(context.Background(), "\x00")
		var re *RequestError
		if !errors.As(err, &re) || re.Op != "create request" {
			t.Fatalf("expected create request RequestError, got %v", err)
		}
	})
	t.Run("DownloadHTTPError", func(t *testing.T) {
		mockRT.RoundTripFunc = func(req *http.Request) (*http.Response, error) {
			return NewMockResponse(502, "srverr"), nil
		}
		_, err := c.Download(context.Background(), "u")
		var he *HTTPError
		if !errors.As(err, &he) || he.Status != 502 || he.Body != "srverr" {
			t.Fatalf("expected HTTPError 502 srverr, got %v", err)
		}
	})
}

func TestOptionsAndErrorUnwrap(t *testing.T) {
	cli := NewHTTPClient(nil, WithUserAgent("UA"), WithBaseURL("B"), WithCache(time.Minute))
	if cli.UserAgent != "UA" || cli.BaseURL != "B" || cli.cache == nil {
		t.Fatalf("options not applied %v", cli)
	}
	he := &HTTPError{Status: 1, URL: "u", Body: "b", Cause: errors.New("c")}
	if he.Error() == "" || !errors.Is(he, he.Cause) {
		t.Fatalf("http error unwrap bad %v", he)
	}
	re := &RequestError{Op: "o", Err: errors.New("e")}
	if re.Error() == "" || !errors.Is(re, re.Err) {
		t.Fatalf("request error unwrap bad %v", re)
	}
}

type impl struct{}

func (impl) Do(req *http.Request) (*http.Response, error) {
	return NewMockResponse(200, "ok"), nil
}

func TestWithHTTPClient(t *testing.T) {
	cli := NewHTTPClient(impl{}, WithHTTPClient(impl{}))
	if cli.HTTP == nil {
		t.Fatal("http client not set")
	}
}
