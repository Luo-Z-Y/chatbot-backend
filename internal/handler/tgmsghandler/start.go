package tgmsghandler

import (
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/model"
	"backend/internal/ws"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	StartCmdWord = "start"
	StartCmdDesc = "Start a new chat"
)

const (
	ChatAlreadyExistsResponse = "You have already started the chat :)"
	ChatCreatedResponse       = "New chat created"
)

var (
	SuccessChatCreationResponse = fmt.Sprintf(
		"New chat created, you can now start a new query using /%s, /%s, or /%s",
		AskCmdWord, QueryCmdWord, RequestCmdWord,
	)
)

func HandleStartCommand(bot *tgbotapi.BotAPI, hub *ws.Hub, msg *tgbotapi.Message) error {
	// There is no need to broadcast the chat creation message to the since the chat is just created
	db := database.GetDb()

	tgChat := model.Chat{
		TelegramChatId: msg.Chat.ID,
	}

	if err := chat.Create(db, &tgChat); err != nil {
		_, err = sendTelegramMessage(bot, msg, ChatAlreadyExistsResponse)
		return err
	}

	_, err := sendTelegramMessage(bot, msg, ChatCreatedResponse)
	return err
}
