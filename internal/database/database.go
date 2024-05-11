package database

import (
	"backend/internal/configs"
	"backend/internal/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var globalDb *gorm.DB

func SetupDb(cfg *configs.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDb, cfg.PostgresPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		panic("Error opening the database.")
	}

	_ = db.AutoMigrate(&model.User{})

	globalDb = db
}

// Assumption: SetupDb is called before this function
func GetDb() *gorm.DB {
	return globalDb
}
