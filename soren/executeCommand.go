package soren

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type APIResponse struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   string          `json:"error,omitempty"`
}

func (c *Client) executeCommand(ctx context.Context, apiPath string, payload any) (*APIResponse, error) {
	fullURL := c.baseURL + apiPath

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON payload: %w", err)
	}
	jsonString := string(jsonBytes)

	formData := url.Values{
		"data": {jsonString},
		"key":  {c.apiKey},
	}

	body := formData.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to form request for %s: %w", fullURL, err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to complete request to %s: %w", fullURL, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body for %s: %w", fullURL, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request to %s failed with status: %s. Response body: %s", fullURL, resp.Status, string(respBody))
	}

	var apiResponse APIResponse
	if err := json.Unmarshal(respBody, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode API response wrapper for %s: %w. Raw response: %s", fullURL, err, string(respBody))
	}

	if !apiResponse.Success {
		return nil, errors.New("API command failed: " + apiResponse.Error)
	}

	return &apiResponse, nil
}
