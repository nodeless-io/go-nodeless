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

func TestGetStoreWebhooks(t *testing.T) {
	var tests = []struct {
		name    string
		storeID string
		err     error
	}{
		{"basic", integrationSecrets.StoreID, nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			webhooks, err := client.GetStoreWebhooks(ctx, tt.storeID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, webhooks)
		})
	}
}

func TestCreateStoreWebhook(t *testing.T) {
	var tests = []struct {
		name string
		req  CreateStoreWebhookRequest
		err  error
	}{
		{
			name: "basic",
			req: CreateStoreWebhookRequest{
				StoreID: integrationSecrets.StoreID,
				Type:    WebhookTypeStore,
				URL:     "https://example.com",
				Events:  []WebhookEvent{WebhookEventNew},
				Secret:  "testSecret",
				Status:  WebhookStatusActive,
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
			webhook, err := client.CreateStoreWebhook(ctx, tt.req)
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

func TestGetStoreWebhook(t *testing.T) {
	var tests = []struct {
		name      string
		storeID   string
		webhookID string
		err       error
	}{
		{"basic", integrationSecrets.StoreID, integrationSecrets.StoreWebhookID, nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			webhook, err := client.GetStoreWebhook(ctx, tt.storeID, tt.webhookID)
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

func TestUpdateStoreWebhook(t *testing.T) {
	var tests = []struct {
		name string
		req  UpdateStoreWebhookRequest
		err  error
	}{
		{
			name: "basic",
			req: UpdateStoreWebhookRequest{
				StoreID:   integrationSecrets.StoreID,
				WebhookID: integrationSecrets.StoreWebhookID,
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
			webhook, err := client.UpdateStoreWebhook(ctx, tt.req)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, webhook)
		})
	}
}

func TestDeleteStoreWebhook(t *testing.T) {
	var tests = []struct {
		name      string
		storeID   string
		webhookID string
		err       error
	}{
		{"not exists", integrationSecrets.StoreID, "foo", ErrWebhookNotFound},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			err := client.DeleteStoreWebhook(ctx, tt.storeID, tt.webhookID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
