package viewmodel

type BaseMessageView struct {
	TelegramMessageId int64     `json:"telegram_message_id"`
	By                string    `json:"by"`
	MessageBody       string    `json:"message_body"`
	Timestamp         string    `json:"timestamp"`
	HotelStaffId      *uint     `json:"hotel_staff_id,omitempty"`
	HotelStaff        *UserView `json:"hotel_staff,omitempty"`
	RequestQueryId    uint      `json:"request_query_id,omitempty"`
}

type MessageView struct {
	BaseMessageView
	RequestQuery BaseRequestQueryView `json:"request_query"`
}

type MessageWebSocketView struct {
	BaseMessageView
	ChatID uint `json:"chat_id,omitempty"`
}
