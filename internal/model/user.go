package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	StaffRole Role = "staff"
	AdminRole Role = "admin"
)

type User struct {
	gorm.Model
	Username          string `gorm:"unique"`
	EncryptedPassword string
	Messages          []Message `gorm:"foreignKey:HotelStaffID"`
	Role              Role
}

func (u *User) Create(db *gorm.DB) error {
	err := db.Create(u).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return errors.New("username already taken")
	}
	return err
}

func (u *User) Update(db *gorm.DB) error {
	return db.Updates(u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.EncryptedPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.EncryptedPassword = string(bytes)
	return nil
}
