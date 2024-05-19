package authhandler

import (
	"backend/internal/dataaccess/booking"
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/params/bookingparams"
	"backend/internal/telegram"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
)

const (
	AuthSuccessMsg = "You have been authenticated successfully"
)

// The current logic is that, as long as there exist a booking with the same chat id, the user is authenticated.
func AuthenticateTgUser(c echo.Context) error {
	bkParams := bookingparams.Params{}

	if err := c.Bind(&bkParams); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&bkParams); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db := database.GetDb()

	bk := bkParams.ToModel()

	chat, err := chat.Read(db, bk.ChatId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := booking.Create(db, bk); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	tgBotApi := telegram.GetBotAPI()

	response := tgbotapi.NewMessage(int64(chat.TelegramChatId), AuthSuccessMsg)
	_, err = tgBotApi.Send(response)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
