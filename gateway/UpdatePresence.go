package gateway

import "brandenly.com/go/packages/discord-bot/common"

// Reference: https://discord.com/developers/docs/events/gateway-events#update-presence

type UpdatePresence struct {
	Since      *int              `json:"since"`      // Unix time (in milliseconds) of when the client went idle, or null if the client is not idle
	Activities []common.Activity `json:"activities"` // User's activities
	Status     string            `json:"status"`     // User's new status; 'online', 'dnd', 'idle', 'invisible', or 'offline'
	Afk        bool              `json:"afk"`        // Whether or not the client is afk
}
