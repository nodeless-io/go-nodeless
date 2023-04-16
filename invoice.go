package nodeless

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// CreateInvoiceRequest is a request to create an Invoice.
type CreateInvoiceRequest struct {
	StoreID     string            `json:"-"`
	Amount      float64           `json:"amount"`
	Currency    string            `json:"currency"`
	BuyerEmail  string            `json:"buyerEmail"`
	RedirectUrl string            `json:"redirectUrl"`
	Metadata    map[string]string `json:"metadata"`
}

// Invoice represents a Store Invoice.
type Invoice struct {
	ID               string         `json:"id"`
	CheckoutLink     string         `json:"checkoutLink"`
	SatsAmount       float64        `json:"satsAmount"`
	Status           InvoiceStatus  `json:"status"`
	BuyerEmail       string         `json:"buyerEmail"`
	RedirectURL      string         `json:"redirectUrl"`
	Metadata         map[string]any `json:"metadata"`
	CreatedAt        time.Time      `json:"createdAt"`
	PaidAt           *time.Time     `json:"paidAt"`
	OnchainAddress   string         `json:"onchainAddress"`
	LightningInvoice string         `json:"lightningInvoice"`
	Store            Store          `json:"store"`
	QRCodes          QRCodes        `json:"qrCodes"`
}

func (i *Invoice) String() string {
	return fmt.Sprintf("id=%s amount=%v status=%s", i.ID, i.SatsAmount, i.Status)
}

// QRCodes represents Invoice QRCodes.
type QRCodes struct {
	Unified   string `json:"unified"`
	Onchain   string `json:"onchain"`
	Lightning string `json:"lightning"`
}

// InvoiceStatus represents an Invoice's Status.
type InvoiceStatus string

const (
	InvoiceStatusUnknown InvoiceStatus = ""
	InvoiceStatusNew     InvoiceStatus = "new"
	InvoiceStatusPaid    InvoiceStatus = "paid"
	InvoiceStatusExpired InvoiceStatus = "expired"
)

// CreateStoreInvoice creates a new Invoice.
func (c *Client) CreateStoreInvoice(ctx context.Context, r CreateInvoiceRequest) (*Invoice, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/store", r.StoreID, "invoice")
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data Invoice `json:"data"`
	}

	if err := c.do(ctx, http.MethodPost, endpoint, r, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// GetStoreInvoice fetchs an existing Invoice.
func (c *Client) GetStoreInvoice(ctx context.Context, storeID, invoiceID string) (*Invoice, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/store", storeID, "invoice", invoiceID)
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data Invoice `json:"data"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

// GetStoreInvoiceStatus fetchs an existing Invoice Status.
func (c *Client) GetStoreInvoiceStatus(ctx context.Context, id, invoiceID string) (InvoiceStatus, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/store", id, "invoice", invoiceID, "status")
	if err != nil {
		return InvoiceStatusUnknown, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Status InvoiceStatus `json:"status"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return InvoiceStatusUnknown, err
	}

	return resp.Status, nil
}
