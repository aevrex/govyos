package client

import (
	"errors"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewClient(baseURL string, apiKey string) (*Client, error) {
	if baseURL == "" {
		return nil, errors.New("URL cannot be empty")
	}

	if apiKey == "" {
		return nil, errors.New("API key cannot be empty")
	}

	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	newClient := &Client{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: httpClient,
	}

	return newClient, nil
}