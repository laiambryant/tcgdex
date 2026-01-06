package endpoint

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/laiambryant/tcgdex/client"
	"github.com/laiambryant/tcgdex/query"
)

type fakeHTTP struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (f *fakeHTTP) Do(req *http.Request) (*http.Response, error) { return f.fn(req) }

func TestNewEndpoint(t *testing.T) {
	c := client.NewHTTPClient(&fakeHTTP{fn: func(req *http.Request) (*http.Response, error) { return client.NewMockResponse(200, `[]`), nil }}, client.WithBaseURL("http://example"))
	e := New[struct{}, struct{}](c, "cards")
	if e.Client != c {
		t.Fatalf("expected client to be set")
	}
	if e.Path != "cards" {
		t.Fatalf("expected path set, got %q", e.Path)
	}
}

func TestGetSuccessAndDecodeErrorAndRequestError(t *testing.T) {
	type Item struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	// success
	c := client.NewHTTPClient(&fakeHTTP{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.String() != "http://example/cards/123" {
			t.Fatalf("unexpected url: %s", req.URL.String())
		}
		return client.NewMockResponse(200, `{"id":"123","name":"bob"}`), nil
	}}, client.WithBaseURL("http://example"))
	e := New[Item, Item](c, "cards")
	it, err := e.Get(context.Background(), "123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if it.ID != "123" || it.Name != "bob" {
		t.Fatalf("unexpected item: %#v", it)
	}

	// decode error
	c2 := client.NewHTTPClient(&fakeHTTP{fn: func(req *http.Request) (*http.Response, error) {
		return client.NewMockResponse(200, `not-json`), nil
	}}, client.WithBaseURL("http://example"))
	e2 := New[Item, Item](c2, "cards")
	_, err = e2.Get(context.Background(), "123")
	if err == nil {
		t.Fatalf("expected decode error")
	}
	var derr *DecodeError
	if !errors.As(err, &derr) {
		t.Fatalf("expected DecodeError, got %T", err)
	}
	if !strings.Contains(derr.Error(), "/cards/123") {
		t.Fatalf("error message should contain resource: %v", derr.Error())
	}
	if derr.Unwrap() == nil {
		t.Fatalf("expected wrapped error")
	}

	// request error from underlying client
	c3 := client.NewHTTPClient(&fakeHTTP{fn: func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("boom")
	}}, client.WithBaseURL("http://example"))
	e3 := New[Item, Item](c3, "cards")
	_, err = e3.Get(context.Background(), "123")
	if err == nil {
		t.Fatalf("expected request error")
	}
	var re *client.RequestError
	if !errors.As(err, &re) {
		t.Fatalf("expected client.RequestError, got %T", err)
	}
}

func TestListSuccessNilAndWithQueryAndErrors(t *testing.T) {
	type Item struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	// nil query -> no qs
	c := client.NewHTTPClient(&fakeHTTP{fn: func(req *http.Request) (*http.Response, error) {
		if req.URL.String() != "http://example/cards" {
			t.Fatalf("unexpected url for nil query: %s", req.URL.String())
		}
		return client.NewMockResponse(200, `[{"id":"1","name":"a"},{"id":"2","name":"b"}]`), nil
	}}, client.WithBaseURL("http://example"))
	e := New[Item, Item](c, "cards")
	items, err := e.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}

	// with query -> appended qs
	q := query.New().Contains("name", "bob")
	c2 := client.NewHTTPClient(&fakeHTTP{fn: func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.String(), "http://example/cards") || !strings.Contains(req.URL.String(), "name=bob") {
			t.Fatalf("unexpected url for query: %s", req.URL.String())
		}
		return client.NewMockResponse(200, `[{"id":"9","name":"bob"}]`), nil
	}}, client.WithBaseURL("http://example"))
	e2 := New[Item, Item](c2, "cards")
	items2, err := e2.List(context.Background(), q)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items2) != 1 || items2[0].ID != "9" {
		t.Fatalf("unexpected items: %#v", items2)
	}

	// decode error
	c3 := client.NewHTTPClient(&fakeHTTP{fn: func(req *http.Request) (*http.Response, error) {
		return client.NewMockResponse(200, `not-json`), nil
	}}, client.WithBaseURL("http://example"))
	e3 := New[Item, Item](c3, "cards")
	_, err = e3.List(context.Background(), nil)
	if err == nil {
		t.Fatalf("expected decode error")
	}
	var derr *DecodeError
	if !errors.As(err, &derr) {
		t.Fatalf("expected DecodeError for list, got %T", err)
	}

	// request error
	c4 := client.NewHTTPClient(&fakeHTTP{fn: func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("boom")
	}}, client.WithBaseURL("http://example"))
	e4 := New[Item, Item](c4, "cards")
	_, err = e4.List(context.Background(), nil)
	if err == nil {
		t.Fatalf("expected request error")
	}
}
