package soren

import (
	"context"
	"fmt"
)

type SystemInfo struct {
	Version  string `json:"version"`
	Hostname string `json:"hostname"`
	Banner   string `json:"banner"`
}

func (c *Client) GetInfo(ctx context.Context) (*SystemInfo, error) {
	var payload struct{}

	var info SystemInfo

	apiPath := "/info"

	if err := c.Call(ctx, apiPath, payload, &info); err != nil {
		return nil, fmt.Errorf("failed to retrieve system info: %w", err)
	}

	return &info, nil
}
