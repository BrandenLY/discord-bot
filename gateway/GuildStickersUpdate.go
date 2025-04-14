package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-stickers-update
type GuildStickersUpdate struct {
	GuildId  string           `json:"guild_id"` // ID of the guild
	Stickers []common.Sticker `json:"stickers"` // Array of stickers
}
