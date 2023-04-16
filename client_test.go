package nodeless_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nodeless-io/go-nodeless"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		name   string
		config nodeless.Config
		err    error
	}{
		{"testnet", nodeless.Config{
			APIKey:     "some-api-key",
			UseTestnet: true,
		}, nil},
		{"not testnet", nodeless.Config{
			APIKey:     "some-api-key",
			UseTestnet: false,
		}, nil},
		{"missing api key", nodeless.Config{
			APIKey:     "",
			UseTestnet: true,
		}, nodeless.ErrMissingAPIKey},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := nodeless.New(tt.config)
			if err != nil {
				assert.Equal(t, tt.err, err)
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, client)
		})
	}
}
