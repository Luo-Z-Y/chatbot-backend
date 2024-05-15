package database

import (
	"backend/internal/configs"
	"backend/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var globalDb *gorm.DB

func SetupDb(cfg *configs.PostgresConfig) {
	dsn, err := BuildDsn(cfg)
	if err != nil {
		panic("Error building the DSN.")
	}

	gormCfg := GetConfig()
	db, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		panic("Error opening the database.")
	}

	_ = db.AutoMigrate(&model.User{})
	_ = db.AutoMigrate(&model.Booking{})
	_ = db.AutoMigrate(&model.Chat{})
	_ = db.AutoMigrate(&model.RequestQuery{})
	_ = db.AutoMigrate(&model.Message{})

	globalDb = db
}

// Assumption: SetupDb is called before this function
func GetDb() *gorm.DB {
	return globalDb.Session(&gorm.Session{NewDB: true})
}
