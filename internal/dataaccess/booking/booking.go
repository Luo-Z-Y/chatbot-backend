package booking

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

func Read(db *gorm.DB, id uint) (*model.Booking, error) {
	var booking model.Booking
	result := db.Model(&model.Booking{}).
		Where("id = ?", id).
		First(&booking, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &booking, nil
}

func ReadByChatID(db *gorm.DB, chatID uint) (*model.Booking, error) {
	var booking model.Booking
	result := db.Model(&model.Booking{}).
		Where("chat_id = ?", chatID).
		First(&booking)
	if result.Error != nil {
		return nil, result.Error
	}

	return &booking, nil
}

func Create(db *gorm.DB, booking *model.Booking) error {
	return booking.Create(db)
}

func Update(db *gorm.DB, booking *model.Booking) error {
	return booking.Update(db)
}

func Delete(db *gorm.DB, booking *model.Booking) error {
	return booking.Delete(db)
}
