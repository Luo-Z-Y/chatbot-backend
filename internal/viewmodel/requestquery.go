package viewmodel

type BaseRequestQueryView struct {
	ID        uint   `json:"id"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	BookingID *uint  `json:"booking_id,omitempty"`
	ChatID    uint   `json:"chat_id"`
}

type RequestQueryView struct {
	BaseRequestQueryView
	Messages []BaseMessageView `json:"messages"`
}
