package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-member-remove
type GuildMemberRemove struct {
	GuildId string      `json:"guild_id"` // ID of the guild
	User    common.User `json:"user"`     // User who was removed
}
