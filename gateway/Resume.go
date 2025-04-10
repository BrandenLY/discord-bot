package gateway

// Reference: https://discord.com/developers/docs/events/gateway-events#resume

type Resume struct {
	Token     string `json:"token"`      // Session token
	SessionId string `json:"session_id"` // Session ID
	Seq       int    `json:"seq"`        // Last sequence number received
}
