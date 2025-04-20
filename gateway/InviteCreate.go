package gateway

import (
	"encoding/json"
	"time"

	"brandenly.com/go/packages/discord-bot/common"
)

// External reference: https://discord.com/developers/docs/events/gateway-events#invite-create
type InviteCreate struct {
	ChannelId         string           `json:"channel_id"`                   // Channel the invite is for
	Code              string           `json:"code"`                         // Unique invite code
	CreatedAt         time.Time        `json:"created_at"`                   // Time at which the invite was created
	GuildId           *string          `json:"guild_id,omitempty"`           // Guild of the invite
	Inviter           *common.User     `json:"inviter,omitempty"`            // User that created the invite
	MaxAge            int              `json:"max_age"`                      // How long the invite is valid for (in seconds)
	MaxUses           int              `json:"max_uses"`                     // Maximum number of times the invite can be used
	TargetType        *int             `json:"target_type,omitempty"`        // Type of target for this voice channel invite
	TargetUser        *common.User     `json:"target_user,omitempty"`        // User whose stream to display for this voice channel stream invite
	TargetApplication *json.RawMessage `json:"target_application,omitempty"` // Embedded application to open for this voice channel embedded application invite
	Temporary         bool             `json:"temporary"`                    // Whether or not the invite is temporary (invited users will be kicked on disconnect unless they're assigned a role)
	Uses              int              `json:"uses"`                         // How many times the invite has been used (always will be 0)
}
