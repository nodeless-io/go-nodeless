package nodeless

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Store represents a Store.
type Store struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func (s *Store) String() string {
	return fmt.Sprintf("id=%s name=%s", s.ID, s.Name)
}

// GetStores gets a list of Stores.
func (c *Client) GetStores(ctx context.Context) ([]Store, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/store")
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data []Store `json:"data"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// GetStore gets a Store.
func (c *Client) GetStore(ctx context.Context, id string) (*Store, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/store", id)
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data Store `json:"data"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
