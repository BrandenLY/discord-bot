package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#soundboard-sounds
type SoundboardSounds struct {
	SoundboardSounds []common.SoundboardSound `json:"soundboard_sounds"` // The guild's soundboard sounds
	GuildId          string                   `json:"guild_id"`          // ID of the guild
}
