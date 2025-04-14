package common

import "time"

// External reference: https://discord.com/developers/docs/resources/channel#channel-object
type Channel struct {
	Id                            string           `json:"id"`
	Type                          int              `json:"type"`
	GuildId                       *string          `json:"guild_id,omitempty"`
	Position                      *int             `json:"position,omitempty"`
	PermissionOverwrites          *[]Overwrite     `json:"permission_overwrites,omitempty"`
	Name                          *string          `json:"name,omitempty"`
	Topic                         *string          `json:"topic,omitempty"`
	Nsfw                          *bool            `json:"nsfw,omitempty"`
	LastMessageId                 *string          `json:"last_message_id,omitempty"`
	Bitrate                       *int             `json:"bitrate,omitempty"`
	UserLimit                     *int             `json:"user_limit,omitempty"`
	RateLimitPerUser              *int             `json:"rate_limit_per_user,omitempty"`
	Recipients                    *[]User          `json:"recipients,omitempty"`
	Icon                          *string          `json:"icon,omitempty"`
	OwnerId                       *string          `json:"owner_id,omitempty"`
	ApplicationId                 *string          `json:"application_id,omitempty"`
	Managed                       *bool            `json:"managed,omitempty"`
	ParentId                      *string          `json:"parent_id,omitempty"`
	LastPinTimestamp              *time.Time       `json:"last_pin_timestamp,omitempty"`
	RtcRegion                     *string          `json:"rtc_region,omitempty"`
	VideoQualityMode              *int             `json:"video_quality_mode,omitempty"`
	MessageCount                  *int             `json:"message_count,omitempty"`
	MemberCount                   *int             `json:"member_count,omitempty"`
	ThreadMetadata                *ThreadMetadata  `json:"thread_metadata,omitempty"`
	Member                        *Member          `json:"member,omitempty"`
	DefaultAutoArchiveDuration    *int             `json:"default_auto_archive_duration"`
	Permissions                   *string          `json:"permissions,omitempty"`
	Flags                         *int             `json:"flags,omitempty"`
	TotalMessageSent              *int             `json:"total_messages_sent,omitempty"`
	AvailableTags                 *[]Tag           `json:"available_tags,omitempty"`
	AppliedTags                   *[]string        `json:"applied_tags,omitempty"`
	DefaultReactionEmoji          *DefaultReaction `json:"default_reaction_emoji"`
	DefaultThreadRateLimitPerUser *int             `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              *int             `json:"default_sort_order,omitempty"`
	DefaultForumLayout            *int             `json:"default_forum_layout"`

	NewlyCreated *bool         `json:"newly_created,omitempty"` // When a thread is created, includes an additional newly_created boolean field.
	ThreadMember *ThreadMember `json:"thread_member"`           //  When being added to an existing private thread, includes a thread member object.
}

// External reference: https://discord.com/developers/docs/resources/channel#thread-member-object-thread-member-structure
type ThreadMember struct {
	Id            *string   `json:"id,omitempty"`       // ID of the thread
	UserId        *string   `json:"user_id,omitempty"`  // ID of the user
	JoinTimestamp time.Time `json:"join_timestamp"`     // Time the user last joined the thread
	Flags         int       `json:"flags"`              // Any user-thread settings, currently only used for notifications
	Member        *Member   `json:"member"`             // Additional information about the user
	GuildId       *string   `json:"guild_id,omitempty"` // ID of the guild, sent as part of Thread Member Update gateway events
	Presence      *Presence `json:"presence,omitempty"` // Included in the thread members update gateway event
}

// External reference: https://discord.com/developers/docs/resources/channel#overwrite-object
type Overwrite struct {
	Id    string `json:"id"`
	Type  int    `json:"type"`
	Allow string `json:"allow"`
	Deny  string `json:"deny"`
}

// External reference: https://discord.com/developers/docs/resources/channel#thread-metadata-object
type ThreadMetadata struct {
	Archived            bool       `json:"archived"`
	AutoArchiveDuration int        `json:"auto_archive_duration"`
	ArchiveTimestamp    time.Time  `json:"archive_timestamp"`
	Locked              bool       `json:"locked"`
	Invitable           *bool      `json:"invitable,omitempty"`
	CreateTimestamp     *time.Time `json:"create_timestamp,omitempty"`
}

// External reference: https://discord.com/developers/docs/resources/channel#forum-tag-object
type Tag struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Moderated bool    `json:"moderated"`
	EmojiId   *string `json:"emoji_id"`
	EmojiName *string `json:"emoji_name"`
}

// External reference: https://discord.com/developers/docs/resources/channel#default-reaction-object
type DefaultReaction struct {
	EmojiId   *string `json:"emoji_id"`
	EmojiName *string `json:"emoji_name"`
}
