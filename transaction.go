package nodeless

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// TransactionStatus represents an Transaction's Status.
type TransactionStatus string

const (
	TransactionStatusUnknown TransactionStatus = ""
	TransactionStatusSettled TransactionStatus = "settled"
)

// TransactableType represents an Transaction's Transactable type.
type TransactableType string

const (
	TransactableTypeUnknown  TransactableType = ""
	TransactableTypeDonation TransactableType = "Donation"
)

// Transaction represents a Transaction.
type Transaction struct {
	ID               string            `json:"id"`
	Amount           float64           `json:"amount"`
	Type             string            `json:"type"`
	Status           TransactionStatus `json:"status"`
	TransactableType TransactableType  `json:"transactable_type"`
	Transactable     Transactable      `json:"transactable"`
	IsFee            bool              `json:"is_fee"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
}

func (t *Transaction) String() string {
	return fmt.Sprintf("id=%s amount=%v status=%v", t.ID, t.Amount, t.Status)
}

// Transaction represents a Transactable.
type Transactable struct {
	ID           int       `json:"id"`
	UUID         string    `json:"uuid"`
	DonationPage int       `json:"donation_page_id"`
	Amount       float64   `json:"amount"`
	AmountPaid   float64   `json:"amount_paid"`
	Name         string    `json:"name"`
	Message      string    `json:"message"`
	Status       string    `json:"status"`
	Type         string    `json:"type"`
	Metadata     any       `json:"metadata"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	// TODO: This should be *time.Time if/when the API returns a RFC-3339.
	PaidAt time.Time `json:"updated_at"`
}

func (t *Transactable) String() string {
	return fmt.Sprintf("id=%d uuid=%s name=%s amount=%v", t.ID, t.UUID, t.Name, t.Amount)
}

// GetTransactions gets a list of Transactions.
func (c *Client) GetTransactions(ctx context.Context) ([]Transaction, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/transaction")
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data []Transaction `json:"data"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// GetTransaction gets a Transaction.
func (c *Client) GetTransaction(ctx context.Context, id string) (*Transaction, error) {
	endpoint, err := url.JoinPath(c.config.apiBase(), "api/v1/transaction", id)
	if err != nil {
		return nil, fmt.Errorf("url JoinPath: %w", err)
	}

	var resp struct {
		Data Transaction `json:"data"`
	}

	if err := c.do(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
