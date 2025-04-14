package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-integrations-update
type GuildIntegrationsUpdate struct {
	GuildId string `json:"guild_id"` // ID of the guild
}
