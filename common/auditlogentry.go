package common

// External reference: https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object
type AuditLogEntry struct {
	TargetId   *string           `json:"target_id"`         // ID of the affected entity (webhook, user, role, etc.)
	Changes    *[]AuditLogChange `json:"changes,omitempty"` // Changes made to the target_id
	UserId     *string           `json:"user_id"`           // User or app that made the changes
	Id         string            `json:"id"`                // ID of the entry
	ActionType int               `json:"action_type"`       // Type of action that occurred
	Options    any               `json:"options"`           // Additional info for certain event types
	Reason     *string           `json:"reason,omitempty"`  // Reason for the change (1-512 characters)
}

// External reference: https://discord.com/developers/docs/resources/audit-log#audit-log-change-object
type AuditLogChange struct {
	NewValue any    `json:"new_value"` // New value of the key
	OldValue any    `json:"old_value"` // Old value of the key
	Key      string `json:"key"`       // Name of the changed entity, with a few exceptions
}
