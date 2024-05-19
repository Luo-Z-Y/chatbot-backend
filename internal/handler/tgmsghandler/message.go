package tgmsghandler

import (
	"backend/internal/dataaccess/message"
	"backend/internal/dataaccess/requestquery"
	"backend/internal/database"
	"backend/internal/model"
	"backend/internal/ws"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, hub *ws.Hub, tgMsg *tgbotapi.Message) error {
	db := database.GetDb()
	var response string

	rqq, err := requestquery.ReadLatestByChatID(db, uint(tgMsg.Chat.ID))
	if err != nil {
		response = fmt.Sprintf(
			"Please start a new query first. Use /%s, /%s, or /%s",
			AskCmdWord, QueryCmdWord, RequestCmdWord,
		)
		return SendTextMessage(bot, tgMsg, response)
	}

	msg := model.Message{
		TelegramMessageId: int64(tgMsg.MessageID),
		By:                model.ByGuest,
		MessageBody:       tgMsg.Text,
		Timestamp:         tgMsg.Time(),
		RequestQueryId:    rqq.ID,
	}

	if err := message.Create(db, &msg); err != nil {
		return err
	}

	response = GetAIResponse(tgMsg.Chat.ID)
	return SendTextMessage(bot, tgMsg, response)
}
