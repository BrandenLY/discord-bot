package discord

import "brandenly.com/go/packages/discord-bot/gateway"

// Called when an event matching the handlers type is received.
type GatewayEventHandler struct {
	Type string
	Fn   func(*gateway.Event, *App) error
}
