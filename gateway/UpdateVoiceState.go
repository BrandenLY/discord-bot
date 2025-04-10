package gateway

// Reference: https://discord.com/developers/docs/events/gateway-events#update-voice-state

type UpdateVoiceState struct {
	GuildId   uint64  `json:"guild_id"`             // ID of the guild
	ChannelId *uint64 `json:"channel_id,omitempty"` // ID of the voice channel client wants to join (null if disconnecting)
	SelfMute  bool    `json:"self_mute"`            // Whether the client is muted
	SelfDeaf  bool    `json:"self_deaf"`            //	Whether the client deafened
}
