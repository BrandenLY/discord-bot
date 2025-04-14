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
	// "AUTO_MODERATION_ACTION_EXECUTION": ,
	"CHANNEL_CREATE":        func() any { return &common.Channel{} },
	"CHANNEL_UPDATE":        func() any { return &common.Channel{} },
	"CHANNEL_DELETE":        func() any { return &common.Channel{} },
	"THREAD_CREATE":         func() any { return &common.Channel{} },
	"THREAD_UPDATE":         func() any { return &common.Channel{} },
	"THREAD_DELETE":         func() any { return &common.Channel{} },
	"CHANNEL_PINS_UPDATE":   func() any { return &ChannelPinsUpdate{} },
	"THREAD_LIST_SYNC":      func() any { return &ThreadListSync{} },
	"THREAD_MEMBER_UPDATE":  func() any { return &common.ThreadMember{} },
	"THREAD_MEMBERS_UPDATE": func() any { return &ThreadMembersUpdate{} },
	"ENTITLEMENT_CREATE":    func() any { return &common.Entitlement{} },
	"ENTITLEMENT_UPDATE":    func() any { return &common.Entitlement{} },
	"ENTITLEMENT_DELETE":    func() any { return &common.Entitlement{} },
	"GUILD_CREATE":          func() any { return &GuildCreate{} },
	"GUILD_UPDATE":          func() any { return &common.Guild{} },
	"GUILD_DELETE":          func() any { return &common.UnavailableGuild{} },
	// "GUILD_AUDIT_LOG_ENTRY_CREATE": ,
	"GUILD_BAN_ADD":             func() any { return &GuildBanAdd{} },
	"GUILD_BAN_REMOVE":          func() any { return &GuildBanRemove{} },
	"GUILD_EMOJIS_UPDATE":       func() any { return &GuildEmojisUpdate{} },
	"GUILD_STICKERS_UPDATE":     func() any { return &GuildStickersUpdate{} },
	"GUILD_INTEGRATIONS_UPDATE": func() any { return &GuildIntegrationsUpdate{} },
	// "GUILD_MEMBER_ADD": ,
	// "GUILD_MEMBER_REMOVE": ,
	// "GUILD_MEMBER_UPDATE": ,
	// "GUILD_MEMBERS_CHUNK": ,
	// "GUILD_ROLE_CREATE": ,
	// "GUILD_ROLE_UPDATE": ,
	// "GUILD_ROLE_DELETE": ,
	// "GUILD_SCHEDULED_EVENT_CREATE": ,
	// "GUILD_SCHEDULED_EVENT_UPDATE": ,
	// "GUILD_SCHEDULED_EVENT_DELETE": ,
	// "GUILD_SCHEDULED_EVENT_USER_ADD": ,
	// "GUILD_SCHEDULED_EVENT_USER_REMOVE": ,
	// "GUILD_SOUNDBOARD_SOUND_CREATE": ,
	// "GUILD_SOUNDBOARD_SOUND_UPDATE": ,
	// "GUILD_SOUNDBOARD_SOUND_DELETE": ,
	// "GUILD_SOUNDBOARD_SOUNDS_UPDATE": ,
	// "SOUNDBOARD_SOUNDS": ,
	// "INTEGRATION_CREATE": ,
	// "INTEGRATION_UPDATE": ,
	// "INTEGRATION_DELETE": ,
	// "INTERACTION_CREATE": ,
	// "INVITE_CREATE": ,
	// "INVITE_DELETE": ,
	// "MESSAGE_CREATE": ,
	// "MESSAGE_UPDATE": ,
	// "MESSAGE_DELETE": ,
	// "MESSAGE_DELETE_BULK": ,
	// "MESSAGE_REACTION_ADD": ,
	// "MESSAGE_REACTION_REMOVE": ,
	// "MESSAGE_REACTION_REMOVE_ALL": ,
	// "MESSAGE_REACTION_REMOVE_EMOJI": ,
	// "PRESENCE_UPDATE": ,
	// "STAGE_INSTANCE_CREATE": ,
	// "STAGE_INSTANCE_UPDATE": ,
	// "STAGE_INSTANCE_DELETE": ,
	// "SUBSCRIPTION_CREATE": ,
	// "SUBSCRIPTION_UPDATE": ,
	// "SUBSCRIPTION_DELETE": ,
	// "TYPING_START": ,
	// "USER_UPDATE": ,
	// "VOICE_CHANNEL_EFFECT_SEND": ,
	// "VOICE_STATE_UPDATE": ,
	// "VOICE_SERVER_UPDATE": ,
	// "WEBHOOKS_UPDATE": ,
	// "MESSAGE_POLL_VOTE_ADD": ,
	// "MESSAGE_POLL_VOTE_REMOVE": ,
}
