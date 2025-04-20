package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#integration-delete
type IntegrationDelete struct {
	Id            string  `json:"id"`                       // Integration ID
	GuildId       string  `json:"guild_id"`                 // ID of the guild
	ApplicationId *string `json:"application_id,omitempty"` // ID of the bot/OAuth2 application for this discord integration
}
