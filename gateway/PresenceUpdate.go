package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#presence-update-presence-update-event-fields
type PresenceUpdate struct {
	User         common.User       `json:"user"`                    // User whose presence is being updated
	GuildId      string            `json:"guild_id"`                // ID of the guild
	Status       string            `json:"status"`                  // Either "idle", "dnd", "online", or "offline"
	Activities   []common.Activity `json:"activities"`              // User's current activities
	ClientStatus ClientStatus      `json:"client_status,omitempty"` // User's platform-dependent status
}

// External reference: https://discord.com/developers/docs/events/gateway-events#client-status-object
type ClientStatus struct {
	Desktop *string `json:"desktop,omitempty"` // User's status set for an active desktop (Windows, Linux, Mac) application session
	Mobile  *string `json:"mobile,omitempty"`  // User's status set for an active mobile (iOS, Android) application session
	Web     *string `json:"web,omitempty"`     // User's status set for an active web (browser, bot user) application session
}
