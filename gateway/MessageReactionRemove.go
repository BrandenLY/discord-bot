package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#message-reaction-remove
type MessageReactionRemove struct {
	UserId    string       `json:"user_id"`            // ID of the user
	ChannelId string       `json:"channel_id"`         // ID of the channel
	MessageId string       `json:"message_id"`         // ID of the message
	GuildId   *string      `json:"guild_id,omitempty"` // ID of the guild
	Emoji     common.Emoji `json:"emoji"`              // Emoji used to react - example
	Burst     bool         `json:"burst"`              // true if this was a super-reaction
	Type      int          `json:"type"`               // The type of reaction
}
