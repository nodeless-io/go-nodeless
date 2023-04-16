package nodeless

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// APIStatus represents the Nodeless API Status.
type APIStatus struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Node   string `json:"node"`
}

func (s *APIStatus) String() string {
	return fmt.Sprintf("code=%d status=%s node=%s", s.Code, s.Status, s.Node)
}

// GetAPIStatus gets the Nodeless API Status.
func (c *Client) GetAPIStatus(ctx context.Context) (*APIStatus, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/status")
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data APIStatus `json:"data"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
