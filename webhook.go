package nodeless

import (
	"fmt"
	"time"
)

// Webhook represents a Webhook.
type Webhook struct {
	ID             string         `json:"id"`
	URL            string         `json:"url"`
	Secret         string         `json:"secret"`
	Status         WebhookStatus  `json:"status"`
	Events         []WebhookEvent `json:"events"`
	LastDeliveryAt *time.Time     `json:"lastDeliveryAt"`
	CreatedAt      time.Time      `json:"createdAt"`
}

func (s *Webhook) String() string {
	return fmt.Sprintf("id=%s url=%s status=%s", s.ID, s.URL, s.Status)
}

// WebhookType represents a Webhook's Type.
type WebhookType string

const (
	WebhookTypeUnknown      WebhookType = ""
	WebhookTypeStore        WebhookType = "store"
	WebhookTypeDonationPage WebhookType = "donation_page"
	WebhookTypePaywall      WebhookType = "paywall"
	WebhookTypeInbox        WebhookType = "inbox"
)

// WebhookEvent represents a Webhook Event.
type WebhookEvent string

const (
	WebhookEventUnknown             WebhookEvent = ""
	WebhookEventNew                 WebhookEvent = "new"
	WebhookEventPendingConfirmation WebhookEvent = "pending_confirmation"
	WebhookEventPaid                WebhookEvent = "paid"
	WebhookEventOverpaid            WebhookEvent = "overpaid"
	WebhookEventUnderpaid           WebhookEvent = "underpaid"
	WebhookEventInflight            WebhookEvent = "inflight"
	WebhookEventExpired             WebhookEvent = "expired"
	WebhookEventCancelled           WebhookEvent = "cancelled"
)

// WebhookStatus represents a Webhook's Status.
type WebhookStatus string

const (
	WebhookStatusUnknown  WebhookStatus = ""
	WebhookStatusActive   WebhookStatus = "active"
	WebhookStatusInactive WebhookStatus = "inactive"
)
