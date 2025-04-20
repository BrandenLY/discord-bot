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
