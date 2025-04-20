package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#message-delete-bulk
type MessageDeleteBulk struct {
	Ids       []string `json:"ids"`                // IDs of the messages
	ChannelId string   `json:"channel_id"`         // ID of the channel
	GuildId   *string  `json:"guild_id,omitempty"` // ID of the guild
}
