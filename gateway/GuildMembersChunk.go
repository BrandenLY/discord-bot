package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#guild-members-chunk
type GuildMembersChunk struct {
	GuildId    string             `json:"guild_id"`            // ID of the guild
	Members    []common.Member    `json:"members"`             // Set of guild members
	ChunkIndex int                `json:"chunk_index"`         // Chunk index in the expected chunks for this response (0 <= chunk\_index < chunk\_count)
	ChunkCount int                `json:"chunk_count"`         // Total number of expected chunks for this response
	NotFound   *[]string          `json:"not_found,omitempty"` // When passing an invalid ID to REQUEST_GUILD_MEMBERS, it will be returned here
	Presences  *[]common.Presence `json:"presences,omitempty"` // When passing true to REQUEST_GUILD_MEMBERS, presences of the returned members will be here
	Nonce      *string            `json:"nonce,omitempty"`     // Nonce used in the Guild Members Request
}
