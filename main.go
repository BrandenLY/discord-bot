package main

import "brandenly.com/go/packages/discord-bot/gateway"

var BotToken string = "MTM0MDY5Mzk4OTg0OTk2MDU2MA.Gvs6Kl.GCW_nQvwLAxny9Tkvll-Iqe1fNxObjbP85IFrk"

func main() {

	// Testing currently

	// Test gateway connection
	var Gateway gateway.Connection

	Gateway.Connect(BotToken)
}
