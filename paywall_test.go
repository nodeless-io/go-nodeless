//go:build integration
// +build integration

package nodeless

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPaywalls(t *testing.T) {
	var tests = []struct {
		name string
		err  error
	}{
		{"basic", nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			paywalls, err := client.GetPaywalls(ctx)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, paywalls)
		})
	}
}

func TestCreatePaywall(t *testing.T) {
	var tests = []struct {
		name string
		req  CreatePaywallRequest
		err  error
	}{
		{
			name: "basic",
			req: CreatePaywallRequest{
				Name:  "test1",
				Type:  PaywallTypeRedirect,
				Price: 1000,
			},
			err: nil,
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
			paywall, err := client.CreatePaywall(ctx, tt.req)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, paywall)
			t.Logf("created Paywall:\n%#v\n", *paywall)
		})
	}
}

func TestGetPaywall(t *testing.T) {
	var tests = []struct {
		name      string
		paywallID string
		err       error
	}{
		{"basic", integrationSecrets.PaywallID, nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			paywall, err := client.GetPaywall(ctx, tt.paywallID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, paywall)

			assert.Equal(t, tt.paywallID, paywall.ID)
			assert.NotEmpty(t, paywall.CreatedAt)
		})
	}
}

func TestUpdatePaywall(t *testing.T) {
	var tests = []struct {
		name string
		req  UpdatePaywallRequest
		err  error
	}{
		{
			name: "basic",
			req: UpdatePaywallRequest{
				ID:    integrationSecrets.PaywallID,
				Name:  fmt.Sprintf("test1-%d", time.Now().Unix()),
				Type:  PaywallTypeRedirect,
				Price: 1001,
			},
			err: nil,
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
			paywall, err := client.UpdatePaywall(ctx, tt.req)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, paywall)
		})
	}
}

func TestDeletePaywall(t *testing.T) {
	var tests = []struct {
		name      string
		paywallID string
		err       error
	}{
		{"not exists", "not-exists", ErrPaywallNotFound},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			err := client.DeletePaywall(ctx, tt.paywallID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
