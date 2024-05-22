package tgmsghandler

import (
	"backend/internal/database"
	"backend/internal/model"
	"backend/internal/ws"
	"backend/pkg/error/internalerror"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, hub *ws.Hub, tgMsg *tgbotapi.Message) error {
	db := database.GetDb()

	msgModel, err := saveTgMessageToDB(db, tgMsg, model.ByGuest)
	if err != nil {
		if internalerror.IsChatNotFoundError(err) {
			_, err := SendTelegramMessage(bot, tgMsg, NoChatFoundResponse)
			return err
		}
		if internalerror.IsRequestQueryNotFoundError(err) {
			_, err := SendTelegramMessage(bot, tgMsg, NoQueryFoundResponse)
			return err
		}
		return err
	}

	if err := broadcastMessage(hub, msgModel); err != nil {
		return err
	}

	res, err := getAIResponse(db, tgMsg.Chat.ID)
	if err != nil {
		return err
	}

	_, err = SendTelegramMessage(bot, tgMsg, res)
	return err
}
