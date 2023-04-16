package nodeless

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// PaywallRequest represents a Paywall Request
type PaywallRequest struct {
	ID               string               `json:"id"`
	SatsAmount       float64              `json:"satsAmount"`
	Status           PaywallRequestStatus `json:"status"`
	OnchainAddress   string               `json:"onchainAddress"`
	LightningInvoice string               `json:"lightningInvoice"`
	Paywall          Paywall              `json:"paywall"`
	QRCodes          QRCodes              `json:"qrCodes"`
	Metadata         map[string]string    `json:"metadata"`
	CreatedAt        time.Time            `json:"createdAt"`
	PaidAt           *time.Time           `json:"paidAt"`
}

func (r *PaywallRequest) String() string {
	return fmt.Sprintf("id=%s amount=%v status=%v", r.ID, r.SatsAmount, r.Status)
}

// CreatePaywallRequest creates a new Paywall Request.
func (c *Client) CreatePaywallRequest(ctx context.Context, paywallID string) (*PaywallRequest, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall", paywallID, "request")
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data PaywallRequest `json:"data"`
	}

	if err := c.do(ctx, http.MethodPost, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// GetPaywallRequest gets a Paywall Request.
func (c *Client) GetPaywallRequest(ctx context.Context, paywallID, id string) (*PaywallRequest, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall", paywallID, "request", id)
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data PaywallRequest `json:"data"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// TODO: Is this the same as InvoiceStatus?
// PaywallRequestStatus represents an PaywallRequest's Status.
type PaywallRequestStatus string

const (
	PaywallRequestStatusUnknown PaywallRequestStatus = ""
	PaywallRequestStatusNew     PaywallRequestStatus = "new"
	PaywallRequestStatusPaid    PaywallRequestStatus = "paid"
	PaywallRequestStatusExpired PaywallRequestStatus = "expired"
)

// GetPaywallRequestStatus gets the status of a Paywall Request.
func (c *Client) GetPaywallRequestStatus(ctx context.Context, paywallID, id string) (PaywallRequestStatus, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/paywall", paywallID, "request", id, "status")
	if err != nil {
		return PaywallRequestStatusUnknown, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Status PaywallRequestStatus `json:"status"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return PaywallRequestStatusUnknown, err
	}

	return resp.Status, nil
}
