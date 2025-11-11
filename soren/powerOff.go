package soren

import (
	"context"
)

type PowerOffPayload struct {
	Op   string   `json:"op"`
	Path []string `json:"path"`
}

func (c *Client) RePowerOff(ctx context.Context, path []string) (map[string]any, error) {
	payload := PowerOffPayload{
		Op:   "poweroff",
		Path: path,
	}

	var result map[string]any

	err := Call(c, ctx, "/reboot", payload, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
