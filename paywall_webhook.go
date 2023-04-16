package nodeless

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// GetPaywallWebooks gets a list of PaywallWebhooks.
func (c *Client) GetPaywallWebhooks(ctx context.Context, paywallID string) ([]Webhook, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall", paywallID, "webhook")
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

// CreatePaywallWebhookRequest is a request to create a Paywall Webhook.
type CreatePaywallWebhookRequest struct {
	PaywallID string         `json:"-"`
	Type      WebhookType    `json:"type"`
	URL       string         `json:"url"`
	Events    []WebhookEvent `json:"events"`
	Secret    string         `json:"secret"`
	Status    WebhookStatus  `json:"status"`
}

// CreatePaywallWebhook creates a new Paywall Webhook.
func (c *Client) CreatePaywallWebhook(ctx context.Context, r CreatePaywallWebhookRequest) (*Webhook, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall", r.PaywallID, "webhook")
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

// GetPaywallWebhook gets a Paywall Webhook.
func (c *Client) GetPaywallWebhook(ctx context.Context, paywallID, id string) (*Webhook, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall", paywallID, "webhook", id)
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

// DeletePaywallWebhook deletes a Paywall Webhook.
func (c *Client) DeletePaywallWebhook(ctx context.Context, paywallID, id string) error {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall", paywallID, "webhook", id)
	if err != nil {
		return fmt.Errorf("url JoinPath: %w", err)
	}

	if err := c.do(ctx, http.MethodDelete, endpoint, nil, nil); err != nil {
		return err
	}

	return nil
}

// UpdatePaywallWebhookRequest is a request to update a Paywall Webhook.
type UpdatePaywallWebhookRequest struct {
	PaywallID string         `json:"-"`
	WebhookID string         `json:"-"`
	URL       string         `json:"url"`
	Events    []WebhookEvent `json:"events"`
	Status    WebhookStatus  `json:"status"`
}

// UpdatePaywallWebhook updates a Paywall Webhook.
func (c *Client) UpdatePaywallWebhook(ctx context.Context, r UpdatePaywallWebhookRequest) (*Webhook, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall", r.PaywallID, "webhook", r.WebhookID)
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
