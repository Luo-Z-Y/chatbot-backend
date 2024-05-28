package requestquery

import (
	"backend/internal/dataaccess/message"
	"backend/internal/model"

	"gorm.io/gorm"
)

func Read(db *gorm.DB, id uint) (*model.RequestQuery, error) {
	var rqq model.RequestQuery
	result := db.Model(&model.RequestQuery{}).
		Where("id = ?", id).
		First(&rqq, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &rqq, nil
}

func ReadLatestByChatID(db *gorm.DB, chatID uint) (*model.RequestQuery, error) {
	var rqq model.RequestQuery
	result := db.Model(&model.RequestQuery{}).
		Where("chat_id = ?", chatID).
		Order("created_at desc").
		First(&rqq)
	if result.Error != nil {
		return nil, result.Error
	}

	return &rqq, nil
}

func Create(db *gorm.DB, rqq *model.RequestQuery) error {
	if err := ensureChatExists(db, rqq.ChatID); err != nil {
		return err
	}

	return rqq.Create(db)
}

func Update(db *gorm.DB, rqq *model.RequestQuery) error {
	if err := ensureChatExists(db, rqq.ChatID); err != nil {
		return err
	}

	return rqq.Update(db)
}

func Delete(db *gorm.DB, rqq *model.RequestQuery) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := pruneAssociations(tx, rqq); err != nil {
			return err
		}

		return rqq.Delete(tx)
	})

	return err
}

func ensureChatExists(db *gorm.DB, id uint) error {
	var chat model.Chat
	result := db.Model(&model.Chat{}).
		Where("id = ?", id).
		First(&chat)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func pruneAssociations(db *gorm.DB, rqq *model.RequestQuery) error {
	var msgs []model.Message

	result := db.Model(&model.Message{}).
		Where("request_query_id = ?", rqq.ID).
		Find(&msgs)
	if result.Error != nil {
		return result.Error
	}

	for _, msg := range msgs {
		if err := message.Delete(db, &msg); err != nil {
			return err
		}
	}

	return nil
}
