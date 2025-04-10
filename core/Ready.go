package discord

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#ready
type Ready struct {
	V                int                      `json:"v"`                  // API version
	User             common.User              `json:"user"`               // Information about the user including email
	Guilds           []UnavailableGuild       `json:"guilds"`             // Guilds the user is in
	SessionId        string                   `json:"session_id"`         // Used for resuming connections
	ResumeGatewayUrl string                   `json:"resume_gateway_url"` // Gateway URL for resuming connections
	Shard            [2]int                   `json:"shard"`              // Shard information associated with this session, if sent when identifying
	Application      PartialApplicationObject `json:"application"`        // Contains id and flags
}

// External reference: https://discord.com/developers/docs/resources/guild#unavailable-guild-object
type UnavailableGuild struct {
	Id          string `json:"id"`          // Guild id
	Unavailable bool   `json:"unavailable"` //
}

// External reference: https://discord.com/developers/docs/resources/application#application-object-application-structure
type PartialApplicationObject struct {
	Id    string `json:"id"`    // ID of the app
	Flags uint   `json:"flags"` // App's public flags
}
