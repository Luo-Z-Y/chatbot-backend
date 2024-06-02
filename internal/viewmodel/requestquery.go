package viewmodel

import (
	"backend/internal/model"
)

type BaseRequestQueryView struct {
	ID        uint   `json:"id"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	BookingID *uint  `json:"bookingId,omitempty"`
	ChatID    uint   `json:"chatId"`
}

type RequestQueryView struct {
	BaseRequestQueryView
	Messages []BaseMessageView `json:"messages"`
}

func RequestQueryViewFrom(rq *model.RequestQuery) RequestQueryView {
	messageViews := make([]BaseMessageView, len(rq.Messages))
	for i, msg := range rq.Messages {
		messageViews[i] = BaseMessageViewFrom(&msg)
	}
	return RequestQueryView{
		BaseRequestQueryView: BaseRequestQueryView{
			ID:        rq.ID,
			Status:    string(rq.Status),
			Type:      string(rq.Type),
			BookingID: rq.BookingID,
			ChatID:    rq.ChatID,
		},
		Messages: messageViews,
	}
}
