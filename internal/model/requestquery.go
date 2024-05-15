package model

import (
	"errors"
	"gorm.io/gorm"
)

type Status string
type Type string

const (
	Ongoing   Status = "ongoing"
	Autoreply Status = "autoreply"
	Pending   Status = "pending"
	Closed    Status = "closed"
	Reviewed  Status = "reviewed"
)

const (
	Unknown Type = "unknown"
	Query   Type = "query"
	Request Type = "request"
)

type RequestQuery struct {
	gorm.Model
	Status    Status `gorm:"type:enum('ongoing','autoreply','pending','closed','reviewed')"`
	Type      Type   `gorm:"type:enum('unknown', query','request')"`
	BookingId *uint
	Booking   *Booking
	ChatId    uint
	Chat      Chat
	Messages  []Message
}

var ErrRequestHasNilBookingId = errors.New("booking id is required for requests")

func (r *RequestQuery) Create(db *gorm.DB) error {
	return db.Create(r).Error
}

func (r *RequestQuery) Update(db *gorm.DB) error {
	return db.Updates(r).Error
}

func (r *RequestQuery) Delete(db *gorm.DB) error {
	return db.Delete(r).Error
}

func (r *RequestQuery) BeforeSave(tx *gorm.DB) (err error) {
	if r.Type == Request && r.BookingId == nil {
		return ErrRequestHasNilBookingId
	}
	return
}
