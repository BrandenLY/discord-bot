package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-soundboard-sounds-update-guild-soundboard-sounds-update-event-fields
type GuildSoundboardSoundsUpdate struct {
	SoundboardSounds []common.SoundboardSound `json:"soundboard_sounds"` // The guild's soundboard sounds
	GuildId          string                   `json:"guild_id"`          // ID of the guild
}
