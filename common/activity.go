package common

// Reference: https://discord.com/developers/docs/events/gateway-events#activity-object
type Activity struct {
	Name  string  `json:"name"`            // Activity's name
	Type  uint8   `json:"type"`            // Activity type
	State *string `json:"state,omitempty"` // User's current party status, or text used for a custom status
	Url   *string `json:"url,omitempty"`   // Stream URL, is validated when type is 1
	// Bots cannot modify the fields below
	CreatedAt     int     `json:"created_at"`               // Unix timestamp (in milliseconds) of when the activity was added to the user's session
	ApplicationId *uint64 `json:"application_id,omitempty"` // Application ID for the game
	Details       *string `json:"details,omitempty"`        // What the player is currently doing
	Instance      *bool   `json:"instance,omitempty"`       // Whether or not the activity is an instanced game session
	Flags         *int    `json:"flags,omitempty"`          // Activity flags ORd together, describes what the payload includes
}
