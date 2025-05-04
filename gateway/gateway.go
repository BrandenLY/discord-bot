package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Connection struct {
	ShardIndex int
	Outgoing   chan Event
	Incoming   chan Event
	BotToken   *string
	Logger     *log.Logger
	gatewayUrl *url.URL
	conn       *websocket.Conn
	ctx        context.Context
	cancel     context.CancelFunc
	active     bool

	heartbeatInterval uint
	session           session
	parentCtx         context.Context
}

type session struct {
	Id               *string
	ResumeGatewayUrl *url.URL
	LastSequence     *int
	Wg               sync.WaitGroup
	InvalidSession   chan any
	SuccessfulResume chan any
	NewSession       chan any
}

// External reference: https://discord.com/developers/docs/events/gateway#connecting
func (c *Connection) Connect(ctx context.Context, wg *sync.WaitGroup, config *GatewayBotUrlResponse, identifyEvent *Event) error {

	// Inform external api's that the connection is fully closed.
	defer wg.Done()

	// Prevent consecutive calls to Connection.Connect()
	if c.active {
		return fmt.Errorf("cannot call Connect() on an already active connection instance") // Connection is already active
	}

	// Configure logger, this can be set externally so it is only created if not provided
	if c.Logger == nil {
		c.Logger = log.New(os.Stdout, fmt.Sprintf("[gateway-connection:%d]", c.ShardIndex+1), log.LstdFlags|log.Lshortfile)
	}

	// Cache WSS Url
	gatewayUrl, err := url.Parse(config.Url)
	if err != nil {
		return fmt.Errorf("unable to format gateway url: %w", err)
	}

	urlQuery := gatewayUrl.Query()
	urlQuery.Set("v", "10")
	urlQuery.Set("encoding", "json")
	gatewayUrl.RawQuery = urlQuery.Encode()

	c.gatewayUrl = gatewayUrl

	// Create invalid session channel
	c.session.InvalidSession = make(chan any)

	// Create invalid session channel
	c.session.NewSession = make(chan any)

	// Create successful resume channel
	c.session.SuccessfulResume = make(chan any)

	// Cache outer context
	c.parentCtx = ctx

	for {
		select {
		case <-ctx.Done():
			err := c.Disconnect()
			if err != nil {
				return fmt.Errorf("unable to close connection gracefully: %w", err)
			}
			return nil
		default:

			// Reset wait group
			c.session.Wg = sync.WaitGroup{}

			// Reset context
			c.ctx, c.cancel = context.WithCancel(c.parentCtx)

			// Store websocket connection object
			conn, _, err := websocket.DefaultDialer.Dial(c.gatewayUrl.String(), nil)
			if err != nil {
				return err
			}

			c.conn = conn
			c.active = true

			c.session.Wg.Add(3) // send(), receive(), sendHeartbeats()

			go c.send()    // Begin sending events
			go c.receive() // Begin listening to events

			// Receive Hello event & Identify
			select {
			case hello := <-c.Incoming:

				helloData, ok := hello.D.(Hello)
				if !ok {
					return fmt.Errorf("unable to access hello event data")
				}

				c.heartbeatInterval = helloData.HeartbeatInterval
				go c.sendHeartbeats(c.heartbeatInterval) // Send heartbeats

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
			case <-time.After(10 * time.Second):
				return fmt.Errorf("timed out waiting for hello event")
			}

			// Wait for active session to close
			<-c.session.NewSession
			c.Logger.Printf("Attempting full reconnection")
		}
	}

}

// External reference: https://discord.com/developers/docs/events/gateway#disconnecting
func (c *Connection) Disconnect() error {

	c.Logger.Printf("Disconnecting gateway %d\n", c.ShardIndex+1)

	c.active = false // Mark connection as inactive
	c.cancel()       // Halt goroutines

	// Close gateway connection gracefully
	err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "brb"))
	if err != nil {
		// Do nothing
	}

	err = c.conn.Close()
	if err != nil {
		// Do nothing
	}

	// Ensure all goroutines exited
	c.session.Wg.Wait()

	return nil
}

// External reference: https://discord.com/developers/docs/events/gateway#resuming
func (c *Connection) Resume() error {

	c.Logger.Printf("Attempting to resume gateway connection %d\n", c.ShardIndex+1)

	// Disconnect
	if err := c.Disconnect(); err != nil {
		return fmt.Errorf("disconnect failed: %w", err)
	}

	// Create new gateway connection
	conn, _, err := websocket.DefaultDialer.Dial(c.session.ResumeGatewayUrl.String(), nil)
	if err != nil {
		return fmt.Errorf("unable to connect to prior session: %w", err)
	}

	c.conn = conn
	c.active = true
	c.session.Wg = sync.WaitGroup{}
	c.ctx, c.cancel = context.WithCancel(c.parentCtx)

	c.session.Wg.Add(3)
	go c.receive()                           // Begin listening to events
	go c.send()                              // Begin sending events
	go c.sendHeartbeats(c.heartbeatInterval) // Begin sending heartbeats

	// Send Resume Event
	var resume Event = Event{
		Op: 6,
		D: Resume{
			Token:     *c.BotToken,
			SessionId: *c.session.Id,
			Seq:       *c.session.LastSequence,
		},
	}
	c.Outgoing <- resume

	// Wait for resume result update
	select {
	case <-time.After(1 * time.Minute):
		return fmt.Errorf("resume function never received status update")

	case msg := <-c.session.InvalidSession:

		// Parse event
		E, ok := msg.(Event)
		if !ok {
			return fmt.Errorf("resume received unexpected channel message type")
		}

		// Parse event data
		canResume, ok := E.D.(bool)
		if !ok {
			return fmt.Errorf("resume received unexpected channel message event data type")
		}

		// Check resumability
		if canResume {
			c.Resume()
			return nil
		} else {
			c.Disconnect()
			c.session.NewSession <- E
			return nil
		}

	case <-c.session.SuccessfulResume:
		c.Logger.Printf("Successfully resumed gateway connection")
		return nil
	}

}

// Receives gateway events from the outgoing channel, marshals them, and sends them to the discord gateway.
func (c *Connection) send() {

	c.Logger.Printf("Gateway connection %d started sending (GID:%d)\n", c.ShardIndex+1, GetGID())
	defer c.Logger.Printf("Gateway connection %d stopped sending (GID:%d)\n", c.ShardIndex+1, GetGID())
	defer c.session.Wg.Done()

	for {
		select {

		case <-c.ctx.Done():
			return // Stop sending

		case msg, ok := <-c.Outgoing:

			if !ok {
				return // Stop sending
			}

			// Convert to byte array
			payload, err := json.Marshal(msg)
			if err != nil {
				panic(fmt.Errorf("failed to Serialize payload to json: %+v", msg))
			}

			// Send full payload
			writeError := c.conn.WriteMessage(websocket.TextMessage, payload)
			if writeError != nil {
				panic(fmt.Errorf("an error occurred while sending data to the gateway: %v", writeError))
			}
		}
	}

}

// Listens to the active gateway connection and unmarshals incoming payloads before passing them to the incoming channel.
func (c *Connection) receive() {

	c.Logger.Printf("Gateway connection %d started receiving (GID:%d)", c.ShardIndex+1, GetGID())
	defer c.Logger.Printf("Gateway connection %d stopped receiving (GID:%d)", c.ShardIndex+1, GetGID())
	defer c.session.Wg.Done()

	for {
		select {

		case <-c.ctx.Done():

			return // Stop receiving

		default:

			_, msg, err := c.conn.ReadMessage()

			if err != nil {

				if websocketCloseErr, ok := err.(*websocket.CloseError); ok {

					// Handle clean close

					c.Logger.Printf("Connection closed with close code: %+v", websocketCloseErr)

				} else {

					// Handle abrupt close

					c.Logger.Printf("Connection closed without close code: %+v", err)

				}

				if c.active { // Attempt to resume

					go func() {
						err := c.Resume()
						if err != nil {
							c.Logger.Printf("Failed to resume after websocket closure: %s", err.Error())
						}
					}()

				}

				return

			}

			// Process incoming message
			if len(msg) > 0 {
				go c.processIncoming(msg)
			}

		}
	}

}

func (c *Connection) processIncoming(message []byte) {

	var E Event
	err := json.Unmarshal(message, &E)
	if err != nil {
		c.Logger.Printf("FAILED TO PARSE INCOMING GATEWAY EVENT: %s", err.Error())
		return
	}

	// Update sequence number
	if E.S != nil {
		c.session.LastSequence = E.S
	}

	// Handle successful resume
	if E.Op == 9 {
		c.session.InvalidSession <- E
	}

	// Handle reconnect events
	if E.Op == 7 {
		go func() {
			err := c.Resume()
			if err != nil {
				c.Logger.Printf("Failed to resume after websocket closure: %s", err.Error())
			}
		}()
		return
	}

	// Handle various dispatch event state updates
	if E.Op == 0 {

		if *E.T == "READY" {

			readyData, ok := E.D.(*Ready)
			if !ok {
				panic(fmt.Errorf("ready event data could not be accessed"))
			}

			c.session.Id = &readyData.SessionId

			// Cache WSS Url
			if resumeGatewayUrl, err := url.Parse(readyData.ResumeGatewayUrl); err == nil {

				urlQuery := resumeGatewayUrl.Query()
				urlQuery.Set("v", "10")
				urlQuery.Set("encoding", "json")
				resumeGatewayUrl.RawQuery = urlQuery.Encode()

				c.session.ResumeGatewayUrl = resumeGatewayUrl

			} else {
				panic(fmt.Errorf("unable to parse session resume gateway url: %w", err))
			}

		}

		if *E.T == "RESUMED" {
			c.session.SuccessfulResume <- E
		}

	}

	// Forward for additional processing
	c.Incoming <- E
}

// Handles the sending of heartbeat events at the specified interval
func (c *Connection) sendHeartbeats(interval uint) {

	defer c.session.Wg.Done()

	intervalDur, err := time.ParseDuration(
		fmt.Sprintf("%dms", interval),
	)
	if err != nil {
		panic(err)
	}
	var t time.Ticker = *time.NewTicker(intervalDur)

	for {
		select {

		case <-c.ctx.Done():
			return // Stop sending heartbeats

		case <-t.C:
			// Process heartbeat

			// send heartbeat here
			var Pulse Event = Event{
				Op: 1,
				D:  c.session.LastSequence,
			}

			c.Outgoing <- Pulse

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

// GetGID returns the current goroutine's ID (for debugging purposes only)
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		return 0
	}
	gid, err := strconv.ParseUint(string(b[:i]), 10, 64)
	if err != nil {
		return 0
	}
	return gid
}
