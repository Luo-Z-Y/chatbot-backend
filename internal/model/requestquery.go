package model

import (
	"errors"
	"gorm.io/gorm"
)

type Status string
type Type string

const (
	StatusOngoing   Status = "ongoing"
	StatusAutoreply Status = "autoreply"
	StatusPending   Status = "pending"
	StatusClosed    Status = "closed"
	StatusReviewed  Status = "reviewed"
)

const (
	TypeUnknown Type = "unknown"
	TypeQuery   Type = "query"
	TypeRequest Type = "request"
)

type RequestQuery struct {
	gorm.Model
	Status    Status
	Type      Type
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
	if r.Type == TypeRequest && r.BookingId == nil {
		return ErrRequestHasNilBookingId
	}
	return
}
