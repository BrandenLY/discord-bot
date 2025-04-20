package gateway

import (
	"time"

	"brandenly.com/go/packages/discord-bot/common"
)

// External reference: https://discord.com/developers/docs/events/gateway-events#message-create
type MessageCreate struct {
	Id                  string                     `json:"id"`
	ChannelId           string                     `json:"channel_id"`
	Author              common.User                `json:"author"`
	Content             string                     `json:"content"`
	Timestamp           time.Time                  `json:"timestamp"`
	EditedTimestamp     *time.Time                 `json:"edited_timestamp"`
	Tts                 bool                       `json:"tts"`
	MentionEveryone     bool                       `json:"mention_everyone"`
	Mentions            []common.User              `json:"mentions"`
	MentionRoles        []common.Role              `json:"mention_roles"`
	MentionChannels     *[]common.ChannelMention   `json:"mention_channels,omitempty"`
	ReferencedMessage   *common.MessageReference   `json:"referenced_message,omitempty"`
	Attachments         []common.MessageAttachment `json:"attachments"`
	Embeds              []common.Embed             `json:"embeds"`
	Interaction         *MessageInteraction        `json:"interaction,omitempty"`
	InteractionMetaData *MessageInteraction        `json:"interaction_metadata,omitempty"`
	AllowedMentions     *common.AllowedMention     `json:"allowed_mentions,omitempty"`
	Components          *[]common.MessageComponent `json:"components,omitempty"`
	Flags               *int                       `json:"flags,omitempty"`

	GuildId *string        `json:"guild_id,omitempty"` // ID of the guild the message was sent in - unless it is an ephemeral message
	Member  *common.Member `json:"member,omitempty"`   // Member properties for this message's author. Missing for ephemeral messages and messages from webhooks
}
