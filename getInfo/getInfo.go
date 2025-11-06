package soren

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/aevrex/soren/client"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type SystemInfo struct {
	Version  string `json:"version"`
	Hostname string `json:"hostname"`
	Banner   string `json:"banner"`
}

func (c *Client) GetInfo(ctx context.Context) (*SystemInfo, error) {
	infoURL := c.baseURL + "/info"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, infoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, responseBody)
	}

	var rawResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&rawResponse); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %w", err)
	}

	if !rawResponse.Success {
		return nil, fmt.Errorf("VyOS command failed: %s", rawResponse.Error)
	}

	dataBytes, err := json.Marshal(rawResponse.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data field: %w", err)
	}

	var info SystemInfo
	if err := json.Unmarshal(dataBytes, &info); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SystemInfo: %w", err)
	}

	return &info, nil
}