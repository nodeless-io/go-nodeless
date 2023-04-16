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

func TestGetPaywallWebhooks(t *testing.T) {
	var tests = []struct {
		name      string
		PaywallID string
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
			webhooks, err := client.GetPaywallWebhooks(ctx, tt.PaywallID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, webhooks)
		})
	}
}

func TestCreatePaywallWebhook(t *testing.T) {
	var tests = []struct {
		name string
		req  CreatePaywallWebhookRequest
		err  error
	}{
		{
			name: "basic",
			req: CreatePaywallWebhookRequest{
				PaywallID: integrationSecrets.PaywallID,
				Type:      WebhookTypePaywall,
				URL:       "https://example.com",
				Events:    []WebhookEvent{WebhookEventNew},
				Secret:    "testSecret",
				Status:    WebhookStatusActive,
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
			webhook, err := client.CreatePaywallWebhook(ctx, tt.req)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, webhook)
			t.Logf("created webhook:\n%#v\n", *webhook)
		})
	}
}

func TestGetPaywallWebhook(t *testing.T) {
	var tests = []struct {
		name      string
		paywallID string
		webhookID string
		err       error
	}{
		{"basic", integrationSecrets.PaywallID, integrationSecrets.PaywallWebhookID, nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			webhook, err := client.GetPaywallWebhook(ctx, tt.paywallID, tt.webhookID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, webhook)

			assert.Equal(t, tt.webhookID, webhook.ID)
			assert.NotEmpty(t, webhook.CreatedAt)
		})
	}
}

func TestUpdatePaywallWebhook(t *testing.T) {
	var tests = []struct {
		name string
		req  UpdatePaywallWebhookRequest
		err  error
	}{
		{
			name: "basic",
			req: UpdatePaywallWebhookRequest{
				PaywallID: integrationSecrets.PaywallID,
				WebhookID: integrationSecrets.PaywallWebhookID,
				URL:       fmt.Sprintf("https://example.com/%v", time.Now().Unix()),
				Events:    []WebhookEvent{WebhookEventNew},
				Status:    WebhookStatusActive,
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
			webhook, err := client.UpdatePaywallWebhook(ctx, tt.req)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, webhook)
		})
	}
}

func TestDeletePaywallWebhook(t *testing.T) {
	var tests = []struct {
		name      string
		PaywallID string
		webhookID string
		err       error
	}{
		{"not exists", integrationSecrets.PaywallID, "foo", ErrWebhookNotFound},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			err := client.DeletePaywallWebhook(ctx, tt.PaywallID, tt.webhookID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
