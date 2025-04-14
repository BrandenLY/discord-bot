package common

// External reference: https://discord.com/developers/docs/resources/soundboard#soundboard-sound-object
type SoundboardSound struct {
	Name      string  `json:"name"`           // the name of this sound
	SoundId   string  `json:"sound_id"`       // the id of this sound
	Volume    float64 `json:"volume"`         // the volume of this sound, from 0 to 1
	EmojiId   *string `json:"emoji_id"`       // the id of this sound's custom emoji
	EmojiName *string `json:"emoji_name"`     // the unicode character of this sound's standard emoji
	GuildId   string  `json:"guild_id"`       // the id of the guild this sound is in
	Available bool    `json:"available"`      // whether this sound can be used, may be false due to loss of Server Boosts
	User      *User   `json:"user,omitempty"` // the user who created this sound
}
