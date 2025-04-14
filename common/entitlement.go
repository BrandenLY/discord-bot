package common

import "time"

// External reference: https://discord.com/developers/docs/resources/entitlement#entitlement-object-entitlement-structure
type Entitlement struct {
	Id            string     `json:"id"`                 // ID of the entitlement
	SkuId         string     `json:"sku_id"`             // ID of the SKU
	ApplicationId string     `json:"application_id"`     // ID of the parent application
	UserId        *string    `json:"user_id,omitempty"`  // ID of the user that is granted access to the entitlement's sku
	Type          int        `json:"type"`               // Type of entitlement
	Deleted       bool       `json:"deleted"`            // Entitlement was deleted
	StartsAt      *time.Time `json:"starts_at"`          // Start date at which the entitlement is valid.
	EndsAt        *time.Time `json:"ends_at"`            // Date at which the entitlement is no longer valid.
	GuildId       *string    `json:"guild_id,omitempty"` // ID of the guild that is granted access to the entitlement's sku
	Consumed      *bool      `json:"consumed,omitempty"` // For consumable items, whether or not the entitlement has been consumed
}
