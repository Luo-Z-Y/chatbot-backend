package model

import (
	"gorm.io/gorm"
)

// Booking represents our "authentication" data.
// Our "authentication" process should be just adding to this table
// after trivially assuming they are all correct.
type Booking struct {
	gorm.Model
	RoomNumber string
	LastName   string
	ChatId     uint
	Chat       Chat
}

func (b *Booking) Create(db *gorm.DB) error {
	return db.Create(b).Error
}

func (b *Booking) Update(db *gorm.DB) error {
	return db.Updates(b).Error
}

func (b *Booking) Delete(db *gorm.DB) error {
	return db.Delete(b).Error
}
