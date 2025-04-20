package gateway

import (
	"encoding/json"
	"fmt"
	"time"

	"brandenly.com/go/packages/discord-bot/common"
)

var InteractionType map[string]uint8 = map[string]uint8{
	"PING":                             1,
	"APPLICATION_COMMAND":              2,
	"MESSAGE_COMPONENT":                3,
	"APPLICATION_COMMAND_AUTOCOMPLETE": 4,
	"MODAL_SUBMIT":                     5,
}

var InteractionContextType map[string]uint8 = map[string]uint8{
	"GUILD":           0,
	"BOT_DM":          1,
	"PRIVATE_CHANNEL": 2,
}

type InteractionCreate struct {
	Id                           string            `json:"id"`
	ApplicationId                string            `json:"application_id"`
	Type                         uint8             `json:"type"`
	Data                         *json.RawMessage  `json:"data,omitempty"`
	Guild                        *common.Guild     `json:"guild,omitempty"`
	GuildId                      *string           `json:"guild_id,omitempty"`
	Channel                      *common.Channel   `json:"channel,omitempty"`
	ChannelId                    *string           `json:"channel_id,omitempty"`
	Member                       *common.Member    `json:"member,omitempty"`
	User                         *common.User      `json:"user,omitempty"`
	Token                        string            `json:"token"`
	Version                      uint              `json:"version"`
	Message                      *common.Message   `json:"message,omitempty"`
	AppPermissions               string            `json:"app_permissions"`
	Locale                       *string           `json:"locale,omitempty"`
	GuildLocale                  *string           `json:"guild_locale,omitempty"`
	Entitlements                 []Entitlement     `json:"entitlments"`
	AuthorizingIntegrationOwners map[string]string `json:"authorizing_integration_owners"`
	Context                      *int              `json:"context,omitempty"`
}

type InteractionApplicationCommandData struct {
	Id       string                                   `json:"id"`
	Name     string                                   `json:"name"`
	Type     int                                      `json:"type"`
	Resolved *ResolvedData                            `json:"resolved,omitempty"`
	Options  *[]common.ApplicationCommandOptionChoice `json:"options,omitempty"`
	GuildId  *string                                  `json:"guild_id,omitempty"`
	TargetId *string                                  `json:"target_id,omitempty"`
}

type MessageComponentData struct {
	CustomId      string           `json:"custom_id"`
	ComponentType int              `json:"component_type"`
	Values        *json.RawMessage `json:"values,omitempty"`
	Resolved      *ResolvedData    `json:"resolved,omitempty"`
}

// Reference: https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-option-structure
type SelectOptionValue struct {
}

type ModalSubmitData struct {
	CustomId   string                    `json:"custom_id"`
	Components []common.MessageComponent `json:"components"`
}

// Reference: https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-resolved-data-structure
type ResolvedData struct {
	Users       *map[string]common.User              `json:"users,omitempty"`
	Members     *map[string]common.Member            `json:"members,omitempty"`
	Roles       *map[string]common.Role              `json:"roles,omitempty"`
	Channels    *map[string]common.Channel           `json:"channels,omitempty"`
	Messages    *map[string]common.Message           `json:"messages,omitempty"`
	Attachments *map[string]common.MessageAttachment `json:"attachments,omitempty"`
}

type MessageInteraction struct {
	Id     string         `json:"id"`
	Type   uint8          `json:"type"`
	Name   string         `json:"name"`
	User   common.User    `json:"user"`
	Member *common.Member `json:"member"`
}

var EntitlementTypes map[string]int = map[string]int{
	"PURCHASE":                 1,
	"PREMIUM_SUBSCRIPTION":     2,
	"DEVELOPER_GIFT":           3,
	"TEST_MODE_PURCHASE":       4,
	"FREE_PURCHASE":            5,
	"USER_GIFT":                6,
	"PREMIUM_PURCHASE":         7,
	"APPLICATION_SUBSCRIPTION": 8,
}

// Reference: https://discord.com/developers/docs/resources/entitlement#entitlement-object
type Entitlement struct {
	Id            string     `json:"id"`                 // ID of the entitlement
	SkuId         string     `json:"sku_id"`             // ID of the SKU
	ApplicationId string     `json:"application_id"`     // ID of the parent application
	UserId        *string    `json:"user_id,omitempty"`  // ID of the user that is granted access to the entitlement's sku
	Type          int        `json:"type"`               // Type of entitlement
	Deleted       bool       `json:"deleted"`            // Entitlement was deleted
	StartsAt      *time.Time `json:"starts_at"`          // Start date at which the entitlement is valid.
	EndsAt        *time.Time `json:"ends_at"`            // Date at which the entitlement is no longer valid.
	GuildId       *string    `json:"guild_id,omitempty"` // ID of the guild that is granted access to the entitlement's sku
	Consumed      *bool      `json:"consumed,omitempty"` // For consumable items, whether or not the entitlement has been consumed
}

// Reference:
var InteractionContext map[string]int = map[string]int{
	"GUILD":           0,
	"BOT_DM":          1,
	"PRIVATE_CHANNEL": 2,
}

type GenericCommandOption interface {
	GetValueAsString() string
}

type CommandOptionChoice struct {
	Name  string `json:"name"`  // 1-100 character choice name
	Value any    `json:"value"` // Value for the choice, up to 100 characters if string
	Type  int    `json:"type"`
}

func (o *CommandOptionChoice) GetValueAsString() string {

	var stringValue string = "unsupported type"

	switch o.Type {
	case 1: // SUB_COMMAND
	case 2: // SUB_COMMAND_GROUP
	case 3: // STRING
		stringValue = fmt.Sprintf("%s", o.Value)
	case 4: // INTEGER
		stringValue = fmt.Sprintf("%.0f", o.Value)
	case 5: // BOOL
	case 6: // USER
	case 7: // CHANNEL
	case 8: // ROLE
	case 9: // MENTIONABLE
	case 10: // NUMBER
		stringValue = fmt.Sprintf("%d", o.Value)
	case 11: // ATTACHMENT
	}

	return stringValue
}
