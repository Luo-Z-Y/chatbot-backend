package tgmsghandler

import (
	"backend/internal/dataaccess/booking"
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/model"
	"backend/internal/ws"
	"backend/pkg/error/externalerror"
	"backend/pkg/error/internalerror"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	RequestCmdWord = "request"
	RequestCmdDesc = "Make a logistical request"
)

const (
	AuthRequiredErrorResponse = "Authentication required before request can be processed, please provide your booking ID"
	RequestCreatedResponse    = "New request created, please wait for a response"
)

func HandleRequestCommand(bot *tgbotapi.BotAPI, hub *ws.Hub, msg *tgbotapi.Message) error {
	db := database.GetDb()

	chat, err := chat.ReadByTgChatID(db, msg.Chat.ID)
	if err != nil {
		if internalerror.IsRecordNotFoundError(err) {
			_, err := SendTelegramMessage(bot, msg, NoChatFoundResponse)
			return err
		}
		return err
	}

	bk, err := booking.ReadByChatID(db, chat.ID)
	if err != nil && !internalerror.IsRecordNotFoundError(err) {
		return err
	}

	if err := createRequestQuery(db, model.TypeRequest, chat, bk); err != nil {
		if externalerror.IsAuthRequiredError(err) {
			_, err := SendTelegramMessage(bot, msg, AuthRequiredErrorResponse)
			return err
		}
		return err
	}

	msgModel, err := saveTgMessageToDB(db, msg, model.ByGuest)
	if err != nil {
		return err
	}

	if err := broadcastMessage(hub, msgModel); err != nil {
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

	if err := broadcastMessage(hub, msgModel); err != nil {
		return err
	}

	return nil
}
