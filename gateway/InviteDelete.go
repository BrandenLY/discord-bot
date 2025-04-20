package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#invite-delete
type InviteDelete struct {
	ChannelId string `json:"channel_id"`         // Channel of the invite
	GuildId   string `json:"guild_id,omitempty"` // Guild of the invite
	Code      string `json:"code"`               // Unique invite code
}
