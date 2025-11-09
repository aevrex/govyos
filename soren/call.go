package soren

import (
	"context"
	"encoding/json"
	"fmt"
)

func Call[T any](c *Client, ctx context.Context, apiPath string, payload any, result *T) error {
    apiResp, err := c.executeCommand(ctx, apiPath, payload)
    if err != nil {
        return err 
    }

    if err := json.Unmarshal(apiResp.Data, result); err != nil {
        return fmt.Errorf("failed to unmarshal API data into target type: %w", err)
    }

    return nil
}