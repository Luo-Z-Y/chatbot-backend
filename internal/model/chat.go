package model

import (
	"gorm.io/gorm"
)

// Chat represents the telegram chat and has information required to send messages to it
// For simplicity, assume one booking to one chat for now. We can change this in future iterations if time permits.
type Chat struct {
	gorm.Model
	TelegramChatId int64 `gorm:"unique"`

	Booking        *Booking       `gorm:"->"`
	RequestQueries []RequestQuery `gorm:"->;<-"`
}

func (c *Chat) Create(db *gorm.DB) error {
	return db.Create(c).Error
}

func (c *Chat) Update(db *gorm.DB) error {
	return db.Updates(c).Error
}

func (c *Chat) Delete(db *gorm.DB) error {
	return db.Delete(c).Error
}
