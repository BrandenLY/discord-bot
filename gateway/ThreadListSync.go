package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#thread-list-sync-thread-list-sync-event-fields
type ThreadListSync struct {
	GuildId    string                `json:"guild_id"`    // ID of the guild
	ChannelIds *[]string             `json:"channel_ids"` // Parent channel IDs whose threads are being synced. If omitted, then threads were synced for the entire guild. This array may contain channel_ids that have no active threads as well, so you know to clear that data.
	Threads    []common.Channel      `json:"threads"`     // All active threads in the given channels that the current user can access
	Members    []common.ThreadMember `json:"members"`     // All thread member objects from the synced threads for the current user, indicating which threads the current user has been added to
}
