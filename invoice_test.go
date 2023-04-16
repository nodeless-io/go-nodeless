//go:build integration
// +build integration

package nodeless

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInvoice(t *testing.T) {
	var tests = []struct {
		name     string
		storeID  string
		amount   float64
		currency string
		status   InvoiceStatus
		err      error
	}{
		{"new", integrationSecrets.StoreID, 1000, "SATS", InvoiceStatusNew, nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			resp, err := client.CreateStoreInvoice(ctx, CreateInvoiceRequest{
				StoreID:  tt.storeID,
				Amount:   tt.amount,
				Currency: tt.currency,
			})
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)

			assert.Equal(t, tt.status, resp.Status)
		})
	}
}

func TestGetInvoice(t *testing.T) {
	var tests = []struct {
		name      string
		storeID   string
		invoiceID string
		status    InvoiceStatus
		err       error
	}{
		{"paid", integrationSecrets.StoreID, integrationSecrets.PaidInvoiceID, InvoiceStatusPaid, nil},
		{"expired", integrationSecrets.StoreID, integrationSecrets.ExpiredInvoiceID, InvoiceStatusExpired, nil},
		{"missing invoice", integrationSecrets.StoreID, "does-not-exist", InvoiceStatusExpired, ErrInvoiceNotFound},
		// NOTE: should probably return an error, but doesn't
		{"missing store", "does-not-exist", integrationSecrets.ExpiredInvoiceID, InvoiceStatusExpired, nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			resp, err := client.GetStoreInvoice(ctx, tt.storeID, tt.invoiceID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)

			assert.Equal(t, tt.status, resp.Status)
			if tt.status == InvoiceStatusPaid {
				assert.NotNil(t, resp.PaidAt)
			}
		})
	}
}

func TestGetInvoiceStatus(t *testing.T) {
	var tests = []struct {
		name      string
		storeID   string
		invoiceID string
		status    InvoiceStatus
		err       error
	}{
		{"paid", integrationSecrets.StoreID, integrationSecrets.PaidInvoiceID, InvoiceStatusPaid, nil},
		{"expired", integrationSecrets.StoreID, integrationSecrets.ExpiredInvoiceID, InvoiceStatusExpired, nil},
		{"missing invoice", integrationSecrets.StoreID, "does-not-exist", InvoiceStatusExpired, ErrInvoiceNotFound},
		// NOTE: should probably return an error, but doesn't
		{"missing store", "does-not-exist", integrationSecrets.ExpiredInvoiceID, InvoiceStatusExpired, nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			status, err := client.GetStoreInvoiceStatus(ctx, tt.storeID, tt.invoiceID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				assert.Equal(t, InvoiceStatusUnknown, status)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, status)

			assert.Equal(t, tt.status, status)
		})
	}
}
