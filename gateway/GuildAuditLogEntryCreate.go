package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-audit-log-entry-create
type GuildAuditLogEntryCreate struct {
	TargetId   *string                  `json:"target_id"`         // ID of the affected entity (webhook, user, role, etc.)
	Changes    *[]common.AuditLogChange `json:"changes,omitempty"` // Changes made to the target_id
	UserId     *string                  `json:"user_id"`           // User or app that made the changes
	Id         string                   `json:"id"`                // ID of the entry
	ActionType int                      `json:"action_type"`       // Type of action that occurred
	Options    any                      `json:"options"`           // Additional info for certain event types
	Reason     *string                  `json:"reason,omitempty"`  // Reason for the change (1-512 characters)
	GuildId    string                   `json:"guild_id"`          // ID of the guild
}
