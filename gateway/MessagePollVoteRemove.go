package gateway

// External reference: https://discord.com/developers/docs/events/gateway-events#message-poll-vote-remove
type MessagePollVoteRemove struct {
	UserId    string  `json:"user_id"`    // ID of the user
	ChannelId string  `json:"channel_id"` // ID of the channel
	MessageId string  `json:"message_id"` // ID of the message
	GuildId   *string `json:"guild_id"`   // ID of the guild
	AnswerId  int     `json:"answer_id"`  // ID of the answer
}
