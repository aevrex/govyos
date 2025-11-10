package soren

import (
	"fmt"
	"net/http"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewClient(baseURL string, apiKey string, opts ...Option) (*Client, error) {
	c := &Client{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.baseURL == "" || c.apiKey == "" {
		return nil, fmt.Errorf("baseURL and token must be provided")
	}

	return c, nil
}
