package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#auto-moderation-action-execution
type AutoModerationActionExecution struct {
	GuildId              string               `json:"guild_id"`                          // ID of the guild in which action was executed
	Action               AutoModerationAction `json:"action"`                            // Action which was executed
	RuleId               string               `json:"rule_id"`                           // ID of the rule which action belongs to
	RuleTriggerType      int                  `json:"rule_trigger_type"`                 // Trigger type of rule which was triggered
	UserId               string               `json:"user_id"`                           // ID of the user which generated the content which triggered the rule
	ChannelId            *string              `json:"channel_id,omitempty"`              // ID of the channel in which user content was posted
	MessageId            *string              `json:"message_id,omitempty"`              // ID of any user message which content belongs to *
	AlertSystemMessageId *string              `json:"alert_system_message_id,omitempty"` // ID of any system auto moderation messages posted as a result of this action **
	Content              string               `json:"content"`                           // User-generated text content
	MatchedKeyword       *string              `json:"matched_keyword"`                   // Word or phrase configured in the rule that triggered the rule
	MatchedContent       *string              `json:"matched_content"`                   // Substring in content that triggered the rule
}

// External reference:
type AutoModerationAction struct {
}

// External reference: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-trigger-types
var AutoModerationTriggerTypes map[string]int = map[string]int{
	"KEYWORD":        1,
	"SPAM":           3,
	"KEYWORD_PRESET": 4,
	"MENTION_SPAM":   5,
	"MEMBER_PROFILE": 6,
}
