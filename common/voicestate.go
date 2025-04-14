package common

import (
	"time"
)

// External reference: https://discord.com/developers/docs/resources/voice#voice-state-object
type VoiceState struct {
	GuildId                 *string    `json:"guild_id,omitempty"`         // the guild id this voice state is for
	ChannelId               *string    `json:"channel_id"`                 // the channel id this user is connected to
	UserId                  string     `json:"user_id"`                    // the user id this voice state is for
	Member                  *Member    `json:"member,omitempty"`           // the guild member this voice state is for
	SessionId               string     `json:"session_id"`                 // the session id for this voice state
	Deaf                    bool       `json:"deaf"`                       // whether this user is deafened by the server
	Mute                    bool       `json:"mute"`                       // whether this user is muted by the server
	SelfDeaf                bool       `json:"self_deaf"`                  // whether this user is locally deafened
	SelfMute                bool       `json:"self_mute"`                  // whether this user is locally muted
	SelfStream              bool       `json:"self_stream"`                // whether this user is streaming using "Go Live"
	SelfVideo               bool       `json:"self_video"`                 // whether this user's camera is enabled
	Suppress                bool       `json:"suppres"`                    // whether this user's permission to speak is denied
	RequestToSpeakTimestamp *time.Time `json:"request_to_speak_timestamp"` // the time at which the user requested to speak
}
