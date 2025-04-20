package gateway

import (
	"time"

	"brandenly.com/go/packages/discord-bot/common"
)

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-member-add-guild-member-add-extra-fields
type GuildMemberAdd struct {
	User                       *common.User                 `json:"user,omitempty"`
	Nick                       *string                      `json:"nick,omitempty"`
	Avatar                     *string                      `json:"avatar,omitempty"`
	Banner                     *string                      `json:"banner,omitempty"`
	Roles                      []string                     `json:"roles"`
	JoinedAt                   time.Time                    `json:"joined_at"`
	PremiumSince               *time.Time                   `json:"premium_since,omitempty"`
	Deaf                       bool                         `json:"deaf"`
	Mute                       bool                         `json:"mute"`
	Flags                      int                          `json:"flags"`
	Pending                    *bool                        `json:"pending,omitempty"`
	Permissions                *string                      `json:"permissions,omitempty"`
	CommunicationDisabledUntil *time.Time                   `json:"communication_disabled_until,omitempty"`
	AvatarDecorationData       *common.AvatarDecorationData `json:"avatar_decoration_data,omitempty"`

	GuildId string `json:"guild_id"` // ID of the guild
}
