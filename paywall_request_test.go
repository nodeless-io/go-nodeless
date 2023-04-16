//go:build integration
// +build integration

package nodeless

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePaywallRequest(t *testing.T) {
	var tests = []struct {
		name      string
		paywallID string
		err       error
	}{
		{
			name:      "basic",
			paywallID: integrationSecrets.PaywallID,
			err:       nil,
		},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			paywallRequest, err := client.CreatePaywallRequest(ctx, tt.paywallID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, paywallRequest)
			t.Logf("created PaywallRequest:\n%#v\n", *paywallRequest)
		})
	}
}

func TestGetPaywallRequest(t *testing.T) {
	var tests = []struct {
		name             string
		paywallID        string
		paywallRequestID string
		err              error
	}{
		{"basic", integrationSecrets.PaywallID, integrationSecrets.PaywallRequestID, nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			paywallRequest, err := client.GetPaywallRequest(ctx, tt.paywallID, tt.paywallRequestID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, paywallRequest)

			assert.Equal(t, tt.paywallRequestID, paywallRequest.ID)
			assert.NotEmpty(t, paywallRequest.CreatedAt)
		})
	}
}

func TestGetPaywallRequestStatus(t *testing.T) {
	var tests = []struct {
		name             string
		paywallID        string
		paywallRequestID string
		status           PaywallRequestStatus
		err              error
	}{
		{"basic", integrationSecrets.PaywallID, integrationSecrets.PaywallRequestID, PaywallRequestStatusExpired, nil},
		{"not exists", integrationSecrets.PaywallID, "not-exists", PaywallRequestStatusUnknown, ErrPaywallRequestNotFound},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			status, err := client.GetPaywallRequestStatus(ctx, tt.paywallID, tt.paywallRequestID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				assert.Equal(t, PaywallRequestStatusUnknown, status)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.status, status)
		})
	}
}
