package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#thread-members-update
type ThreadMembersUpdate struct {
	Id               string                 `json:"id"`                      // ID of the thread
	GuildId          string                 `json:"guild_id"`                // ID of the guild
	MemberCount      uint                   `json:"member_count"`            // Approximate number of members in the thread, capped at 50
	AddedMembers     *[]common.ThreadMember `json:"added_members,omitempty"` // Users who were added to the thread
	RemovedMemberIds *[]string              `json:"removed_member_ids"`      // ID of the users who were removed from the thread
}
