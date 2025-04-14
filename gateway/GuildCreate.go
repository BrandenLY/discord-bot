package gateway

import (
	"time"

	"brandenly.com/go/packages/discord-bot/common"
)

// External Reference: https://discord.com/developers/docs/resources/guild#guild-object
type GuildCreate struct {
	Id                          string                `json:"id"`                                      // guild id
	Name                        string                `json:"name"`                                    // guild name (2-100 characters, excluding trailing and leading whitespace)
	Icon                        *string               `json:"icon"`                                    // icon hash
	IconHash                    *string               `json:"icon_hash,omitempty"`                     // icon hash, returned when in the template object
	Splash                      *string               `json:"splash"`                                  // splash hash
	DiscoverySplash             *string               `json:"discover_splash"`                         // discovery splash hash; only present for guilds with the "DISCOVERABLE" feature
	Owner                       *bool                 `json:"owner,omitempty"`                         // true if the user is the owner of the guild
	OwnerId                     string                `json:"owner_id"`                                // id of owner
	Permissions                 *string               `json:"permissions,omitempty"`                   // total permissions for the user in the guild (excludes overwrites and implicit permissions)
	AfkChannelId                *string               `json:"afk_channel_id"`                          // voice region id for the guild (deprecated)
	AfkTimeout                  int                   `json:"afk_timeout"`                             // id of afk channel
	WidgetEnabled               *bool                 `json:"widget_enabled,omitempty"`                // afk timeout in seconds
	WidgetChannelId             *string               `json:"widget_channel_id"`                       // true if the server widget is enabled
	VerificationLevel           int                   `json:"verification_level"`                      // the channel id that the widget will generate an invite to, or null if set to no invite
	DefaultMessageNotifications int                   `json:"default_message_notifications"`           // verification level required for the guild
	ExplicitContentFilter       int                   `json:"explicit_content_filter"`                 // default message notifications level
	Roles                       []common.Role         `json:"roles"`                                   // roles in the guild
	Emojis                      []common.Emoji        `json:"emojis"`                                  // custom guild emojis
	Features                    []string              `json:"features"`                                // enabled guild features
	MfaLevel                    int                   `json:"mfa_level"`                               // required MFA level for the guild
	ApplicationId               *string               `json:"application_id"`                          // application id of the guild creator if it is bot-created
	SystemChannelId             *string               `json:"system_channel_id"`                       // the id of the channel where guild notices such as welcome messages and boost events are posted
	SystemChannelFlags          int                   `json:"system_channel_flags"`                    // system channel flags
	RulesChannelId              *string               `json:"rules_channel_id"`                        // the id of the channel where Community guilds can display rules and/or guidelines
	MaxPresences                *int                  `json:"max_presences"`                           // the maximum number of presences for the guild (null is always returned, apart from the largest of guilds)
	MaxMembers                  int                   `json:"max_members"`                             // the maximum number of members for the guild
	VanityUrlCode               *string               `json:"vanity_url_code"`                         // the vanity url code for the guild
	Description                 *string               `json:"description"`                             // the description of a guild
	Banner                      *string               `json:"banner"`                                  // banner hash
	PremiumTier                 int                   `json:"premium_tier"`                            // premium tier (Server Boost level)
	PremiumSubscriptionCount    *int                  `json:"premium_subscription_count,omitempty"`    // the number of boosts this guild currently has
	PreferredLocale             string                `json:"preferred_locale"`                        // the preferred locale of a Community guild; used in server discovery and notices from Discord, and sent in interactions; defaults to "en-US"
	PublicUpdatesChannelId      *string               `json:"public_updates_channel_id"`               // the id of the channel where admins and moderators of Community guilds receive notices from Discord
	MaxVideoChannelUsers        *int                  `json:"max_video_channel_users,omitempty"`       // the maximum amount of users in a video channel
	MaxStageVideoChannelUsers   *int                  `json:"max_stage_video_channel_users,omitempty"` // the maximum amount of users in a stage video channel
	ApproximateMemberCount      *int                  `json:"approximate_member_count,omitempty"`      // approximate number of members in this guild, returned from the GET /guilds/<id> and /users/@me/guilds endpoints when with_counts is true
	ApproximatePresenceCount    *int                  `json:"approximate_presence_count,omitempty"`    // approximate number of non-offline members in this guild, returned from the GET /guilds/<id> and /users/@me/guilds endpoints when with_counts is true
	WelcomeScreen               *common.WelcomeScreen `json:"welcome_screen,omitempty"`                // the welcome screen of a Community guild, shown to new members, returned in an Invite's guild object
	NsfwLevel                   int                   `json:"nsfw_level"`                              // guild NSFW level
	Stickers                    *[]common.Sticker     `json:"stickers,omitempty"`                      // custom guild stickers
	PremiumProgressBarEnabled   bool                  `json:"premium_progress_bar_enabled"`            // whether the guild has the boost progress bar enabled
	SafetyAlertsChannelId       *string               `json:"safety_alerts_channel_id"`                // the id of the channel where admins and moderators of Community guilds receive safety alerts from Discord
	IncidentsData               *common.Incidents     `json:"incidents_data"`                          // the incidents data for this guild

	JoinedAt             time.Time                    `json:"joined_at"`              // When this guild was joined at
	Large                bool                         `json:"large"`                  // true if this is considered a large guild
	Unavailable          *bool                        `json:"unavailable"`            // true if this guild is unavailable due to an outage
	Membercount          int                          `json:"member_count"`           // Total number of members in this guild
	VoiceStates          []common.VoiceState          `json:"voice_states"`           // States of members currently in voice channels; lacks the guild_id key
	Members              []common.Member              `json:"members"`                // Users in the guild
	Channel              []common.Channel             `json:"channels"`               // Channels in the guild
	Threads              []common.Channel             `json:"threads"`                // All active threads in the guild that current user has permission to view
	Presences            []common.Presence            `json:"presences"`              // Presences of the members in the guild, will only include non-offline members if the size is greater than large threshold
	StageInstances       []common.Stage               `json:"stage_instances"`        // Stage instances in the guild
	GuildScheduledEvents []common.GuildScheduledEvent `json:"guild_Scheduled_events"` // Scheduled events in the guild
	SoundboardSounds     []common.SoundboardSound     `json:"soundboard_sounds"`      // Soundboard sounds in the guild
}
