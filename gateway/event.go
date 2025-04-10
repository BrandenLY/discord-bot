package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#gateway-events
type Event struct {
	Op int     `json:"op"`          // Gateway opcode, which indicates the payload type
	D  any     `json:"d,omitempty"` // Event data
	S  *int    `json:"s,omitempty"` // Sequence number of event used for resuming sessions and heartbeating
	T  *string `json:"t,omitempty"` // Event name
}
