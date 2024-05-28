package tgmessagehandler

import (
	"backend/internal/dataaccess/booking"
	"backend/internal/database"
	"backend/internal/model"
	"backend/internal/ws"
	"backend/pkg/error/internalerror"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	QueryCmdWord = "query"
	QueryCmdDesc = "Make a simple query regarding the hotel"
)

func HandleQueryCommand(bot *tgbotapi.BotAPI, hub *ws.Hub, msg *tgbotapi.Message) error {
	db := database.GetDb()

	chat, err := readChatByTgChatIDOrCreate(db, msg.Chat.ID)
	if err != nil {
		return err
	}

	bk, err := booking.ReadByChatID(db, chat.ID)
	if err != nil && !internalerror.IsRecordNotFoundError(err) {
		return err
	}

	if err := createRequestQuery(db, model.TypeQuery, chat, bk); err != nil {
		return err
	}

	msgModel, err := saveTgMessageToDB(db, msg, model.ByGuest)
	if err != nil {
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
