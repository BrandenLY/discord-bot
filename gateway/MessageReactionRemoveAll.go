package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#message-reaction-remove-all-message-reaction-remove-all-event-fields
type MessageReactionRemoveAll struct {
	ChannelId string  `json:"channel_id"`         // ID of the channel
	MessageId string  `json:"message_id"`         // ID of the message
	GuildId   *string `json:"guild_id,omitempty"` // ID of the guild
}
