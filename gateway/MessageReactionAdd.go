package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#message-reaction-add
type MessageReactionAdd struct {
	UserId          string         `json:"user_id"`                     // ID of the user
	ChannelId       string         `json:"channel_id"`                  // ID of the channel
	MessageId       string         `json:"message_id"`                  // ID of the message
	GuildId         *string        `json:"guild_id,omitempty"`          // ID of the guild
	Member          *common.Member `json:"member,omitempty"`            // Member who reacted if this happened in a guild
	Emoji           common.Emoji   `json:"emoji"`                       // Emoji used to react - example
	MessageAuthorId *string        `json:"message_author_id,omitempty"` // ID of the user who authored the message which was reacted to
	Burst           bool           `json:"burst"`                       // true if this is a super-reaction
	BurstColors     *[]string      `json:"burst_colors"`                // Colors used for super-reaction animation in "#rrggbb" format
	Type            int            `json:"type"`                        // The type of reaction
}
