package bookingparams

import "backend/internal/model"

type Params struct {
	ChatID     uint   `json:"chat_id" validate:"required"`
	RoomNumber string `json:"room_number" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
}

func (p *Params) ToModel() *model.Booking {
	return &model.Booking{
		RoomNumber: p.RoomNumber,
		LastName:   p.LastName,
		ChatId:     p.ChatID,
	}
}
