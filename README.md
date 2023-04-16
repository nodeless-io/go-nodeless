[![ci](https://github.com/nodeless-io/go-nodeless/actions/workflows/main_push.yml/badge.svg)](https://github.com/nodeless-io/go-nodeless/actions/workflows/main_push.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/nodeless-io/go-nodeless.svg)](https://pkg.go.dev/github.com/nodeless-io/go-nodeless)

# go-nodeless

The [Nodeless.io](https://nodeless.io) Go SDK

## Quickstart

```go
package main

import (
	"context"

	"github.com/nodeless-io/go-nodeless"
)

func main() {
	// Create a client
	client := nodeless.New(nodeless.Config{
		APIKey: "my-api-key",
	})

	ctx := context.Background()

	// Create an invoice
	invoice, _ := client.CreateStoreInvoice(ctx, nodeless.CreateInvoiceRequest{
		StoreID:  "my-store-id",
		Amount:   1000,
		Currency: "SATS",
	})
}
```

For complete documentation, see the [Godoc](https://pkg.go.dev/github.com/nodeless-io/go-nodeless).

## Status

Impementation status of full [API](https://nodeless.io/api-docs#) support:

### Paywall Requests
- [x] [Create a Paywall Request](https://nodeless.io/api-docs#paywall-requests-POSTapi-v1-paywall--id--request)
- [x] [Get a Paywall Request](https://nodeless.io/api-docs#paywall-requests-GETapi-v1-paywall--id--request--requestId-)
- [x] [Get a Paywall Request Status](https://nodeless.io/api-docs#paywall-requests-GETapi-v1-paywall--id--request--requestId--status)

### Paywall Webhooks
- [x] [Get Paywall Webhooks](https://nodeless.io/api-docs#paywall-webhooks-GETapi-v1-paywall--id--webhook)
- [x] [Create Paywall Webhook](https://nodeless.io/api-docs#paywall-webhooks-POSTapi-v1-paywall--id--webhook)
- [x] [Get Paywall Webhook](https://nodeless.io/api-docs#paywall-webhooks-GETapi-v1-paywall--id--webhook--webhookId-)
- [x] [Delete Paywall Webhook](https://nodeless.io/api-docs#paywall-webhooks-DELETEapi-v1-paywall--id--webhook--webhookId-)
- [x] [Update Paywall Webhook](https://nodeless.io/api-docs#paywall-webhooks-PUTapi-v1-paywall--id--webhook--webhookId-)

### Paywalls
- [x] [Get Paywalls](https://nodeless.io/api-docs#paywalls-GETapi-v1-paywall)
- [x] [Create Paywall](https://nodeless.io/api-docs#paywalls-POSTapi-v1-paywall)
- [x] [Get Paywall](https://nodeless.io/api-docs#paywalls-GETapi-v1-paywall--id-)
- [x] [Update Paywall](https://nodeless.io/api-docs#paywalls-PUTapi-v1-paywall--id-)
- [x] [Delete Paywall](https://nodeless.io/api-docs#paywalls-DELETEapi-v1-paywall--id-)

### Server Info
- [x] [Get API Status](https://nodeless.io/api-docs#server-info-GETapi-v1-status)

### Store Invoices
- [x] [Create Store Invoice](https://nodeless.io/api-docs#store-invoices-POSTapi-v1-store--id--invoice)
- [x] [Get Store Invoice](https://nodeless.io/api-docs#store-invoices-GETapi-v1-store--id--invoice--invoiceId-)
- [x] [Get Store Invoice Status](https://nodeless.io/api-docs#store-invoices-GETapi-v1-store--id--invoice--invoiceId--status)

### Store Webhooks
- [x] [Get Store Webhooks](https://nodeless.io/api-docs#store-webhooks-GETapi-v1-store--id--webhook)
- [x] [Create Store Webhook](https://nodeless.io/api-docs#store-webhooks-POSTapi-v1-store--id--webhook)
- [x] [Get Store Webhook](https://nodeless.io/api-docs#store-webhooks-GETapi-v1-store--id--webhook--webhookId-)
- [x] [Delete Store Webhook](https://nodeless.io/api-docs#store-webhooks-DELETEapi-v1-store--id--webhook--webhookId-)
- [x] [Update Store Webhook](https://nodeless.io/api-docs#store-webhooks-PUTapi-v1-store--id--webhook--webhookId-)

### Stores
- [x] [Get Stores](https://nodeless.io/api-docs#stores-GETapi-v1-store)
- [x] [Get Store](https://nodeless.io/api-docs#stores-GETapi-v1-store--id-)

### Transactions
- [x] [Get All Transactions](https://nodeless.io/api-docs#transactions-GETapi-v1-transaction)
- [x] [Get Transaction](https://nodeless.io/api-docs#transactions-GETapi-v1-transaction--id-)

## Integration Tests

go-nodeless is equiped with a full integration test suite that may be run to
test the SDK's integration with the Nodeless API.

1. Copy `integration.secrets.example` to `integration.secrets` and fill in the
fields with testnet values.
2. Run `make test-integration`
