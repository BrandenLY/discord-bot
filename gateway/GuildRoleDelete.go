package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-role-delete
type GuildRoleDelete struct {
	GuildId string `json:"guild_id"` // ID of the guild
	RoleId  string `json:"role_id"`  // ID of the role
}
