package chat

import (
	"backend/internal/model"
	"errors"

	"gorm.io/gorm"
)

const invalidID = "invalid chat id"

func Read(db *gorm.DB, id uint) (*model.Chat, error) {
	var chat model.Chat
	db.First(&chat, id)
	if db.Error != nil {
		return nil, db.Error
	}
	if chat.ID == 0 {
		return nil, errors.New(invalidID)
	}
	return &chat, nil
}

func Create(db *gorm.DB, chat *model.Chat) error {
	return chat.Create(db)
}

func Update(db *gorm.DB, chat *model.Chat) error {
	return chat.Update(db)
}

func Delete(db *gorm.DB, chat *model.Chat) error {
	return chat.Delete(db)
}
