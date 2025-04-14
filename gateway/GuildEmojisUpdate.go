package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-emojis-update
type GuildEmojisUpdate struct {
	GuildId string         `json:"guild_id"` // ID of the guild
	Emojis  []common.Emoji `json:"emojis"`   // Array of emojis
}
