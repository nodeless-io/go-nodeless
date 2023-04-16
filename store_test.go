//go:build integration
// +build integration

package nodeless

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStores(t *testing.T) {
	var tests = []struct {
		name string
		// Only expecting a single store from the testnet account
		store Store
		err   error
	}{
		{"basic", Store{
			ID:   integrationSecrets.StoreID,
			Name: integrationSecrets.StoreName,
		}, nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stores, err := client.GetStores(ctx)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, stores)

			assert.Len(t, stores, 1)
			assert.Equal(t, tt.store.ID, stores[0].ID)
			assert.Equal(t, tt.store.Name, stores[0].Name)
			assert.NotEmpty(t, stores[0].CreatedAt)
		})
	}
}

func TestGetStore(t *testing.T) {
	var tests = []struct {
		name  string
		store Store
		err   error
	}{
		{"basic", Store{
			ID:   integrationSecrets.StoreID,
			Name: integrationSecrets.StoreName,
		}, nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			store, err := client.GetStore(ctx, tt.store.ID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, store)

			assert.Equal(t, tt.store.ID, store.ID)
			assert.Equal(t, tt.store.Name, store.Name)
			assert.NotEmpty(t, store.CreatedAt)
		})
	}
}
