package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#hello
type Hello struct {
	HeartbeatInterval uint     `json:"heartbeat_interval"` // Interval (in milliseconds) an app should heartbeat with
	Trace             []string `json:"_trace"`
}
