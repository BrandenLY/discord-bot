package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#message-reaction-remove-emoji
type MessageReactionRemoveEmoji struct {
	ChannelId string       `json:"channel_id"`         // ID of the channel
	GuildId   *string      `json:"guild_id,omitempty"` // ID of the guild
	MessageId string       `json:"message_id"`         // ID of the message
	Emoji     common.Emoji `json:"emoji"`              // Emoji that was removed
}
