package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-ban-add-guild-ban-add-event-fields
type GuildBanAdd struct {
	GuildId string      `json:"guild_id"` // ID of the guild
	User    common.User `json:"user"`     // User who was banned
}

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-ban-remove
type GuildBanRemove struct {
	GuildId string      `json:"guild_id"` // ID of the guild
	User    common.User `json:"user"`     // User who was unbanned
}
