package nodeless

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// GetStoreWebooks gets a list of StoreWebhooks.
func (c *Client) GetStoreWebhooks(ctx context.Context, storeID string) ([]Webhook, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/store", storeID, "webhook")
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data []Webhook `json:"data"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// CreateStoreWebhookRequest is a request to create a Store Webhook.
type CreateStoreWebhookRequest struct {
	StoreID string         `json:"-"`
	Type    WebhookType    `json:"type"`
	URL     string         `json:"url"`
	Events  []WebhookEvent `json:"events"`
	Secret  string         `json:"secret"`
	Status  WebhookStatus  `json:"status"`
}

// CreateStoreWebhook creates a new Store Webhook.
func (c *Client) CreateStoreWebhook(ctx context.Context, r CreateStoreWebhookRequest) (*Webhook, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/store", r.StoreID, "webhook")
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data Webhook `json:"data"`
	}

	if err := c.do(ctx, http.MethodPost, endpoint, r, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// GetStoreWebhook gets a Store Webhook.
func (c *Client) GetStoreWebhook(ctx context.Context, storeID, id string) (*Webhook, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/store", storeID, "webhook", id)
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data Webhook `json:"data"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// DeleteStoreWebhook deletes a Store Webhook.
func (c *Client) DeleteStoreWebhook(ctx context.Context, storeID, id string) error {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/store", storeID, "webhook", id)
	if err != nil {
		return fmt.Errorf("url JoinPath: %w", err)
	}

	if err := c.do(ctx, http.MethodDelete, endpoint, nil, nil); err != nil {
		return err
	}

	return nil
}

// UpdateStoreWebhookRequest is a request to update a Store Webhook.
type UpdateStoreWebhookRequest struct {
	StoreID   string         `json:"-"`
	WebhookID string         `json:"-"`
	URL       string         `json:"url"`
	Events    []WebhookEvent `json:"events"`
	Status    WebhookStatus  `json:"status"`
}

// UpdateStoreWebhook updates a Store Webhook.
func (c *Client) UpdateStoreWebhook(ctx context.Context, r UpdateStoreWebhookRequest) (*Webhook, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/store", r.StoreID, "webhook", r.WebhookID)
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data Webhook `json:"data"`
	}

	if err := c.do(ctx, http.MethodPut, endpoint, r, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
