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
)

func HandleRequestCommand(msg *tgbotapi.Message) (string, error) {
	tgChatID := msg.Chat.ID

	db := database.GetDb()

	chat, err := chat.ReadByTgChatID(db, tgChatID)
	if err != nil {
		return NoChatFoundResponse, err
	}

	bk, err := booking.ReadByChatID(db, chat.ID)
	if err != nil && !internalerror.IsRecordNotFoundError(err) {
		return "An error occurred while fetching booking", err
	}

	if err := createRequestQueryTransaction(db, msg, chat, bk, model.TypeRequest); err != nil {
		if externalerror.IsAuthRequiredError(err) {
			return AuthRequiredErrorResponse, err
		}
		return "An error occurred while creating a new request", err
	}

	return "New query created", nil
}
