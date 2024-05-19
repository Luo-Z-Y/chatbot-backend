package message

import (
	"backend/internal/model"
	"errors"

	"gorm.io/gorm"
)

const invalidID = "invalid chat id"

func Read(db *gorm.DB, id uint) (*model.Message, error) {
	var msg model.Message
	db.First(&msg, id)
	if db.Error != nil {
		return nil, db.Error
	}
	if msg.ID == 0 {
		return nil, errors.New(invalidID)
	}
	return &msg, nil
}

func Create(db *gorm.DB, msg *model.Message) error {
	return msg.Create(db)
}

func Update(db *gorm.DB, msg *model.Message) error {
	return msg.Update(db)
}

func Delete(db *gorm.DB, msg *model.Message) error {
	return msg.Delete(db)
}
