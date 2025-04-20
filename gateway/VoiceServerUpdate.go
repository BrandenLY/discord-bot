package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#voice-server-update
type VoiceServerUpdate struct {
	Token    string  `json:"token"`              // Voice connection token
	GuildId  string  `json:"guild_id"`           // Guild this voice server update is for
	Endpoint *string `json:"endpoint,omitempty"` // Voice server host
}
