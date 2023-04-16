package nodeless

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// New creates a new Nodeless client.
func New(config Config) (*Client, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &Client{
		config: config,
	}, nil
}

// Client implements the Nodeless API.
// Use 'New' to instantiate.
type Client struct {
	config Config
}

func (c *Client) authHeader() string {
	return fmt.Sprintf("Bearer %s", c.config.APIKey)
}

func (c *Client) do(ctx context.Context, method, endpoint string, body, resp any) error {
	var payload io.Reader
	if body != nil {
		jsonb, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("json marshal: %w", err)
		}
		payload = bytes.NewBuffer(jsonb)
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, payload)
	if err != nil {
		return fmt.Errorf("http NewRequest: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", c.authHeader())
	req.Header.Set("Content-Type", "application/json")

	result, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	if result.StatusCode > 300 {
		var resp struct {
			Message string `json:"message"`
		}
		if err := json.NewDecoder(result.Body).Decode(&resp); err != nil {
			return fmt.Errorf("json decode: %w", err)
		}
		return parseError(resp.Message)
	}

	// NOTE: delete requests return 200 with no body
	if method == http.MethodDelete {
		return nil
	}

	if err := json.NewDecoder(result.Body).Decode(resp); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	return nil
}
