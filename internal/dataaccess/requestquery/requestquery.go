package requestquery

import (
	"backend/internal/model"
	"errors"

	"gorm.io/gorm"
)

const invalidID = "invalid chat id"

func Read(db *gorm.DB, id uint) (*model.RequestQuery, error) {
	var rqq model.RequestQuery
	db.First(&rqq, id)
	if db.Error != nil {
		return nil, db.Error
	}
	if rqq.ID == 0 {
		return nil, errors.New(invalidID)
	}
	return &rqq, nil
}

func Create(db *gorm.DB, rqq *model.RequestQuery) error {
	return rqq.Create(db)
}

func Update(db *gorm.DB, rqq *model.RequestQuery) error {
	return rqq.Update(db)
}

func Delete(db *gorm.DB, rqq *model.RequestQuery) error {
	return rqq.Delete(db)
}
