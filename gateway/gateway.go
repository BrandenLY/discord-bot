package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Connection struct {
	conn                    *websocket.Conn
	ConnLock                sync.Mutex
	Outgoing                chan Event
	Incoming                chan Event
	sessionId               *string
	sessionsResetAfter      time.Time
	remainingSessionStarts  int
	sessionResumeGatewayUrl *string
	active                  bool
	recommendedShards       int
}

// External reference: https://discord.com/developers/docs/events/gateway#connecting
func (c *Connection) Connect(botToken string) {

	if c.remainingSessionStarts < 1 {
		panic(fmt.Errorf("this bot cannot start any new gateway events until %v", c.sessionsResetAfter.Format("Jan _2 2006, 3:04PM")))
	}

}

// Used to send fully-formed Gateway Event Objects to the gateway as a json string.
func (c *Connection) send() {

	for msg := range c.Outgoing {

		// Convert to byte array
		payload, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Failed to Serialize payload to json: %+v", msg)
		}

		log.Printf("\033[36;1m===OUTGOING===\033[0m\n%v\n\n", string(payload))

		// Send full payload
		c.ConnLock.Lock()
		writeError := c.conn.WriteMessage(websocket.TextMessage, payload)
		c.ConnLock.Unlock()

		if writeError != nil {
			log.Printf("an error occurred while sending data to the gateway%v\n", writeError)
		}

	}

}

// Used to receive json from the gateway and convert them to native go objects.
func (c *Connection) receive() {

}

////

// External reference: https://discord.com/developers/docs/events/gateway#get-gateway-bot-json-response
type GatewayBotUrlResponse struct {
	Url                string         `json:"url"`
	Shards             int            `json:"shards"`
	SessionStartLimits map[string]int `json:"session_start_limit"`
}

// Helper function: requests the wss url for the discord gateway.
func getGatewayBot(botToken string) GatewayBotUrlResponse {

	var retrieveFrom string = "https://discord.com/api/gateway/bot"
	var client http.Client = http.Client{
		Timeout: 10 * time.Second,
	}

	// Format Request
	request, err := http.NewRequest("GET", retrieveFrom, nil)
	if err != nil {
		panic(fmt.Errorf("failed to initialize http request to retrieve websocket url: %w", err))
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bot %s", botToken))

	// Make Request
	res, err := client.Do(request)
	if err != nil {
		panic(fmt.Errorf("failed to make http request to retrieve websocket url: %w", err))
	}

	// Read Response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(fmt.Errorf("failed to parse http response body for websocket url: %w", err))
	}
	defer res.Body.Close()

	// Parse Response
	var payload GatewayBotUrlResponse
	err = json.Unmarshal(body, &payload)
	if err != nil {
		panic(fmt.Errorf("failed to parse websocket url response body json: %w", err))
	}

	return payload
}

// External reference: https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-close-event-codes
var closeIsResumable map[int]bool = map[int]bool{
	4000: true,  // Unknown error
	4001: true,  // Unknown opcode
	4002: true,  // Decode error
	4003: true,  // Not authenticated
	4004: false, // Authentication failed
	4005: true,  // Already authenticated
	4007: true,  // Invalid seq
	4008: true,  // Rate limited
	4009: true,  // Session timed out
	4010: false, // Invalid shard
	4011: false, // Sharding required
	4012: false, // Invalid API version
	4013: false, // Invalid intent(s)
	4014: false, // Disallowed intent(s)
}
