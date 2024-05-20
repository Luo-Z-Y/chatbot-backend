package api

const (
	AuthType = "auth"
)

type WebSocketMessage struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
