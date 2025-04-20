package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#message-delete
type MessageDelete struct {
	Id        string  `json:"id"`                 // ID of the message
	ChannelId string  `json:"channel_id"`         // ID of the channel
	GuildId   *string `json:"guild_id,omitempty"` // ID of the guild
}
