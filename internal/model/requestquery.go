package model

import (
	"backend/pkg/error/externalerror"
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
	BookingID *uint
	Booking   *Booking
	ChatID    uint
	Chat      Chat

	Messages []Message `gorm:"->;<-"`
}

var ErrRequestHasNilBookingId = externalerror.AuthRequiredError{}
var ErrBookingIdDoesNotExist = errors.New("booking id does not exist")

func (r *RequestQuery) Create(db *gorm.DB) error {
	return db.Create(r).Error
}

func (r *RequestQuery) Update(db *gorm.DB) error {
	return db.Updates(r).Error
}

func (r *RequestQuery) Delete(db *gorm.DB) error {
	return db.Delete(r).Error
}

func (r *RequestQuery) BeforeSave(tx *gorm.DB) error {
	if r.Type == TypeRequest && r.BookingID == nil {
		return ErrRequestHasNilBookingId
	}
	if r.BookingID != nil {
		var booking Booking
		tx.First(&booking, *r.BookingID)
		if tx.Error != nil || booking.ID == 0 {
			return ErrBookingIdDoesNotExist
		}
	}
	return nil
}
