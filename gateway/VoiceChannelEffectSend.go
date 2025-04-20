package gateway

import "brandenly.com/go/packages/discord-bot/common"

// External reference: https://discord.com/developers/docs/events/gateway-events#voice-channel-effect-send
type VoiceChannelEffectSend struct {
	ChannelId     string        `json:"channel_id"`               // ID of the channel the effect was sent in
	GuildId       string        `json:"guild_id"`                 // ID of the guild the effect was sent in
	UserId        string        `json:"user_id"`                  // ID of the user who sent the effect
	Emoji         *common.Emoji `json:"emoji"`                    // The emoji sent, for emoji reaction and soundboard effects
	AnimationType *int          `json:"animation_type,omitempty"` // The type of emoji animation, for emoji reaction and soundboard effects
	AnimationId   *int          `json:"animation_id,omitempty"`   // The ID of the emoji animation, for emoji reaction and soundboard effects
	SoundId       *string       `json:"sound_id,omitempty"`       // The ID of the soundboard sound, for soundboard effects
	SoundVolume   *float64      `json:"sound_volume,omitempty"`   // The volume of the soundboard sound, from 0 to 1, for soundboard effects
}
