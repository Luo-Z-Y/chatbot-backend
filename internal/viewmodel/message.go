package viewmodel

import (
	"backend/internal/model"
)

type BaseMessageView struct {
	TelegramMessageID int64     `json:"telegram_message_id"`
	By                string    `json:"by"`
	MessageBody       string    `json:"message_body"`
	Timestamp         string    `json:"timestamp"`
	HotelStaffID      *uint     `json:"hotel_staff_id,omitempty"`
	HotelStaff        *UserView `json:"hotel_staff,omitempty"`
	RequestQueryID    uint      `json:"request_query_id,omitempty"`
}

type MessageView struct {
	BaseMessageView
	RequestQuery *BaseRequestQueryView `json:"request_query,omitempty"`
}

type MessageWebSocketView struct {
	BaseMessageView
	ChatID uint `json:"chat_id,omitempty"`
}

func BaseMessageViewFrom(message *model.Message) BaseMessageView {

	var hotelStaff *UserView
	if message.HotelStaff != nil {
		hotelStaff = UserViewFrom(message.HotelStaff)
	} else {
		hotelStaff = nil
	}
	return BaseMessageView{
		TelegramMessageID: message.TelegramMessageID,
		By:                string(message.By),
		MessageBody:       message.MessageBody,
		Timestamp:         message.Timestamp.Format("2006-01-02 15:04:05"),
		HotelStaffID:      message.HotelStaffID,
		HotelStaff:        hotelStaff,
		RequestQueryID:    message.RequestQueryID,
	}
}

func MessageViewFrom(message *model.Message) MessageView {

	return MessageView{
		BaseMessageView: BaseMessageViewFrom(message),
		RequestQuery:    nil,
	}
}
