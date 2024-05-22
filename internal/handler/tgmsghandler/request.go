package tgmsghandler

import (
	"backend/internal/dataaccess/booking"
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/model"
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

func HandleRequestCommand(msg *tgbotapi.Message) (string, error) {
	tgChatID := msg.Chat.ID

	db := database.GetDb()

	chat, err := chat.ReadByTgChatID(db, tgChatID)
	if err != nil {
		if internalerror.IsRecordNotFoundError(err) {
			return NoChatFoundResponse, nil
		}
		return "", err
	}

	bk, err := booking.ReadByChatID(db, chat.ID)
	if err != nil && !internalerror.IsRecordNotFoundError(err) {
		return "", err
	}

	if err := createRequestQuery(db, model.TypeRequest, chat, bk); err != nil {
		if externalerror.IsAuthRequiredError(err) {
			return AuthRequiredErrorResponse, err
		}
		return "", err
	}

	return RequestCreatedResponse, nil
}
