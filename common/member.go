package common

import "time"

// External reference: https://discord.com/developers/docs/resources/guild#guild-member-object
type Member struct {
	User                       *User                 `json:"user,omitempty"`
	Nick                       *string               `json:"nick,omitempty"`
	Avatar                     *string               `json:"avatar,omitempty"`
	Banner                     *string               `json:"banner,omitempty"`
	Roles                      []string              `json:"roles"`
	JoinedAt                   time.Time             `json:"joined_at"`
	PremiumSince               *time.Time            `json:"premium_since,omitempty"`
	Deaf                       bool                  `json:"deaf"`
	Mute                       bool                  `json:"mute"`
	Flags                      int                   `json:"flags"`
	Pending                    *bool                 `json:"pending,omitempty"`
	Permissions                *string               `json:"permissions,omitempty"`
	CommunicationDisabledUntil *time.Time            `json:"communication_disabled_until,omitempty"`
	AvatarDecorationData       *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
}

var MemberPermissions map[string]int = map[string]int{
	"CREATE_INSTANT_INVITE":               1 << 0,
	"KICK_MEMBERS":                        1 << 1,
	"BAN_MEMBERS":                         1 << 2,
	"ADMINISTRATOR":                       1 << 3,
	"MANAGE_CHANNELS":                     1 << 4,
	"MANAGE_GUILD":                        1 << 5,
	"ADD_REACTIONS":                       1 << 6,
	"VIEW_AUDIT_LOG":                      1 << 7,
	"PRIORITY_SPEAKER":                    1 << 8,
	"STREAM":                              1 << 9,
	"VIEW_CHANNEL":                        1 << 10,
	"SEND_MESSAGES":                       1 << 11,
	"SEND_TTS_MESSAGES":                   1 << 12,
	"MANAGE_MESSAGES":                     1 << 13,
	"EMBED_LINKS":                         1 << 14,
	"ATTACH_FILES":                        1 << 15,
	"READ_MESSAGE_HISTORY":                1 << 16,
	"MENTION_EVERYONE":                    1 << 17,
	"USE_EXTERNAL_EMOJIS":                 1 << 18,
	"VIEW_GUILD_INSIGHTS":                 1 << 19,
	"CONNECT":                             1 << 20,
	"SPEAK":                               1 << 21,
	"MUTE_MEMBERS":                        1 << 22,
	"DEAFEN_MEMBERS":                      1 << 23,
	"MOVE_MEMBERS":                        1 << 24,
	"USE_VAD":                             1 << 25,
	"CHANGE_NICKNAME":                     1 << 26,
	"MANAGE_NICKNAMES":                    1 << 27,
	"MANAGE_ROLES":                        1 << 28,
	"MANAGE_WEBHOOKS":                     1 << 29,
	"MANAGE_GUILD_EXPRESSIONS":            1 << 30,
	"USE_APPLICATION_COMMANDS":            1 << 31,
	"REQUEST_TO_SPEAK":                    1 << 32,
	"MANAGE_EVENTS":                       1 << 33,
	"MANAGE_THREADS":                      1 << 34,
	"CREATE_PUBLIC_THREADS":               1 << 35,
	"CREATE_PRIVATE_THREADS":              1 << 36,
	"USE_EXTERNAL_STICKERS":               1 << 37,
	"SEND_MESSAGES_IN_THREADS":            1 << 38,
	"USE_EMBEDDED_ACTIVITIES":             1 << 39,
	"MODERATE_MEMBERS":                    1 << 40,
	"VIEW_CREATOR_MONETIZATION_ANALYTICS": 1 << 41,
	"USE_SOUNDBOARD":                      1 << 42,
	"CREATE_GUILD_EXPRESSIONS":            1 << 43,
	"CREATE_EVENTS":                       1 << 44,
	"USE_EXTERNAL_SOUNDS":                 1 << 45,
	"SEND_VOICE_MESSAGES":                 1 << 46,
	"SEND_POLLS":                          1 << 49,
	"USE_EXTERNAL_APPS":                   1 << 50,
}
