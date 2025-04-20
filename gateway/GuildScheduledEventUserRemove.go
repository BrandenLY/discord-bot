package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-scheduled-event-user-remove
type GuildScheduledEventUserRemove struct {
	GuildScheduledEventId string `json:"guild_scheduled_event_id"` // ID of the guild scheduled event
	UserId                string `json:"user_id"`                  // ID of the user
	GuildId               string `json:"guild_id"`                 // ID of the guild
}
