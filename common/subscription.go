package common

import "time"

// External reference: https://discord.com/developers/docs/resources/subscription#subscription-object
type Subscription struct {
	Id                 string     `json:"id"`                   // ID of the subscription
	UserId             string     `json:"user_id"`              // ID of the user who is subscribed
	SkuIds             []string   `json:"sku_ids"`              // List of SKUs subscribed to
	EntitlementIds     []string   `json:"entitlement_ids"`      // List of entitlements granted for this subscription
	RenewalSkuIds      *[]string  `json:"renewal_sku_ids"`      // List of SKUs that this user will be subscribed to at renewal
	CurrentPeriodStart time.Time  `json:"current_period_start"` // Start of the current subscription period
	CurrentPeriodEnd   time.Time  `json:"current_period_end"`   // End of the current subscription period
	Status             int        `json:"status"`               // Current status of the subscription
	CanceledAt         *time.Time `json:"canceled_at"`          // When the subscription was canceled
	Country            *string    `json:"country,omitempty"`    // ISO3166-1 alpha-2 country code of the payment source used to purchase the subscription. Missing unless queried with a private OAuth scope.
}

var SubscriptionTypes map[string]int = map[string]int{
	"ACTIVE":   0,
	"ENDING":   1,
	"INACTIVE": 2,
}
