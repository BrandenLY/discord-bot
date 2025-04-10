package discord

var OpCodes map[int]string = map[int]string{
	0:  "Dispatch",
	1:  "Heartbeat",
	2:  "Identify",
	3:  "Presence Update",
	4:  "Voice State Update",
	6:  "Resume",
	7:  "Reconnect",
	8:  "Request Guild Members",
	9:  "Invalid Session",
	10: "Hello",
	11: "Heartbeat ACK",
	31: "Request Soundboard Sounds",
}
