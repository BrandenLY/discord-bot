package common

import (
	"time"
)

// Reference: https://discord.com/developers/docs/resources/message#message-object
type Message struct {
	Id                  string              `json:"id"`
	ChannelId           string              `json:"channel_id"`
	Author              User                `json:"author"`
	Content             string              `json:"content"`
	Timestamp           time.Time           `json:"timestamp"`
	EditedTimestamp     *time.Time          `json:"edited_timestamp"`
	Tts                 bool                `json:"tts"`
	MentionEveryone     bool                `json:"mention_everyone"`
	Mentions            []User              `json:"mentions"`
	MentionRoles        []Role              `json:"mention_roles"`
	MentionChannels     *[]ChannelMention   `json:"mention_channels,omitempty"`
	ReferencedMessage   *MessageReference   `json:"referenced_message,omitempty"`
	Attachments         []MessageAttachment `json:"attachments"`
	Embeds              []Embed             `json:"embeds"`
	Interaction         *MessageInteraction `json:"interaction,omitempty"`
	InteractionMetaData *MessageInteraction `json:"interaction_metadata,omitempty"`
	AllowedMentions     *AllowedMention     `json:"allowed_mentions,omitempty"`
	Components          *[]MessageComponent `json:"components,omitempty"`
	Flags               *int                `json:"flags,omitempty"`
}

// Reference: https://discord.com/developers/docs/resources/message#allowed-mentions-object-allowed-mention-types
var AllowedMentionTypes []string = []string{
	"roles",
	"users",
	"everyone",
}

// Reference: https://discord.com/developers/docs/resources/message#allowed-mentions-object-allowed-mentions-structure
type AllowedMention struct {
	Parse       *[]string `json:"parse,omitempty"`        // An array of allowed mention types to parse from the content.
	Roles       *[]string `json:"roles,omitempty"`        // Array of role_ids to mention (Max size of 100)
	Users       *[]string `json:"users,omitempty"`        // Array of user_ids to mention (Max size of 100)
	RepliedUser *bool     `json:"replied_user,omitempty"` // For replies, whether to mention the author of the message being replied to (default false)
}

// Reference: https://discord.com/developers/docs/resources/message#channel-mention-object
type ChannelMention struct {
}

// Reference: https://discord.com/developers/docs/resources/message#attachment-object
type MessageAttachment struct {
}

// Reference: https://discord.com/developers/docs/resources/message#message-reference-structure
type MessageReference struct {
	Type           *int    `json:"type,omitempty"`
	MessageId      *string `json:"message_id,omitempty"`
	ChannelId      *string `json:"channel_id,omitempty"`
	GuildId        *string `json:"guild_id,omitempty"`
	FailIfNotExist *bool   `json:"fail_if_not_exist,omitempty"`
}

var MessageComponentTypes map[string]int = map[string]int{
	"Action Row":         1,
	"Button":             2,
	"String Select":      3,
	"Text Input":         4,
	"User Select":        5,
	"Role Select":        6,
	"Mentionable Select": 7,
	"Channel Select":     8,
}

// Reference: https://discord.com/developers/docs/interactions/message-components#message-components
type MessageComponent struct {
	Type        int                 `json:"type"`
	Components  *[]MessageComponent `json:"components,omitempty"`
	Label       *string             `json:"label,omitempty"`
	Style       *int                `json:"style,omitempty"`
	CustomId    *string             `json:"custom_id,omitempty"`
	Emoji       *Emoji              `json:"emoji,omitempty"`
	SkuId       *string             `json:"sku_id,omitempty"`
	Url         *string             `json:"url,omitempty"`
	Disabled    *bool               `json:"disabled,omitempty"`
	MinLength   *int                `json:"min_length,omitempty"`
	MaxLength   *int                `json:"max_length,omitempty"`
	Placeholder *string             `json:"placeholder,omitempty"`
	Required    *bool               `json:"required,omitempty"`
	Value       *string             `json:"value,omitempty"`
}

type ActionRowComponent struct {
}

type MessageInteraction struct {
	User User   `json:"user"`
	Type int    `json:"type"`
	Name string `json:"name"`
	Id   string `json:"id"`
}
