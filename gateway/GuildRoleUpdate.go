package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-role-update
type GuildRoleUpdate struct {
	GuildId string      `json:"guild_id"` // ID of the guild
	Role    common.Role `json:"role"`     // Role that was updated
}
