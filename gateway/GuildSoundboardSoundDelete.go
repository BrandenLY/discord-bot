package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-soundboard-sound-delete-guild-soundboard-sound-delete-event-fields
type GuildSoundboardSoundDelete struct {
	SoundId string `json:"sound_id"` // ID of the sound that was deleted
	GuildId string `json:"guild_id"` // ID of the guild the sound was in
}
