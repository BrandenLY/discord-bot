package common

const ( // Application Command Types
	ChatInputApplicationCommandType         = 1
	UserApplicationCommandtype              = 2
	MessageApplicationCommandType           = 3
	PrimaryEntryPointApplicationCommandType = 4
)

const ( // Interaction Context Types
	GuildInteractionContextType          = 0
	BotDmInteractionContextType          = 1
	PrivateChannelInteractionContextType = 2
)

const ( // Application Command Handler Types
	AppHandlerCommandHandlerType            = 1
	DiscordLaunchActivityCommandHandlerType = 2
)

// External reference: https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-structure
type ApplicationCommand struct {
	Id                       string                      `json:"id"`                                  // Unique ID of command
	Type                     *uint8                      `json:"type,omitempty"`                      // Type of command, defaults to 1
	ApplicationId            string                      `json:"application_id"`                      // ID of the parent application
	GuildId                  *string                     `json:"guild_id,omitempty"`                  // Guild ID of the command, if not global
	Name                     string                      `json:"name"`                                // Name of command, 1-32 characters
	NameLocalizations        *map[string]string          `json:"name_localizations,omitempty"`        // Localization dictionary for name field. Values follow the same restrictions as name
	Description              string                      `json:"description"`                         // Description for CHAT_INPUT commands, 1-100 characters. Empty string for USER and MESSAGE commands
	DescriptionLocalizations *map[string]string          `json:"description_localizations,omitempty"` // Localization dictionary for description field. Values follow the same restrictions as description
	Options                  *[]ApplicationCommandOption `json:"options,omitempty"`                   // Parameters for the command, max of 25
	DefaultMemberPermissions *string                     `json:"default_member_permissions"`          // Set of permissions represented as a bit set
	DmPermissions            *bool                       `json:"dm_permission,omitempty"`             // Deprecated (use contexts instead); Indicates whether the command is available in DMs with the app, only for globally-scoped commands. By default, commands are visible.
	DefaultPermission        *bool                       `json:"default_permission,omitempty"`        // Not recommended for use as field will soon be deprecated. Indicates whether the command is enabled by default when the app is added to a guild, defaults to true
	Nsfw                     *bool                       `json:"nsfw,omitempty"`                      // Indicates whether the command is age-restricted, defaults to false
	IntegrationTypes         *[]uint8                    `json:"integration_types,omitempty"`         // Installation contexts where the command is available, only for globally-scoped commands. Defaults to your app's configured contexts
	Contexts                 *[]uint8                    `json:"contexts,omitempty"`                  // Interaction context(s) where the command can be used, only for globally-scoped commands. By default, all interaction context types included for new commands.
	Version                  string                      `json:"version"`                             // Autoincrementing version identifier updated during substantial record changes
	Handler                  *uint8                      `json:"handler,omitempty"`                   // Determines whether the interaction is handled by the app's interactions handler or by Discord
}

const (
	SubCommandApplicationCommandOptionType      = 1
	SubCommandGroupApplicationCommandOptionType = 2
	StringApplicationCommandOptionType          = 3
	IntegerApplicationCommandOptionType         = 4
	BooleanApplicationCommandOptionType         = 5
	UserApplicationCommandOptionType            = 6
	ChannelApplicationCommandOptionType         = 7
	RoleApplicationCommandOptionType            = 8
	MentionableApplicationCommandOptionType     = 9
	NumberApplicationCommandOptionType          = 10
	AttachmentApplicationCommandOptionType      = 11
)

// External referencce: https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-structure
type ApplicationCommandOption struct {
	Type                     int                               `json:"type"`                                //
	Name                     string                            `json:"name"`                                //
	NameLocalizations        *map[string]string                `json:"name_localizations,omitempty"`        //
	Description              string                            `json:"description"`                         //
	DescriptionLocalizations *map[string]string                `json:"description_localizations,omitempty"` //
	Required                 *bool                             `json:"required,omitempty"`                  //
	Choices                  *[]ApplicationCommandOptionChoice `json:"choices,omitempty"`                   //
	Options                  *[]ApplicationCommandOption       `json:"options,omitempty"`                   //
	ChannelTypes             *[]int                            `json:"channel_types,omitempty"`             //
}

// External reference:
type ApplicationCommandOptionChoice struct {
	Name              string             `json:"name"`               //
	NameLocalizations *map[string]string `json:"name_localizations"` //
	Value             any                `json:"value"`
}

// External reference: https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object
type ApplicationCommandPermissions struct {
	Id          string                         `json:"id"`             // ID of the command or the application ID
	AppId       string                         `json:"application_id"` // ID of the application the command belongs to
	GuildId     string                         `json:"guild_id"`       // ID of the guild
	Permissions []ApplicationCommandPermission `json:"permissions"`    // Permissions for the command in the guild, max of 100
}

// External reference: https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-application-command-permissions-structure
type ApplicationCommandPermission struct {
	Id         string `json:"id"`         // ID of the role, user, or channel. It can also be a permission constant
	Type       uint8  `json:"type"`       // role (1), user (2), or channel (3)
	Permission bool   `json:"permission"` // true to allow, false, to disallow
}
