package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#identify
type Identify struct {
	Token          string                 `json:"token"`                     // Authentication token
	Properties     IdentifyConnProperties `json:"properties"`                // Connection properties
	Compress       *bool                  `json:"compress,omitempty"`        // Whether this connection supports compression of packets
	LargeThreshold *uint8                 `json:"large_threshold,omitempty"` // Value between 50 and 250, total number of members where the gateway will stop sending offline members in the guild member list
	Shard          *[2]int                `json:"shard,omitempty"`           // Used for [Guild Sharding](https://discord.com/developers/docs/events/gateway#sharding)
	Presence       common.Presence        `json:"presence,omitempty"`        // Presence structure for initial presence information
	Intents        int                    `json:"intents"`                   // Gateway Intents you wish to receive
}

// External reference: https://discord.com/developers/docs/events/gateway-events#identify-identify-connection-properties
type IdentifyConnProperties struct {
	Os      string `json:"os"`      // Your operating system
	Browser string `json:"browser"` // Your library name
	Device  string `json:"device"`  // Your library name
}
