package gateway

import (
	"encoding/json"
	"fmt"

	"brandenly.com/go/packages/discord-bot/common"
)

// External reference: https://discord.com/developers/docs/events/gateway-events#gateway-events
type Event struct {
	Op int     `json:"op"`          // Gateway opcode, which indicates the payload type
	D  any     `json:"d,omitempty"` // Event data
	S  *int    `json:"s,omitempty"` // Sequence number of event used for resuming sessions and heartbeating
	T  *string `json:"t,omitempty"` // Event name
}

func (e *Event) UnmarshalJSON(data []byte) error {

	// create placeholder
	var raw struct {
		Op int             `json:"op"`
		D  json.RawMessage `json:"d,omitempty"`
		S  *int            `json:"s,omitempty"`
		T  *string         `json:"t,omitempty"`
	}

	// Unmarshal into placeholder
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	e.Op = raw.Op // This field does not require modification
	e.S = raw.S   // This field does not require modification
	e.T = raw.T   // This field does not require modification

	if raw.Op == 0 { // Dispatch

		constructor, ok := EventTypeStructs[*raw.T]
		if !ok {

			var dataAsFallback any
			if err := json.Unmarshal(raw.D, &dataAsFallback); err != nil {
				return fmt.Errorf("failed to parse event; could not unmarshal: %s", *raw.T)
			}

			e.D = dataAsFallback
			return nil

		}

		instance := constructor()
		if err := json.Unmarshal(raw.D, instance); err != nil {
			return err
		}

		e.D = instance
		return nil
	}

	if raw.Op == 7 { // Reconnect event
		e.D = raw.D
		return nil
	}

	if raw.Op == 9 { // Invalid Session event
		e.D = raw.D
		return nil
	}

	if raw.Op == 10 { // Hello event

		var data Hello
		err := json.Unmarshal(raw.D, &data)
		if err != nil {
			return err
		}

		e.D = data
		return nil // exit

	}

	if raw.Op == 11 { // Heartbeat ACK event
		e.D = raw.D
		return nil
	}

	return nil
}

var EventTypeStructs map[string]func() any = map[string]func() any{
	"HELLO": func() any { return &Hello{} },
	"READY": func() any { return &Ready{} },
	// "RESUMED": ,
	// "RECONNECT": ,
	// "INVALID_SESSION": ,
	"APPLICATION_COMMAND_PERMISSIONS_UPDATE": func() any { return &common.ApplicationCommandPermissions{} },
	"AUTO_MODERATION_RULE_CREATE":            func() any { return &common.AutoModerationRule{} },
	"AUTO_MODERATION_RULE_UPDATE":            func() any { return &common.AutoModerationRule{} },
	"AUTO_MODERATION_RULE_DELETE":            func() any { return &common.AutoModerationRule{} },
	"AUTO_MODERATION_ACTION_EXECUTION":       func() any { return &AutoModerationActionExecution{} },
	"CHANNEL_CREATE":                         func() any { return &common.Channel{} },
	"CHANNEL_UPDATE":                         func() any { return &common.Channel{} },
	"CHANNEL_DELETE":                         func() any { return &common.Channel{} },
	"THREAD_CREATE":                          func() any { return &common.Channel{} },
	"THREAD_UPDATE":                          func() any { return &common.Channel{} },
	"THREAD_DELETE":                          func() any { return &common.Channel{} },
	"CHANNEL_PINS_UPDATE":                    func() any { return &ChannelPinsUpdate{} },
	"THREAD_LIST_SYNC":                       func() any { return &ThreadListSync{} },
	"THREAD_MEMBER_UPDATE":                   func() any { return &common.ThreadMember{} },
	"THREAD_MEMBERS_UPDATE":                  func() any { return &ThreadMembersUpdate{} },
	"ENTITLEMENT_CREATE":                     func() any { return &common.Entitlement{} },
	"ENTITLEMENT_UPDATE":                     func() any { return &common.Entitlement{} },
	"ENTITLEMENT_DELETE":                     func() any { return &common.Entitlement{} },
	"GUILD_CREATE":                           func() any { return &GuildCreate{} },
	"GUILD_UPDATE":                           func() any { return &common.Guild{} },
	"GUILD_DELETE":                           func() any { return &common.UnavailableGuild{} },
	"GUILD_AUDIT_LOG_ENTRY_CREATE":           func() any { return &GuildAuditLogEntryCreate{} },
	"GUILD_BAN_ADD":                          func() any { return &GuildBanAdd{} },
	"GUILD_BAN_REMOVE":                       func() any { return &GuildBanRemove{} },
	"GUILD_EMOJIS_UPDATE":                    func() any { return &GuildEmojisUpdate{} },
	"GUILD_STICKERS_UPDATE":                  func() any { return &GuildStickersUpdate{} },
	"GUILD_INTEGRATIONS_UPDATE":              func() any { return &GuildIntegrationsUpdate{} },
	"GUILD_MEMBER_ADD":                       func() any { return &GuildMemberAdd{} },
	"GUILD_MEMBER_REMOVE":                    func() any { return &GuildMemberRemove{} },
	"GUILD_MEMBER_UPDATE":                    func() any { return &GuildMemberUpdate{} },
	"GUILD_MEMBERS_CHUNK":                    func() any { return &GuildMembersChunk{} },
	"GUILD_ROLE_CREATE":                      func() any { return &GuildRoleCreate{} },
	"GUILD_ROLE_UPDATE":                      func() any { return &GuildRoleUpdate{} },
	"GUILD_ROLE_DELETE":                      func() any { return &GuildRoleDelete{} },
	"GUILD_SCHEDULED_EVENT_CREATE":           func() any { return &common.GuildScheduledEvent{} },
	"GUILD_SCHEDULED_EVENT_UPDATE":           func() any { return &common.GuildScheduledEvent{} },
	"GUILD_SCHEDULED_EVENT_DELETE":           func() any { return &common.GuildScheduledEvent{} },
	"GUILD_SCHEDULED_EVENT_USER_ADD":         func() any { return &GuildScheduledEventUserAdd{} },
	"GUILD_SCHEDULED_EVENT_USER_REMOVE":      func() any { return &GuildScheduledEventUserRemove{} },
	"GUILD_SOUNDBOARD_SOUND_CREATE":          func() any { return &common.SoundboardSound{} },
	"GUILD_SOUNDBOARD_SOUND_UPDATE":          func() any { return &common.SoundboardSound{} },
	"GUILD_SOUNDBOARD_SOUND_DELETE":          func() any { return &GuildSoundboardSoundDelete{} },
	"GUILD_SOUNDBOARD_SOUNDS_UPDATE":         func() any { return &GuildSoundboardSoundsUpdate{} },
	"SOUNDBOARD_SOUNDS":                      func() any { return &SoundboardSounds{} },
	"INTEGRATION_CREATE":                     func() any { return &common.Integration{} },
	"INTEGRATION_UPDATE":                     func() any { return &common.Integration{} },
	"INTEGRATION_DELETE":                     func() any { return &IntegrationDelete{} },
	"INTERACTION_CREATE":                     func() any { return &InteractionCreate{} },
	"INVITE_CREATE":                          func() any { return &InviteCreate{} },
	"INVITE_DELETE":                          func() any { return &InviteDelete{} },
	"MESSAGE_CREATE":                         func() any { return &MessageCreate{} },
	"MESSAGE_UPDATE":                         func() any { return &MessageUpdate{} },
	"MESSAGE_DELETE":                         func() any { return &MessageDelete{} },
	"MESSAGE_DELETE_BULK":                    func() any { return &MessageDeleteBulk{} },
	"MESSAGE_REACTION_ADD":                   func() any { return &MessageReactionAdd{} },
	"MESSAGE_REACTION_REMOVE":                func() any { return &MessageReactionRemove{} },
	"MESSAGE_REACTION_REMOVE_ALL":            func() any { return &MessageReactionRemoveAll{} },
	"MESSAGE_REACTION_REMOVE_EMOJI":          func() any { return &MessageReactionRemoveEmoji{} },
	"PRESENCE_UPDATE":                        func() any { return &PresenceUpdate{} },
	"STAGE_INSTANCE_CREATE":                  func() any { return &common.Stage{} },
	"STAGE_INSTANCE_UPDATE":                  func() any { return &common.Stage{} },
	"STAGE_INSTANCE_DELETE":                  func() any { return &common.Stage{} },
	"SUBSCRIPTION_CREATE":                    func() any { return &common.Subscription{} },
	"SUBSCRIPTION_UPDATE":                    func() any { return &common.Subscription{} },
	"SUBSCRIPTION_DELETE":                    func() any { return &common.Subscription{} },
	"TYPING_START":                           func() any { return &TypingStart{} },
	"USER_UPDATE":                            func() any { return &common.User{} },
	"VOICE_CHANNEL_EFFECT_SEND":              func() any { return &VoiceChannelEffectSend{} },
	"VOICE_STATE_UPDATE":                     func() any { return &common.VoiceState{} },
	"VOICE_SERVER_UPDATE":                    func() any { return &VoiceServerUpdate{} },
	"WEBHOOKS_UPDATE":                        func() any { return &WebhooksUpdate{} },
	"MESSAGE_POLL_VOTE_ADD":                  func() any { return &MessagePollVoteAdd{} },
	"MESSAGE_POLL_VOTE_REMOVE":               func() any { return &MessagePollVoteRemove{} },
}
