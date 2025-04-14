package gateway

import "time"

// External reference: https://discord.com/developers/docs/events/gateway-events#channel-pins-update-channel-pins-update-event-fields
type ChannelPinsUpdate struct {
	GuildId          *string    `json:"guild_id,omitempty"`           // ID of the guild
	ChannelId        *string    `json:"channel_id,omitempty"`         // ID of the channel
	LastPinTimestamp *time.Time `json:"last_pin_timestamp,omitempty"` // Time at which the most recent pinned message was pinned
}
