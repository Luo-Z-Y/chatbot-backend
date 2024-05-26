package viewmodel

import (
	"backend/internal/model"
)

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
