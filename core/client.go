package core

import (
	"net/http"
	"time"
)

type Client struct {
	BaseURL   string
	APIKey    string
	StoreID   string
	UserAgent string
	HTTP      *http.Client
}

type Option func(*Client)

func WithBaseURL(v string) Option { return func(c *Client) { c.BaseURL = v } }
func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) { c.HTTP = h }
}
func WithUserAgent(v string) Option { return func(c *Client) { c.UserAgent = v } }

func New(apiKey, storeID string, opts ...Option) *Client {
	c := &Client{
		BaseURL:   "https://sellium.site/api/v1",
		APIKey:    apiKey,
		StoreID:   storeID,
		UserAgent: "sellium-go/0.1",
		HTTP:      &http.Client{Timeout: 30 * time.Second},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
