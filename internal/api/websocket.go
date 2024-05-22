package api

const (
	AuthType            = "auth"
	MessageReceivedType = "message_received"
)

type WebSocketMessage struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
