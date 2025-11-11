package soren

import (
	"context"
)

type RebootPayload struct {
	Op   string   `json:"op"`
	Path []string `json:"path"`
}

func (c *Client) Reboot(ctx context.Context, path []string) (map[string]any, error) {
	payload := RebootPayload{
		Op:   "reboot",
		Path: path,
	}

	var result map[string]any

	err := Call(c, ctx, "/reboot", payload, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
