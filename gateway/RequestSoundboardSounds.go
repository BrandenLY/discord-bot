package gateway

// Reference: https://discord.com/developers/docs/events/gateway-events#request-soundboard-sounds

type RequestSoundboardSounds struct {
	GuildIds []uint64 `json:"guild_ids"` //	IDs of the guilds to get soundboard sounds for
}
