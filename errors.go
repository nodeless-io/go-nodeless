package nodeless

import (
	"fmt"
	"strings"
)

var (
	ErrInvoiceNotFound        = fmt.Errorf("invoice not found")
	ErrDonationPageNotFound   = fmt.Errorf("donation page not found")
	ErrWebhookNotFound        = fmt.Errorf("webhook not found")
	ErrPaywallRequestNotFound = fmt.Errorf("paywall request not found")
	ErrPaywallNotFound        = fmt.Errorf("paywall not found")
	ErrNotFound               = fmt.Errorf("not found")
	// Config
	ErrMissingAPIKey = fmt.Errorf("missing api key")
)

// parseError returns strongly typed errors from api error messages
func parseError(msg string) error {
	msg = strings.ToLower(msg)

	// Resource not found
	if strings.HasPrefix(msg, "no query results") {
		if strings.Contains(msg, "storeinvoice") {
			return ErrInvoiceNotFound
		}
		if strings.Contains(msg, "donationpage") {
			return ErrDonationPageNotFound
		}
		if strings.Contains(msg, "webhook") {
			return ErrWebhookNotFound
		}
		if strings.Contains(msg, "paywallrequest") {
			return ErrPaywallRequestNotFound
		}
		if strings.Contains(msg, "paywall") {
			return ErrPaywallNotFound
		}
	}

	// Unknown error
	return fmt.Errorf(msg)
}
