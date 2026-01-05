package client

import "time"

type Option func(*Client)

func WithUserAgent(ua string) Option {
	return func(c *Client) {
		c.UserAgent = ua
	}
}

func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.BaseURL = url
	}
}

func WithCache(ttl time.Duration) Option {
	return func(c *Client) {
		c.cache = NewCache(ttl)
	}
}

func WithHTTPClient(httpClient HTTPClient) Option {
	return func(c *Client) {
		c.HTTP = httpClient
	}
}
