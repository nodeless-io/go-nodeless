//go:build integration
// +build integration

package nodeless

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// This file loads secrets used for integration tests

const integrationSecretsPath = "./integration.secrets"

var integrationSecrets struct {
	APIKey           string `yaml:"api_key"`
	StoreID          string `yaml:"store_id"`
	StoreName        string `yaml:"store_name"`
	ExpiredInvoiceID string `yaml:"expired_invoice_id"`
	PaidInvoiceID    string `yaml:"paid_invoice_id"`
	TransactionID    string `yaml:"transaction_id"`
	StoreWebhookID   string `yaml:"store_webhook_id"`
	PaywallID        string `yaml:"paywall_id"`
	PaywallWebhookID string `yaml:"paywall_webhook_id"`
	PaywallRequestID string `yaml:"paywall_request_id"`
}

func init() {
	f, err := os.Open(integrationSecretsPath)
	if err != nil {
		fmt.Printf("must create 'integration.secrets' file!")
		os.Exit(1)
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(&integrationSecrets); err != nil {
		fmt.Printf("failed to decode integrations.secrets file as yaml: %v", err)
		os.Exit(1)
	}
}
