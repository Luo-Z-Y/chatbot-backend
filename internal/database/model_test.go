package database

import (
	"backend/internal/configs"
	"backend/internal/model"
	"fmt"
	"os"
	"path"
	"runtime"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCreateChat(t *testing.T) {
	SetupTestDb(t)
	t.Run("can create chat without booking id", func(t *testing.T) {
		db := GetDb()
		defer CleanUpTestDb(t, db)
		chat := model.Chat{TelegramChatId: 1}

		err := (&chat).Create(db)

		AssertNoErr(t, err)
	})
	t.Run("can create chat with booking id", func(t *testing.T) {
		db := GetDb()
		defer CleanUpTestDb(t, db)
		booking := model.Booking{LastName: "john", RoomNumber: "02-01"}
		chat := model.Chat{TelegramChatId: 1, Booking: &booking}

		AssertNoErr(t, chat.Create(db))
	})
}

func TestCreateMessage(t *testing.T) {
	createUserChatQuery := func(t testing.TB, db *gorm.DB) (query *model.RequestQuery, user *model.User) {
		chat := &model.Chat{TelegramChatId: 1}
		AssertNoErr(t, chat.Create(db))

		query = &model.RequestQuery{
			Status:    model.StatusOngoing,
			Type:      model.TypeUnknown,
			BookingID: nil,
		}
		err := db.Model(chat).Association("RequestQueries").Append(query)
		AssertNoErr(t, err)

		user = &model.User{
			Username:          "username",
			EncryptedPassword: "password",
		}
		AssertNoErr(t, user.Create(db))
		return
	}
	t.Run("can create message from staff", func(t *testing.T) {
		SetupTestDb(t)
		db := GetDb()
		defer CleanUpTestDb(t, db)

		query, user := createUserChatQuery(t, db)
		message := model.Message{
			TelegramMessageID: 1,
			By:                model.ByStaff,
			MessageBody:       "very cool message",
			Timestamp:         time.Now(),
			HotelStaffId:      &user.ID,
			RequestQueryId:    query.ID,
		}
		AssertNoErr(t, message.Create(db))

		// Association example
		var queryStaff model.User
		db.Preload("Messages").First(&queryStaff, user.ID)
		if len(queryStaff.Messages) != 1 || queryStaff.Messages[0].TelegramMessageID != 1 {
			t.Error("Message in query staff is not correct")
		}
	})
	t.Run("can create message from bot", func(t *testing.T) {
		SetupTestDb(t)
		db := GetDb()
		defer CleanUpTestDb(t, db)

		query, _ := createUserChatQuery(t, db)
		message := model.Message{
			TelegramMessageID: 1,
			By:                model.ByBot,
			MessageBody:       "very cool message",
			Timestamp:         time.Now(),
			RequestQueryId:    query.ID,
		}
		AssertNoErr(t, message.Create(db))
		fmt.Printf("%#v\n", message)
	})
}

func SetupTestDb(t testing.TB) {
	t.Helper()

	// Change directory to project directory
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	cfg, err := configs.GetConfig()
	if err != nil {
		panic(err)
	}
	SetupDb(cfg.GetTestDatabaseConfig())
}

func AssertNoErr(t testing.TB, err error) {
	if err != nil {
		t.Fatal("expected no error, got:", err)
	}
}

func CleanUpTestDb(t testing.TB, db *gorm.DB) {
	t.Helper()
	db.Unscoped().Where("1 = 1").Delete(&model.Message{})
	db.Unscoped().Where("1 = 1").Delete(&model.RequestQuery{})
	db.Unscoped().Where("1 = 1").Delete(&model.Booking{})
	db.Unscoped().Where("1 = 1").Delete(&model.Chat{})
	db.Unscoped().Where("1 = 1").Delete(&model.User{})
}
