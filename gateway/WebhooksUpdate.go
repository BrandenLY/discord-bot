package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#webhooks-update
type WebhooksUpdate struct {
	GuildId   string `json:"guild_id"`   // ID of the guild
	ChannelId string `json:"channel_id"` // ID of the channel
}
