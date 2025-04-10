package common

// Reference: https://discord.com/developers/docs/topics/permissions#role-object
type Role struct {
	Id           uint64    `json:"id"`                      // role id
	Name         string    `json:"name"`                    // role name
	Color        int       `json:"color"`                   // integer representation of hexadecimal color code
	Hoist        bool      `json:"hoist"`                   // if this role is pinned in the user listing
	Icon         *string   `json:"icon,omitempty"`          // role icon hash
	UnicodeEmoji *string   `json:"unicode_emoji,omitempty"` // role unicode emoji
	Position     int       `json:"position"`                // position of this role (roles with the same position are sorted by id)
	Permissions  string    `json:"permissions"`             // permission bit set
	Managed      bool      `json:"managed"`                 // whether this role is managed by an integration
	Mentionable  bool      `json:"mentionable"`             // whether this role is mentionable
	Tags         *RoleTags `json:"tags,omitempty"`          // the tags this role has
	Flags        int       `json:"flags"`                   // role flags combined as a bitfield
}

// Reference: https://discord.com/developers/docs/topics/permissions#role-object-role-tags-structure
type RoleTags struct {
	BotId                 *uint64 `json:"bot_id,omitempty"`                  // the id of the bot this role belongs to
	IntegrationId         *uint64 `json:"integration_id,omitempty"`          // the id of the integration this role belongs to
	SubscriptionListingId *uint64 `json:"subscription_listing_id,omitempty"` // the id of this role's subscription sku and listing

	//Below will be present and set to null if they are "true", and will be not present if they are "false".

	PremiumSubscriber    *string `json:"premium_subscriber"`     // whether this is the guild's Booster role
	AvailableForPurchase *string `json:"available_for_purchase"` // whether this role is available for purchase
	GuildConnections     *string `json:"guild_connections"`      // whether this role is a guild's linked role
}
