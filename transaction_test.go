//go:build integration
// +build integration

package nodeless

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {
	var tests = []struct {
		name string
		// Only expecting a single store from the testnet account
		transaction Transaction
		err         error
	}{
		{"basic", Transaction{ID: integrationSecrets.TransactionID}, nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			transactions, err := client.GetTransactions(ctx)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, transactions)

			var expectedTx *Transaction
			for _, tx := range transactions {
				if strings.EqualFold(tx.ID, tt.transaction.ID) {
					expectedTx = &tx
					break
				}
			}

			if expectedTx == nil {
				t.Errorf("did not get expected transaction %s", tt.transaction.ID)
				t.Logf("got: %#v\n", transactions)
				return
			}

		})
	}
}

func TestGetTransaction(t *testing.T) {
	var tests = []struct {
		name        string
		transaction Transaction
		err         error
	}{
		{"basic", Transaction{ID: integrationSecrets.TransactionID}, nil},
	}

	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			transaction, err := client.GetTransaction(ctx, tt.transaction.ID)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, transaction)
			assert.Equal(t, tt.transaction.ID, transaction.ID)
			assert.NotEmpty(t, transaction.CreatedAt)
		})
	}
}
