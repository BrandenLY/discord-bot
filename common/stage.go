package common

// External reference: https://discord.com/developers/docs/resources/stage-instance#stage-instance-object
type Stage struct {
	Id                    string  `json:"id"`                       // The id of this Stage instance
	GuildId               string  `json:"guild_id"`                 // The guild id of the associated Stage channel
	ChannelId             string  `json:"channel_id"`               // The id of the associated Stage channel
	Topic                 string  `json:"topic"`                    // The topic of the Stage instance (1-120 characters)
	PrivacyLevel          int     `json:"privacy_level"`            // The privacy level of the Stage instance
	DiscoverableDisabled  bool    `json:"discoverable_disabled"`    // Whether or not Stage Discovery is disabled (deprecated)
	GuildScheduledEventId *string `json:"guild_scheduled_event_id"` // The id of the scheduled event for this Stage instance
}
