package common

// External reference: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object
type AutoModerationRule struct {
	Id              string          `json:"id"`               // the id of this rule
	GuildId         string          `json:"guild_id"`         // the id of the guild which this rule belongs to
	Name            string          `json:"name"`             // the rule name
	CreatorId       string          `json:"creator_id"`       // the user which first created this rule
	EventType       int             `json:"event_type"`       // the rule event type
	TriggerType     int             `json:"trigger_type"`     // the rule trigger type
	TriggerMetadata TriggerMetadata `json:"trigger_metadata"` // the rule trigger metadata
	Actions         []Action        `json:"actions"`          // the actions which will execute when the rule is triggered
	Enabled         bool            `json:"enabled"`          // whether the rule is enabled
	ExemptRoles     []string        `json:"exempt_roles"`     // the role ids that should not be affected by the rule (Maximum of 20)
	ExemptChannels  []string        `json:"exempt_channels"`  // the channel ids that should not be affected by the rule (Maximum of 50)
}

// External reference: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-trigger-metadata
type TriggerMetadata struct {
	KeywordFilter                []string `json:"keyword_filter"`                  // substrings which will be searched for in content (Maximum of 1000)
	RegexPatterns                []string `json:"regex_patterns"`                  // regular expression patterns which will be matched against content (Maximum of 10)
	Presets                      []uint8  `json:"presets"`                         // the internally pre-defined wordsets which will be searched for in content
	AllowList                    []string `json:"allow_list"`                      // substrings which should not trigger the rule (Maximum of 100 or 1000)
	MentionTotalLimit            int      `json:"mention_total_limit"`             // total number of unique role and user mentions allowed per message (Maximum of 50)
	MentionRaidProtectionEnabled bool     `json:"mention_raid_protection_enabled"` // whether to automatically detect mention raids
}

// External reference: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-action-object-auto-moderation-action-structure
type Action struct {
	Type     uint8           `json:"type"`     // the type of action
	Metadata *ActionMetadata `json:"metadata"` // additional metadata needed during execution for this specific action type
}

// External reference: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-action-object-action-metadata
type ActionMetadata struct {
	ChannelId       string `json:"channel_id"`       // channel to which user content should be logged
	DurationSeconds int    `json:"duration_seconds"` // timeout duration in seconds
	CustomMessage   string `json:"custom_message"`   // additional explanation that will be shown to members whenever their message is blocked
}
