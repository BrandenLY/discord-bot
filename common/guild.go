package common

import "time"

// External Reference: https://discord.com/developers/docs/resources/guild#guild-object
type Guild struct {
	Id                          string         `json:"id"`                                      // guild id
	Name                        string         `json:"name"`                                    // guild name (2-100 characters, excluding trailing and leading whitespace)
	Icon                        *string        `json:"icon"`                                    // icon hash
	IconHash                    *string        `json:"icon_hash,omitempty"`                     // icon hash, returned when in the template object
	Splash                      *string        `json:"splash"`                                  // splash hash
	DiscoverySplash             *string        `json:"discover_splash"`                         // discovery splash hash; only present for guilds with the "DISCOVERABLE" feature
	Owner                       *bool          `json:"owner,omitempty"`                         // true if the user is the owner of the guild
	OwnerId                     uint64         `json:"owner_id"`                                // id of owner
	Permissions                 *string        `json:"permissions,omitempty"`                   // total permissions for the user in the guild (excludes overwrites and implicit permissions)
	AfkChannelId                *uint64        `json:"afk_channel_id"`                          // voice region id for the guild (deprecated)
	AfkTimeout                  int            `json:"afk_timeout"`                             // id of afk channel
	WidgetEnabled               *bool          `json:"widget_enabled,omitempty"`                // afk timeout in seconds
	WidgetChannelId             *uint64        `json:"widget_channel_id"`                       // true if the server widget is enabled
	VerificationLevel           int            `json:"verification_level"`                      // the channel id that the widget will generate an invite to, or null if set to no invite
	DefaultMessageNotifications int            `json:"default_message_notifications"`           // verification level required for the guild
	ExplicitContentFilter       int            `json:"explicit_content_filter"`                 // default message notifications level
	Roles                       []Role         `json:"roles"`                                   // roles in the guild
	Emojis                      []Emoji        `json:"emojis"`                                  // custom guild emojis
	Features                    []string       `json:"features"`                                // enabled guild features
	MfaLevel                    int            `json:"mfa_level"`                               // required MFA level for the guild
	ApplicationId               *uint64        `json:"application_id"`                          // application id of the guild creator if it is bot-created
	SystemChannelId             *uint64        `json:"system_channel_id"`                       // the id of the channel where guild notices such as welcome messages and boost events are posted
	SystemChannelFlags          int            `json:"system_channel_flags"`                    // system channel flags
	RulesChannelId              *uint64        `json:"rules_channel_id"`                        // the id of the channel where Community guilds can display rules and/or guidelines
	MaxPresences                *int           `json:"max_presences"`                           // the maximum number of presences for the guild (null is always returned, apart from the largest of guilds)
	MaxMembers                  int            `json:"max_members"`                             // the maximum number of members for the guild
	VanityUrlCode               *string        `json:"vanity_url_code"`                         // the vanity url code for the guild
	Description                 *string        `json:"description"`                             // the description of a guild
	Banner                      *string        `json:"banner"`                                  // banner hash
	PremiumTier                 int            `json:"premium_tier"`                            // premium tier (Server Boost level)
	PremiumSubscriptionCount    *int           `json:"premium_subscription_count,omitempty"`    // the number of boosts this guild currently has
	PreferredLocale             string         `json:"preferred_locale"`                        // the preferred locale of a Community guild; used in server discovery and notices from Discord, and sent in interactions; defaults to "en-US"
	PublicUpdatesChannelId      *uint64        `json:"public_updates_channel_id"`               // the id of the channel where admins and moderators of Community guilds receive notices from Discord
	MaxVideoChannelUsers        *int           `json:"max_video_channel_users,omitempty"`       // the maximum amount of users in a video channel
	MaxStageVideoChannelUsers   *int           `json:"max_stage_video_channel_users,omitempty"` // the maximum amount of users in a stage video channel
	ApproximateMemberCount      *int           `json:"approximate_member_count,omitempty"`      // approximate number of members in this guild, returned from the GET /guilds/<id> and /users/@me/guilds endpoints when with_counts is true
	ApproximatePresenceCount    *int           `json:"approximate_presence_count,omitempty"`    // approximate number of non-offline members in this guild, returned from the GET /guilds/<id> and /users/@me/guilds endpoints when with_counts is true
	WelcomeScreen               *WelcomeScreen `json:"welcome_screen,omitempty"`                // the welcome screen of a Community guild, shown to new members, returned in an Invite's guild object
	NsfwLevel                   int            `json:"nsfw_level"`                              // guild NSFW level
	Stickers                    *[]Sticker     `json:"stickers,omitempty"`                      // custom guild stickers
	PremiumProgressBarEnabled   bool           `json:"premium_progress_bar_enabled"`            // whether the guild has the boost progress bar enabled
	SafetyAlertsChannelId       *uint64        `json:"safety_alerts_channel_id"`                // the id of the channel where admins and moderators of Community guilds receive safety alerts from Discord
	IncidentsData               *Incidents     `json:"incidents_data"`                          // the incidents data for this guild
}

// External Reference: https://discord.com/developers/docs/resources/guild#welcome-screen-object
type WelcomeScreen struct {
	Description     *string                `json:"description"`      // the server description shown in the welcome screen
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"` // the channels shown in the welcome screen, up to 5
}

// External Reference: https://discord.com/developers/docs/resources/guild#welcome-screen-object-welcome-screen-channel-structure
type WelcomeScreenChannel struct {
	ChannelId   uint64  `json:"channel_id"`  // the channel's id
	Description string  `json:"description"` // the description shown for the channel
	EmojiId     *uint64 `json:"emoji_id"`    // the emoji id, if the emoji is custom
	EmojiName   *string `json:"emoji_name"`  // the emoji name if custom, the unicode character if standard, or null if no emoji is set
}

// External Reference: https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-types
var StickerTypes map[string]int = map[string]int{
	"STANDARD": 1, // an official sticker in a pack
	"GUILD":    2, // a sticker uploaded to a guild for the guild's members
}

// External Reference: https://discord.com/developers/docs/resources/sticker#sticker-object
type Sticker struct {
	Id          uint64  `json:"id"`                   // id of the sticker
	PackId      *uint64 `json:"pack_id,omitempty"`    // for standard stickers, id of the pack the sticker is from
	Name        string  `json:"name"`                 // name of the sticker
	Description *string `json:"description"`          // description of the sticker
	Tags        string  `json:"tags"`                 // autocomplete/suggestion tags for the sticker (max 200 characters)
	Type        int     `json:"type"`                 // type of sticker
	FormatType  int     `json:"format_type"`          // type of sticker format
	Available   *bool   `json:"available,omitempty"`  // whether this guild sticker can be used, may be false due to loss of Server Boosts
	GuildId     *uint64 `json:"guild_id,omitempty"`   // id of the guild that owns this sticker
	User        *User   `json:"user,omitempty"`       // the user that uploaded the guild sticker
	SortValue   *int    `json:"sort_value,omitempty"` // the standard sticker's sort order within its pack
}

// External Reference: https://discord.com/developers/docs/resources/guild#incidents-data-object-incidents-data-structure
type Incidents struct {
	InvitesDisabledUntil *time.Time `json:"invites_disabled_until"`        // when invites get enabled again
	DmsDisabledUntil     *time.Time `json:"dms_disabled_until"`            // when direct messages get enabled again
	DmSpamDetectedAt     *time.Time `json:"dm_spam_detected_at,omitempty"` // when the dm spam was detected
	RaidDetectedAt       *time.Time `json:"raid_detected_at"`              // when the raid was detected
}
