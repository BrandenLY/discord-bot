package utils

import "brandenly.com/go/packages/discord-bot/gateway"

type EventHandler struct {
	Type string
	Fn   func(*gateway.Event) error
}
