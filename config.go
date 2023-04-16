package nodeless

const (
	testnetAPIBase = "https://testnet.nodeless.io"
	prodAPIBase    = "https://nodeless.io"
)

// Config configures the Nodeless Client
type Config struct {
	APIKey     string
	UseTestnet bool
}

func (c *Config) Validate() error {
	if c.APIKey == "" {
		return ErrMissingAPIKey
	}

	return nil
}

func (c *Config) apiBase() string {
	if c.UseTestnet {
		return testnetAPIBase
	}

	return prodAPIBase
}
