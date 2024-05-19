package viewmodel

type TgAuthView struct {
	ChatID      uint   `json:"chat_id"`
	Credentials string `json:"credentials"`
}
