package message

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

const invalidID = "invalid chat id"

func Read(db *gorm.DB, id uint) (*model.Message, error) {
	var msg model.Message
	result := db.Model(&model.Message{}).
		Where("id = ?", id).
		First(&msg, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &msg, nil
}

func Create(db *gorm.DB, msg *model.Message) error {
	return msg.Create(db)
}

func Update(db *gorm.DB, msg *model.Message) error {
	if err := ensureRequestQueryExists(db, msg.RequestQueryId); err != nil {
		return err
	}

	return msg.Update(db)
}

func Delete(db *gorm.DB, msg *model.Message) error {
	return msg.Delete(db)
}

func ensureRequestQueryExists(db *gorm.DB, id uint) error {
	var rqq model.RequestQuery
	result := db.Model(&model.RequestQuery{}).
		Where("id = ?", id).
		First(&rqq, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
