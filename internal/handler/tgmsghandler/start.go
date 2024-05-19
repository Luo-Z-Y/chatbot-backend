package tgmsghandler

import (
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/model"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	StartCmdWord = "start"
	StartCmdDesc = "Start a new chat"
)

var (
	SuccessChatCreationResponse = fmt.Sprintf(
		"New chat created, you can now start a new query using /%s, /%s, or /%s",
		AskCmdWord, QueryCmdWord, RequestCmdWord,
	)
)

func HandleStartCommand(msg *tgbotapi.Message) (string, error) {
	tgChatID := msg.Chat.ID

	db := database.GetDb()

	tgChat := model.Chat{
		TelegramChatId: tgChatID,
	}

	if err := chat.Create(db, &tgChat); err != nil {
		return "You have already started the chat :)", err
	}

	return "New chat created", nil
}
