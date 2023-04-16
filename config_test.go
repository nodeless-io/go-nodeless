package nodeless_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nodeless-io/go-nodeless"
)

func TestValidateConfig(t *testing.T) {
	var tests = []struct {
		name       string
		apiKey     string
		useTestnet bool
		err        error
	}{
		{"testnet", "some-api-key", true, nil},
		{"not testnet", "some-api-key", false, nil},
		{"missing api key", "", true, nodeless.ErrMissingAPIKey},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := nodeless.Config{
				APIKey:     tt.apiKey,
				UseTestnet: tt.useTestnet,
			}

			err := config.Validate()

			assert.Equal(t, tt.err, err)
		})
	}
}
