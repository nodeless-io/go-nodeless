//go:build integration
// +build integration

package nodeless

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPIStatus(t *testing.T) {
	client, err := New(Config{
		APIKey:     integrationSecrets.APIKey,
		UseTestnet: true,
	})
	assert.NoError(t, err)

	ctx := context.Background()
	status, err := client.GetAPIStatus(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, status)
	t.Logf("APIStatus: %s", status)
}
