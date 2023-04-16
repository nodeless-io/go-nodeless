package nodeless

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// PaywallType represents a Paywall's Type.
type PaywallType string

const (
	PaywallTypeUnknown   PaywallType = ""
	PaywallTypeContent   PaywallType = "content"
	PaywallTypeDownload  PaywallType = "download"
	PaywallTypeRedirect  PaywallType = "redirect"
	PaywallTypeWPArticle PaywallType = "wp_article"
)

// Paywall represents a Paywall
type Paywall struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Type      PaywallType `json:"type"`
	Price     float64     `json:"price"`
	Settings  any         `json:"settings"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func (s *Paywall) String() string {
	return fmt.Sprintf("id=%s name=%s", s.ID, s.Name)
}

// GetPaywalls gets a list of Paywalls.
func (c *Client) GetPaywalls(ctx context.Context) ([]Paywall, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall")
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data []Paywall `json:"data"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// CreatePaywallRequest is a request to create a Store Paywall.
type CreatePaywallRequest struct {
	Name     string            `json:"name"`
	Type     PaywallType       `json:"type"`
	Price    float64           `json:"price"`
	Settings map[string]string `json:"settings"`
}

// CreatePaywall creates a new Store Paywall.
func (c *Client) CreatePaywall(ctx context.Context, r CreatePaywallRequest) (*Paywall, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall")
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data Paywall `json:"data"`
	}

	if err := c.do(ctx, http.MethodPost, endpoint, r, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// GetPaywall gets a Paywall.
func (c *Client) GetPaywall(ctx context.Context, id string) (*Paywall, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall", id)
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data Paywall `json:"data"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// UpdatePaywallRequest is a request to update a Paywall.
type UpdatePaywallRequest struct {
	ID       string            `json:"-"`
	Name     string            `json:"name"`
	Type     PaywallType       `json:"type"`
	Price    float64           `json:"price"`
	Settings map[string]string `json:"settings"`
}

// UpdatePaywall updates a Store Webhook.
func (c *Client) UpdatePaywall(ctx context.Context, r UpdatePaywallRequest) (*Paywall, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall", r.ID)
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data Paywall `json:"data"`
	}

	if err := c.do(ctx, http.MethodPut, endpoint, r, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// DeletePaywall deletes a Paywall.
func (c *Client) DeletePaywall(ctx context.Context, id string) error {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall", id)
	if err != nil {
		return fmt.Errorf("url JoinPath: %w", err)
	}

	if err := c.do(ctx, http.MethodDelete, endpoint, nil, nil); err != nil {
		return err
	}

	return nil
}
