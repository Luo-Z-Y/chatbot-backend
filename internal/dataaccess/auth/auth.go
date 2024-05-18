package auth

import (
	"backend/internal/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	wrongUsernameOrPassword = "username or password is incorrect"
	invalidID               = "invalid user id"
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckLogin(db *gorm.DB, username string, password string) (*model.User, error) {
	var user model.User
	db.First(&user, "username = ?", username)
	if db.Error != nil {
		return nil, db.Error
	}
	if user.ID == 0 {
		return nil, errors.New(wrongUsernameOrPassword)
	}
	if !checkPasswordHash(password, user.EncryptedPassword) {
		return nil, errors.New(wrongUsernameOrPassword)
	}

	return &user, nil
}

func Read(db *gorm.DB, id uint) (*model.User, error) {
	var user model.User
	db.First(&user, id)
	if db.Error != nil {
		return nil, db.Error
	}
	if user.ID == 0 {
		return nil, errors.New(invalidID)
	}
	return &user, nil
}
