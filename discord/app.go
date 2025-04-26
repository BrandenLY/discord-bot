package discord

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"

	"brandenly.com/go/packages/discord-bot/common"
	"brandenly.com/go/packages/discord-bot/gateway"
)

const DiscordApiBaseUrl = "https://discord.com/api"

var BotIntents map[string]int = map[string]int{
	"GUILDS":                        1 << 0,
	"GUILD_MEMBERS":                 1 << 1,
	"GUILD_BANS":                    1 << 2,
	"GUILD_EMOJIS_AND_STICKERS":     1 << 3,
	"GUILD_INTEGRATIONS":            1 << 4,
	"GUILD_WEBHOOKS":                1 << 5,
	"GUILD_INVITES":                 1 << 6,
	"GUILD_VOICE_STATES":            1 << 7,
	"GUILD_PRESENCES":               1 << 8,
	"GUILD_MESSAGES":                1 << 9,
	"GUILD_MESSAGE_REACTIONS":       1 << 10,
	"GUILD_MESSAGE_TYPING":          1 << 11,
	"DIRECT_MESSAGES":               1 << 12,
	"DIRECT_MESSAGE_REACTIONS":      1 << 13,
	"DIRECT_MESSAGE_TYPING":         1 << 14,
	"MESSAGE_CONTENT":               1 << 15,
	"GUILD_SCHEDULED_EVENTS":        1 << 16,
	"AUTO_MODERATION_CONFIGURATION": 1 << 20,
	"AUTO_MODERATION_EXECUTION":     1 << 21,
}

// External reference: https://discord.com/developers/docs/resources/application#install-params-object
type InstallParams struct {
	Scopes      []string `json:"scopes"`
	Permissions string   `json:"permissions"`
}

// External reference: https://discord.com/developers/docs/resources/application#application-object-application-integration-type-configuration-object
type IntegrationTypesConfig struct {
	Oauth2InstallParams *InstallParams `json:"oauth2_install_params,omitempty"` // Install params for each installation context's default in-app authorization link
}

// External reference: https://discord.com/developers/docs/resources/application#application-object
type App struct {
	Id                uint64      `json:"-" discord-bot:"internal"` // ID of the app
	PublicKey         string      `json:"-" discord-bot:"internal"` // App Public Key
	BotToken          string      `json:"-" discord-bot:"internal"` // App Bot Token
	DiscordApiBaseUrl string      `json:"-" discord-bot:"internal"` // Discord base url; the base url used for requests to discord's REST api.
	Color             int         `json:"-" discord-bot:"internal"` // Color
	Logger            *log.Logger `json:"-" discord-bot:"internal"` // The logger to use for application state changes.

	Name                            string                 `json:"name"`                               // Name of the app
	Icon                            string                 `json:"icon"`                               // Icon hash of the app
	Description                     string                 `json:"description"`                        // Description of the app
	RpcOrigins                      []string               `json:"rpc_origins"`                        // List of RPC origin URLs, if RPC is enabled
	BotPublic                       bool                   `json:"bot_public"`                         // When `false`, only the app owner can add the app to guilds
	BotRequireCodeGrant             bool                   `json:"bot_require_code_grant"`             // When `true`, the app's bot will only join upon completion of the full OAuth2 code grant flow
	Bot                             common.User            `json:"bot"`                                // Partial user object for the bot user associated with the app
	TermsOfServiceUrl               string                 `json:"terms_of_service_url"`               // URL of the app's Terms of Service
	PrivacyPolicyUrl                string                 `json:"privacy_policy_url"`                 // URL of the app's Privacy Policy
	Owner                           common.User            `json:"owner"`                              // Partial user object for the owner of the app
	VerifyKey                       string                 `json:"verify_key"`                         // Hex encoded key for verification in interactions and the GameSDK's GetTicket
	Team                            common.Team            `json:"team"`                               // If the app belongs to a team, this will be a list of the members of that team
	GuildId                         string                 `json:"guild_id"`                           // Guild associated with the app. For example, a developer support server.
	Guild                           common.Guild           `json:"guild"`                              // Partial object of the associated guild
	PrimarySkuId                    string                 `json:"primary_sku_id"`                     // If this app is a game sold on Discord, this field will be the id of the "Game SKU" that is created, if exists
	Slug                            string                 `json:"slug"`                               // If this app is a game sold on Discord, this field will be the URL slug that links to the store page
	CoverImage                      string                 `json:"cover_image"`                        // App's default rich presence invite cover image hash
	Flags                           int                    `json:"flags"`                              // App's public flags
	ApproximateGuildCount           int                    `json:"approximate_guild_count"`            // Approximate count of guilds the app has been added to
	ApproximateUserInstallCount     int                    `json:"approximate_user_install_count"`     // Approximate count of users that have installed the app
	RedirectUris                    []string               `json:"redirect_uris"`                      // Array of redirect URIs for the app
	InteractionsEndpointUrl         string                 `json:"interactions_endpoint_url"`          // Interactions endpoint URL for the app
	RoleConnectionsVerificationsUrl string                 `json:"role_connections_verifications_url"` // Role connection verification URL for the app
	EventWebhooksUrl                string                 `json:"event_webhooks_url"`                 // Event webhooks URL for the app to receive webhook events
	EventWebhooksStatus             int8                   `json:"event_webhooks_status"`              // If webhook events are enabled for the app. `1` (default) means disabled, `2` means enabled, and `3` means disabled by Discord
	EventWebhooksTypes              []string               `json:"event_webhooks_types"`               // List of Webhook event types the app subscribes to
	Tags                            []string               `json:"tags"`                               // List of tags describing the content and functionality of the app. Max of 5 tags.
	InstallParams                   InstallParams          `json:"install_params"`                     // Settings for the app's default in-app authorization link, if enabled
	IntegrationTypesConfig          IntegrationTypesConfig `json:"integration_types_config"`           // Default scopes and permissions for each supported installation context. Value for each key is an integration type configuration object
	CustomInstallUrl                string                 `json:"custom_install_url"`                 // Default custom authorization URL for the app, if enabled

	GatewayEventHandlers []GatewayEventHandler `json:"-" discord-bot:"internal"`
	gatewayConnections   []*gateway.Connection `json:"-" discord-bot:"internal"`

	HttpClient          *http.Client
	ExternalConnections sync.WaitGroup
}

//// Core Methods

// Initialize application and establish gateway connections
func (a *App) Start(config gateway.GatewayBotUrlResponse, identify *gateway.Event) error {

	// Verify there are enough remaining session starts for reccomended shards
	remaining, ok := config.SessionStartLimits["remaining"]

	if !ok {
		return fmt.Errorf("application did not receive session start limit data")
	}

	if remaining < config.Shards {
		return fmt.Errorf("not enough session starts remaining for reccomended shards")
	}

	// Introspect application details
	a.HttpClient = &http.Client{}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/applications/@me", DiscordApiBaseUrl),
		nil,
	)

	if err != nil {
		return fmt.Errorf("unable to create request for application details: %w", err)
	}

	data, err := a.Make(req)
	if err != nil {
		return fmt.Errorf("request for application details failed: %w", err)
	}

	err = json.Unmarshal(*data, a)
	if err != nil {
		return fmt.Errorf("unable to unmarshal retrieved application details: %w", err)
	}

	// Start gateway connections

	for i := range config.Shards {

		var shardNum int = i + 1

		a.Logger.Printf("Attempting to establish gateway connection %d/%d\n", shardNum, config.Shards)

		var conn gateway.Connection = gateway.Connection{
			Outgoing:   make(chan gateway.Event),
			Incoming:   make(chan gateway.Event),
			ShardIndex: i,
			BotToken:   &a.BotToken,
		}

		err := conn.Connect(&config, identify)
		if err != nil {
			return fmt.Errorf("unable to start gateway connection %d", shardNum)
		}

		a.Logger.Printf("Gateway connection %d successfully connected", shardNum)

		a.gatewayConnections = append(a.gatewayConnections, &conn)

	}

	a.ExternalConnections.Add(1)
	go a.Receive()

	return nil
}

// Makes an HTTP request to the Discord REST API
func (a *App) Make(req *http.Request) (*[]byte, error) {

	req.Header.Set("Authorization", fmt.Sprintf("Bot %v", a.BotToken)) // Set Authorization Header
	req.Header.Set("Content-Type", "application/json")                 // Set Content-Type Header

	// Make Request
	res, err := a.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Read Response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	a.Logger.Printf("Outbound HTTP request to '%s' (Status code: %d)", req.URL.Path, res.StatusCode)

	// Log error responses
	if res.StatusCode >= 400 {
		return &body, fmt.Errorf("outbound http request received not okay status code: %d", res.StatusCode)
	}

	return &body, nil
}

// Send a gateway event to one of the gateway connections
func (a *App) Send(event gateway.Event, guildId string) error {

	// Determine which shard should handle the event
	guildIdInt, err := strconv.Atoi(guildId)
	if err != nil {
		return fmt.Errorf("unable to retrieve guild id as int: %w", err)
	}

	shardNum := (guildIdInt >> 22) % len(a.gatewayConnections)

	a.Logger.Printf("Sending event to connection %d", shardNum)

	return nil
}

// Handle connections' incoming gateway events
func (a *App) Receive() {
	defer a.ExternalConnections.Done()

	for {
		for _, conn := range a.gatewayConnections {
			select {
			case event, ok := <-conn.Incoming:

				if !ok {
					// Channel is closed wait to see if its restored
					continue
				}

				eventData, err := json.Marshal(event.D)
				if err != nil {
					panic(fmt.Errorf("unable to marshal incoming payload data: %w", err))
				}

				if event.Op == 0 {
					a.Logger.Printf("Incoming event %s (\"%s\"): %s\n", OpCodes[event.Op], *event.T, eventData)
				} else {
					a.Logger.Printf("Incoming event %s: %s\n", OpCodes[event.Op], eventData)
				}

				if event.Op == 0 { // Pass dispatch events to corresponding handlers
					for _, handler := range a.GatewayEventHandlers {

						if *event.T == handler.Type {
							err := handler.Fn(&event, a) // Execute handler
							if err != nil {
								a.Logger.Printf("error occurred while executing event handler: %s", err.Error())
							}
						}

					}
				}
			default:
				// Do nothing
			}
		}
	}

}

//// Additional methods

func (a *App) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if idValue, ok := raw["id"].(string); ok {
		parsedID, err := strconv.ParseUint(idValue, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse id: %w", err)
		}
		a.Id = parsedID
	} else {
		return fmt.Errorf("payload did not include 'id' field")
	}

	return nil

}
