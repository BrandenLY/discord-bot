package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"brandenly.com/go/packages/discord-bot/common"
	"brandenly.com/go/packages/discord-bot/gateway"
	"brandenly.com/go/packages/discord-bot/utils"
)

var lastBeat time.Time = time.Now()

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
	Id                uint64 `json:"-" discord-bot:"internal"` // ID of the app
	PublicKey         string `json:"-" discord-bot:"internal"` // App Public Key
	BotToken          string `json:"-" discord-bot:"internal"` // App Bot Token
	DiscordApiBaseUrl string `json:"-" discord-bot:"internal"` // Discord base url; the base url used for requests to discord's REST api.
	Color             int    `json:"-" discord-bot:"internal"` // Color

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

	send_heartbeat      bool
	send_heartbeat_lock sync.RWMutex
	last_sequence       *int

	GatewayEventHandlers []utils.EventHandler
	registeredCommands   []RegisteredCommand
	// InteractionWorkerChannels map[string]chan *gateway.InteractionCreate
	HttpClient          *http.Client
	ExternalConnections sync.WaitGroup
}

// Establishes the handshake with the gateway and initiates heartbeat signaling.
// func (a *App) ConnectGateway() {

// 	// Before your app can establish a connection to the Gateway, it should
// 	// call the `Get Gateway` or the `Get Gateway Bot` endpoint. Either endpoint
// 	// will return a payload with a url field whose value is the WSS URL you
// 	// can use to open a WebSocket connection. In addition to the URL, Get
// 	// Gateway Bot contains additional information about the recommended
// 	// number of shards and the session start limits for your app.

// 	// https://discord.com/developers/docs/events/gateway#connecting

// 	// ESTABLISH GATEWAY CONNECTION
// 	// Format Gateway Url
// 	gatewayUrl, _ := url.Parse(GatewayUrlResponse.Url)
// 	urlQueryStr := gatewayUrl.Query()
// 	urlQueryStr.Set("v", "10")
// 	urlQueryStr.Set("encoding", "json")
// 	gatewayUrl.RawQuery = urlQueryStr.Encode()

// 	// Connect using the Gorilla WebSocket package
// 	conn, _, err := websocket.DefaultDialer.Dial(gatewayUrl.String(), nil)
// 	if err != nil {
// 		log.Fatal("Failed to connect to Discord WebSocket gateway:", err)
// 	}

// 	a.ws = conn // Save websocket connection

// 	// Receive Hello event
// 	_, msg, err := a.ws.ReadMessage()
// 	if err != nil {
// 		log.Fatal("Unable to detect hello event")
// 	}

// 	log.Printf("===INCOMING===\n%v\n\n", string(msg))
// 	lastBeat = time.Now()

// 	// Parse Hello event
// 	var HelloEvent gateway.Event
// 	helloEventDecodeErr := json.Unmarshal(msg, &HelloEvent)
// 	if helloEventDecodeErr != nil {
// 		log.Fatal("Unable to parse discord hello event", helloEventDecodeErr)
// 	}

// 	if HelloData, ok := HelloEvent.D.(gateway.Hello); ok {
// 		log.Printf("Successfully connected to gateway.\n")

// 		a.ExternalConnections.Add(1)
// 		go a.Receive() // Listen for responses

// 		a.ExternalConnections.Add(1)
// 		go a.keepAlive(&HelloData.HeartbeatInterval) // Begin Heartbeat

// 		a.ExternalConnections.Add(1)
// 		go a.Send() // Send payloads

// 		// Configure Identify Event
// 		e := gateway.Event{
// 			Op: 2,
// 			D: gateway.Identify{
// 				Token: a.BotToken,
// 				Properties: gateway.IdentifyConnProperties{
// 					Os:      "windows",
// 					Browser: "cockorman-inspector",
// 					Device:  "cockorman-inspector",
// 				},
// 				Presence: common.Presence{
// 					Activities: []common.Activity{
// 						{
// 							Name: "Gooning",
// 							Type: 4,
// 						},
// 					},
// 					Status: "online",
// 					Afk:    false,
// 				},
// 				Intents: utils.FormIntents(
// 					BotIntents["GUILDS"],
// 					BotIntents["GUILD_BANS"],
// 					BotIntents["GUILD_VOICE_STATES"],
// 					BotIntents["GUILD_MESSAGES"],
// 					BotIntents["GUILD_MESSAGE_REACTIONS"],
// 					BotIntents["GUILD_MESSAGE_TYPING"],
// 					BotIntents["MODERATE_MEMBERS"],
// 					// PRIVILEGED INTENTS
// 					// BotIntents["GUILD_PRESENCES"],
// 					// BotIntents["GUILD_MEMBERS"],
// 					// BotIntents["MESSAGE_CONTENT"],
// 				),
// 			},
// 		}

// 		// Send Identify Event
// 		a.outgoing <- &e
// 	}

// }

// Makes an HTTP request to the Discord REST API
func (a *App) Make(req *http.Request) (*[]byte, error) {

	req.Header.Set("Authorization", fmt.Sprintf("Bot %v", a.BotToken)) // Set Authorization Header
	req.Header.Set("Content-Type", "application/json")

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

	// Log error responses
	if res.StatusCode >= 400 {
		log.Printf("===FAILED API REQUEST=== %s\n%s\n\n", res.Status, &body)
	}

	return &body, nil
}

// Convert Gateway Event Response to structs
func parseGatewayResponse(payload []byte) (gateway.Event, error) {

	var event gateway.Event
	err := json.Unmarshal(payload, &event)

	return event, err

}

type RegisteredCommand struct {
	Id                       *uint64                           `json:"id,omitempty"`               // Unique ID of command
	Type                     *uint8                            `json:"type,omitempty"`             // Type of command, defaults to 1
	GuildId                  *uint64                           `json:"guild_id,omitempty"`         // Guild ID of the command, if not global
	Name                     string                            `json:"name"`                       // Name of command, 1-32 characters
	Description              string                            `json:"description"`                // Localization dictionary for name field. Values follow the same restrictions as name
	DefaultMemberPermissions *string                           `json:"default_member_permissions"` // Set of permissions represented as a bit set
	NSFW                     *bool                             `json:"nsfw,omitempty"`             // Indicates whether the command is age-restricted, defaults to false
	Version                  uint64                            `json:"version"`
	Fn                       *func(*gateway.Event, *App) error `json:"-"` // Callback function meant to handle interaction events for this command
}

type UnexpectedPayloadFormat struct {
	Msg string
}

func (upf UnexpectedPayloadFormat) Error() string {
	return upf.Msg
}

func (rc *RegisteredCommand) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if idValue, ok := raw["id"].(string); ok {
		parsedID, err := strconv.ParseUint(idValue, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse id: %w", err)
		}
		rc.Id = &parsedID
	} else {
		return UnexpectedPayloadFormat{Msg: "Unexpected type"}
	}

	return nil

}

func (a *App) RegisterCommand(cmd any, handler func(*gateway.Event, *App) error) error {

	// Serialize payload
	cmdPayload, err := json.Marshal(&cmd)
	if err != nil {
		return err
	}

	// Create Request
	request, err := http.NewRequest("POST", fmt.Sprintf("https://discord.com/api/applications/%v/commands", a.Id), bytes.NewReader(cmdPayload))
	if err != nil {
		return err
	}

	// Execute Request
	responseBody, err := a.Make(request)
	if err != nil {
		return err
	}

	// Register Command Handler fn with app
	var NewlyRegisteredCommand RegisteredCommand
	NewlyRegisteredCommand.Fn = &handler                                // Add Callback Fn
	decodeErr := json.Unmarshal(*responseBody, &NewlyRegisteredCommand) // Add remaining details returned by api
	if decodeErr != nil {
		return decodeErr
	}

	a.registeredCommands = append(a.registeredCommands, NewlyRegisteredCommand)

	return nil
}
