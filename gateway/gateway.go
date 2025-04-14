package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Connection struct {
	Outgoing                chan Event
	Incoming                chan Event
	ShardIndex              int
	conn                    *websocket.Conn
	connLock                sync.Mutex
	gatewayUrl              *url.URL
	sessionId               *string
	sessionResumeGatewayUrl *url.URL
	active                  bool
	lastSequence            *int
	stopHeartbeats          chan bool
	BotToken                *string
}

// External reference: https://discord.com/developers/docs/events/gateway#connecting
func (c *Connection) Connect(config *GatewayBotUrlResponse, identifyEvent *Event) error {

	// Cache WSS Url
	if gatewayUrl, err := url.Parse(config.Url); err == nil {

		urlQuery := gatewayUrl.Query()
		urlQuery.Set("v", "10")
		urlQuery.Set("encoding", "json")
		gatewayUrl.RawQuery = urlQuery.Encode()

		c.gatewayUrl = gatewayUrl

	} else {
		return fmt.Errorf("unable to format gateway url: %w", err)
	}

	// Store websocket connection object
	conn, _, err := websocket.DefaultDialer.Dial(c.gatewayUrl.String(), nil)
	if err != nil {
		return err
	}

	c.conn = conn
	c.active = true

	go c.receive() // Begin listening to events
	go c.send()    // Begin sending events

	// Receive Hello event
	hello := <-c.Incoming
	helloData, ok := hello.D.(Hello)
	if !ok {
		return fmt.Errorf("unable to access hello event data")
	}

	go c.sendHeartbeats(helloData.HeartbeatInterval) // Send heartbeats

	// Update identify payload with sharding info
	identify, ok := identifyEvent.D.(Identify)
	if !ok || identifyEvent.Op != 2 {
		return fmt.Errorf("invalid identify payload")
	}

	identify.Shard = &[2]int{
		c.ShardIndex,
		config.Shards,
	}

	identifyEvent.D = identify

	// Send Identify Payload
	c.Outgoing <- *identifyEvent

	return nil
}

// External reference:
func (c *Connection) Disconnect() error {
	// Mark connection as inactive while we attempt to restore connection
	c.active = false

	// Close gateway connection gracefully
	err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "brb"))
	if err != nil {
		return fmt.Errorf("unable to gracefully close gateway connection: %w", err)
	}
	c.conn.Close()

	return nil
}

// External reference:
func (c *Connection) Resume() error {

	// Disconnect
	if err := c.Disconnect(); err != nil {
		return fmt.Errorf("failed to resume: %w", err)
	}

	// Create new gateway connection
	conn, _, err := websocket.DefaultDialer.Dial(c.sessionResumeGatewayUrl.String(), nil)
	if err != nil {
		return fmt.Errorf("unable to connect to prior session: %w", err)
	}

	c.conn = conn
	c.active = true

	go c.receive() // Begin listening to events
	go c.send()    // Begin sending events

	// Send Resume Event
	var resume Event = Event{
		Op: 6,
		D: Resume{
			Token:     *c.BotToken,
			SessionId: *c.sessionId,
			Seq:       *c.lastSequence,
		},
	}

	c.Outgoing <- resume

	// Check for invalid session response
	firstNewEvent := <-c.Incoming

	if firstNewEvent.Op == 9 { // Received an invalid session op code

		canReconnect, ok := firstNewEvent.D.(bool)
		if !ok {
			return fmt.Errorf("invalid session event 'D' is unexpected type; expected bool")
		}

		if canReconnect {
			c.Resume()
		} else {
			c.Disconnect()
			panic(fmt.Errorf("gateway connection could not resume"))
		}

	}

	return nil
}

// Used to send fully-formed Gateway Event Objects to the gateway as a json string.
//
// Receives gateway events from the outgoing channel, marshals them, and sends them to the discord gateway.
func (c *Connection) send() {

	for msg := range c.Outgoing {

		// Convert to byte array
		payload, err := json.Marshal(msg)
		if err != nil {
			panic(fmt.Errorf("failed to Serialize payload to json: %+v", msg))
		}

		// Send full payload
		c.connLock.Lock()
		writeError := c.conn.WriteMessage(websocket.TextMessage, payload)
		c.connLock.Unlock()

		if writeError != nil {
			panic(fmt.Errorf("an error occurred while sending data to the gateway: %v", writeError))
		}

	}

}

// Used to receive json from the gateway and convert them to native go objects.
//
// Listens to the active gateway connection and unmarshals incoming payloads before passing them to the incoming channel.
func (c *Connection) receive() {

	for c.active {

		_, msg, err := c.conn.ReadMessage()

		if err != nil {

			if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
				c.Resume() // Gateway closed unexpectedly, attempt resume.
			} else {
				panic(fmt.Errorf("unable to read incoming gateway event: %w", err))
			}

		}

		if len(msg) > 0 {
			go c.processIncoming(msg)
		}

	}

}

func (c *Connection) processIncoming(message []byte) {

	var E Event
	err := json.Unmarshal(message, &E)
	if err != nil {
		panic(fmt.Errorf("unable to parse incoming gateway event: %w", err))
	}

	// Update sequence number
	if E.S != nil {
		c.lastSequence = E.S
	}

	// Handle various event necessitated app updates
	if E.Op == 0 {

		if *E.T == "READY" {

			readyData, ok := E.D.(*Ready)
			if !ok {
				panic(fmt.Errorf("ready event data could not be accessed"))
			}

			c.sessionId = &readyData.SessionId

			// Cache WSS Url
			if resumeGatewayUrl, err := url.Parse(readyData.ResumeGatewayUrl); err == nil {

				urlQuery := resumeGatewayUrl.Query()
				urlQuery.Set("v", "10")
				urlQuery.Set("encoding", "json")
				resumeGatewayUrl.RawQuery = urlQuery.Encode()

				c.sessionResumeGatewayUrl = resumeGatewayUrl

			} else {
				panic(fmt.Errorf("unable to parse session resume gateway url: %w", err))
			}

		}

	}

	// Forward for additional processing
	c.Incoming <- E
}

// Handles the sending of heartbeat events at the specified interval
func (c *Connection) sendHeartbeats(interval uint) {

	c.stopHeartbeats = make(chan bool) // Create channel to receive notifications to stop.

	for {
		select {

		case <-c.stopHeartbeats:
			return // Stop sending heartbeats

		default:
			// Process heartbeat
			intervalDur, err := time.ParseDuration(
				fmt.Sprintf("%dms", interval),
			)
			if err != nil {
				panic(err)
			}

			// send heartbeat here
			var Pulse Event = Event{
				Op: 1,
				D:  c.lastSequence,
			}

			c.Outgoing <- Pulse

			time.Sleep(intervalDur)

		}

	}

}

////

// External reference: https://discord.com/developers/docs/events/gateway#get-gateway-bot-json-response
type GatewayBotUrlResponse struct {
	Url                string         `json:"url"`
	Shards             int            `json:"shards"`
	SessionStartLimits map[string]int `json:"session_start_limit"`
}

// Helper function: requests the wss url for the discord gateway.
func GetGatewayBot(botToken string) GatewayBotUrlResponse {

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

////

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
