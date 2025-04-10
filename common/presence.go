package common

// External reference: https://discord.com/developers/docs/events/gateway-events#update-presence
// Sent by the client to indicate a presence or status update.
type Presence struct {
	Since      *int       `json:"since"`      // Unix time (in milliseconds) of when the client went idle, or null if the client is not idle
	Activities []Activity `json:"activities"` // User's activities
	Status     string     `json:"status"`     // User's new status
	Afk        bool       `json:"afk"`        // Whether or not the client is afk
}
