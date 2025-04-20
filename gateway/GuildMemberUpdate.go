package gateway

import (
	"time"

	"brandenly.com/go/packages/discord-bot/common"
)

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-member-update
type GuildMemberUpdate struct {
	GuildId                    string                       `json:"guild_id"`                               // ID of the guild
	Roles                      []string                     `json:"roles"`                                  // User role ids
	User                       common.User                  `json:"user"`                                   // User
	Nick                       *string                      `json:"nick,omitempty"`                         // Nickname of the user in the guild
	Avatar                     *string                      `json:"avatar,omitempty"`                       // Member's guild avatar hash
	Banner                     *string                      `json:"banner,omitempty"`                       // Member's guild banner hash
	JoinedAt                   *time.Time                   `json:"joined_at,omitempty"`                    // When the user joined the guild
	PremiumSince               *time.Time                   `json:"premium_since,omitempty"`                // When the user starting boosting the guild
	Deaf                       *bool                        `json:"deaf,omitempty"`                         // Whether the user is deafened in voice channels
	Mute                       *bool                        `json:"mute,omitempty"`                         // Whether the user is muted in voice channels
	Pending                    *bool                        `json:"pending,omitempty"`                      // Whether the user has not yet passed the guild's Membership Screening requirements
	CommunicationDisabledUntil *time.Time                   `json:"communication_disabled_until,omitempty"` // When the user's timeout will expire and the user will be able to communicate in the guild again, null or a time in the past if the user is not timed out
	Flags                      *int                         `json:"flags,omitempty"`                        // Guild member flags represented as a bit set, defaults to 0
	AvatarDecorationData       *common.AvatarDecorationData `json:"avatar_decoration_data,omitempty"`       // Data for the member's guild avatar decoration
}
