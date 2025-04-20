package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#typing-start
type TypingStart struct {
	ChannelId string         `json:"channel_id"`         // ID of the channel
	GuildId   *string        `json:"guild_id,omitempty"` // ID of the guild
	UserId    string         `json:"user_id"`            // ID of the user
	Timestamp int            `json:"timestamp"`          // Unix time (in seconds) of when the user started typing
	Member    *common.Member `json:"member,omitempty"`   // Member who started typing if this happened in a guild
}
