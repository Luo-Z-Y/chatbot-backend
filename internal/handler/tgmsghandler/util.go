package tgmsghandler

import (
	"backend/internal/api"
	"backend/internal/dataaccess/chat"
	"backend/internal/dataaccess/message"
	"backend/internal/dataaccess/requestquery"
	"backend/internal/model"
	"backend/internal/viewmodel"
	"backend/internal/ws"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

var (
	NoQueryFoundResponse = fmt.Sprintf(
		"Please start a new query first. Use /%s, /%s, or /%s",
		AskCmdWord, QueryCmdWord, RequestCmdWord,
	)
	NoQueryFoundErr = errors.New(NoQueryFoundResponse)
)

func _tempRandomType() model.Type {
	decision := rand.Intn(2)
	switch decision {
	case 0:
		return model.TypeQuery
	case 1:
		return model.TypeRequest
	default:
		return model.TypeUnknown
	}
}

func getAIResponse(_ *gorm.DB, _ int64) (string, error) {
	response := "Placeholder AI response"
	return response, nil
}

func saveTgMessageToDB(db *gorm.DB, msg *tgbotapi.Message, by model.By) (*model.Message, error) {
	chat, err := chat.ReadByTgChatID(db, msg.Chat.ID)
	if err != nil {
		return nil, err
	}

	rqq, err := requestquery.ReadLatestByChatID(db, chat.ID)
	if err != nil {
		return nil, err
	}

	msgModel := model.Message{
		TelegramMessageId: chat.TelegramChatId,
		By:                by,
		MessageBody:       msg.Text,
		Timestamp:         time.Now(),
		RequestQueryId:    rqq.ID,
	}

	if err := message.Create(db, &msgModel); err != nil {
		return nil, err
	}

	return &msgModel, nil
}

func createRequestQuery(
	db *gorm.DB,
	queryType model.Type,
	chat *model.Chat,
	booking *model.Booking,
) error {
	query := model.RequestQuery{
		Status: model.StatusOngoing,
		Type:   queryType,
		ChatID: chat.ID,
	}

	if booking != nil {
		query.BookingID = &booking.ID
	}

	if err := requestquery.Create(db, &query); err != nil {
		return err
	}

	return nil
}

func sendTelegramMessage(
	bot *tgbotapi.BotAPI,
	prompt *tgbotapi.Message,
	content string,
) (*tgbotapi.Message, error) {
	response := tgbotapi.NewMessage(prompt.Chat.ID, content)

	if prompt != nil {
		response.ReplyToMessageID = prompt.MessageID
	}

	msg, err := bot.Send(response)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func broadcastMessage(hub *ws.Hub, msg *model.Message, by model.By) error {
	msgView := viewmodel.BaseMessageView{
		TelegramMessageId: msg.TelegramMessageId,
		By:                string(by),
		MessageBody:       msg.MessageBody,
		Timestamp:         msg.Timestamp.Format(time.RFC3339),
		RequestQueryId:    msg.RequestQueryId,
	}

	msgStruct := api.WebSocketMessage{
		Type: api.MessageReceivedType,
		Data: msgView,
	}

	msgBytes, err := json.Marshal(msgStruct)
	if err != nil {
		return err
	}

	hub.Broadcast <- msgBytes
	return nil
}
