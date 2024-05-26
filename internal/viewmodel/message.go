package viewmodel

import (
	"backend/internal/model"
)

type BaseMessageView struct {
	TelegramMessageID int64     `json:"telegramMessageId"`
	By                string    `json:"by"`
	MessageBody       string    `json:"messageBody"`
	Timestamp         string    `json:"timestamp"`
	HotelStaffID      *uint     `json:"hotelStaffId,omitempty"`
	HotelStaff        *UserView `json:"hotelStaff,omitempty"`
	RequestQueryID    uint      `json:"requestQueryId,omitempty"`
}

type MessageView struct {
	BaseMessageView
	RequestQuery *BaseRequestQueryView `json:"requestQuery,omitempty"`
}

type MessageWebSocketView struct {
	BaseMessageView
	ChatID uint `json:"chatId,omitempty"`
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
