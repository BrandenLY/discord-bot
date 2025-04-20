package common

import "time"

// External reference: https://discord.com/developers/docs/resources/guild#integration-object
type Integration struct {
	Id                string                  `json:"id"`                            // integration id
	Name              string                  `json:"name"`                          // integration name
	Type              string                  `json:"type"`                          // integration type (twitch, youtube, discord, or guild_subscription)
	Enabled           bool                    `json:"enabled"`                       // is this integration enabled
	Syncing           *bool                   `json:"syncing,omitempty"`             // is this integration syncing
	RoleId            *string                 `json:"role_id,omitempty"`             // id that this integration uses for "subscribers"
	EnabledEmoticons  *bool                   `json:"enabled_emoticons,omitempty"`   // whether emoticons should be synced for this integration (twitch only currently)
	ExpireBehavior    *int                    `json:"expire_behavior,omitempty"`     // the behavior of expiring subscribers
	ExpireGracePeriod *int                    `json:"expire_grace_period,omitempty"` // the grace period (in days) before expiring subscribers
	User              *User                   `json:"user,omitempty"`                // user for this integration
	Account           *IntegrationAccount     `json:"account"`                       // integration account information
	SyncedAt          *time.Time              `json:"synced_at,omitempty"`           // when this integration was last synced
	SubscriberCount   *int                    `json:"subscriber_count,omitempty"`    // how many subscribers this integration has
	Revoked           *bool                   `json:"revoked,omitempty"`             // has this integration been revoked
	Application       *IntegrationApplication `json:"application,omitempty"`         // The bot/OAuth2 application for discord integrations

	GuildId string `json:"guild_id,omitempty"` // Sent when an integration is created or updated.
}

// External reference: https://discord.com/developers/docs/resources/guild#integration-object-integration-expire-behaviors
var IntegrationExpireBehaviors map[string]int = map[string]int{
	"Remove role": 0,
	"Kick":        1,
}

// External reference: https://discord.com/developers/docs/resources/guild#integration-account-object
type IntegrationAccount struct {
	Id   string `json:"id"`   // id of the account
	Name string `json:"name"` // name of the account
}

// External reference: https://discord.com/developers/docs/resources/guild#integration-application-object
type IntegrationApplication struct {
	Id          string  `json:"id"`            // the id of the app
	Name        string  `json:"name"`          // the name of the app
	Icon        *string `json:"icon"`          // the icon hash of the app
	Description string  `json:"description"`   // the description of the app
	Bot         *User   `json:"bot,omitempty"` // the bot associated with this application
}
