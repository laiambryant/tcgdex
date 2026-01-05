package client

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	BaseURL   string
	HTTP      HTTPClient
	UserAgent string
	cache     *Cache
}

func NewHTTPClient(httpClient HTTPClient, opts ...Option) *Client {
	c := &Client{
		BaseURL:   "https://api.tcgdex.net/v2/en",
		HTTP:      httpClient,
		UserAgent: "tcgdex-go-sdk",
	}
	if c.HTTP == nil {
		c.HTTP = http.DefaultClient
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) Get(ctx context.Context, path string) ([]byte, error) {
	fullURL := c.BaseURL + path
	if c.cache != nil {
		if data, ok := c.cache.Get(fullURL); ok {
			return data, nil
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, &RequestError{Op: "create request", Err: err}
	}
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, &RequestError{Op: "do request", Err: err}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &RequestError{Op: "read body", Err: err}
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &HTTPError{
			Status: resp.StatusCode,
			URL:    fullURL,
			Body:   string(body),
			Cause:  errors.New("api error"),
		}
	}

	if c.cache != nil {
		c.cache.Set(fullURL, body)
	}

	return body, nil
}

func (c *Client) Download(ctx context.Context, urlStr string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, &RequestError{Op: "create request", Err: err}
	}
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, &RequestError{Op: "do request", Err: err}
	}

	if resp.StatusCode == http.StatusNotFound {
		resp.Body.Close()
		return nil, ErrNotFound
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, &HTTPError{
			Status: resp.StatusCode,
			URL:    urlStr,
			Body:   string(body),
			Cause:  errors.New("download error"),
		}
	}

	return resp.Body, nil
}
