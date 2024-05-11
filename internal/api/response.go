package api

type Message struct {
	Message string `json:"message"`
}

type Response struct {
	Data     interface{} `json:"data,omitempty"`
	Messages []Message   `json:"messages,omitempty"`
}
