package tgmessagehandler

import (
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/model"
	"backend/internal/ws"
	"backend/pkg/error/internalerror"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, hub *ws.Hub, msg *tgbotapi.Message) error {
	db := database.GetDb()

	chat, err := chat.ReadByTgChatID(db, msg.Chat.ID)
	if err != nil {
		if internalerror.IsRecordNotFoundError(err) {
			_, err := SendTelegramMessage(bot, msg, NoChatFoundResponse)
			return err
		}
		return err
	}

	msgModel, err := saveTgMessageToDB(db, msg, model.ByGuest)
	if err != nil {
		if internalerror.IsRequestQueryNotFoundError(err) {
			_, err := SendTelegramMessage(bot, msg, NoQueryFoundResponse)
			return err
		}
		return err
	}

	if err := broadcastMessage(hub, msgModel, chat.ID); err != nil {
		return err
	}

	aiResponse, err := getAIResponse(db, msg.Chat.ID)
	if err != nil {
		return err
	}

	aiReplyMsg, err := SendTelegramMessage(bot, msg, aiResponse)
	if err != nil {
		return err
	}

	msgModel, err = saveTgMessageToDB(db, aiReplyMsg, model.ByBot)
	if err != nil {
		return err
	}

	if err := broadcastMessage(hub, msgModel, chat.ID); err != nil {
		return err
	}

	return nil
}
