package tgmessagehandler

import (
	"backend/internal/api"
	"backend/internal/dataaccess/chat"
	"backend/internal/dataaccess/message"
	"backend/internal/dataaccess/requestquery"
	"backend/internal/model"
	"backend/internal/viewmodel"
	"backend/internal/ws"
	"backend/pkg/error/internalerror"
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
		"Please start a new query first. Use /%s, or /%s",
		QueryCmdWord, RequestCmdWord,
	)
	NoQueryFoundErr = internalerror.RequestQueryNotFoundError{}
	NoChatFoundErr  = internalerror.ChatNotFoundError{}
)

// TODO: Implement AI response
func getAICategorisation(_ *gorm.DB, _ string) (model.Type, error) {
	decision := rand.Intn(3)
	switch decision {
	case 0:
		return model.TypeQuery, nil
	case 1:
		return model.TypeRequest, nil
	default:
		return model.TypeUnknown, nil
	}
}

// TODO: Implement AI response
func getAIResponse(_ *gorm.DB, _ int64) (string, error) {
	response := "Placeholder AI response"
	return response, nil
}

func saveTgMessageToDB(db *gorm.DB, msg *tgbotapi.Message, by model.By) (*model.Message, error) {
	chat, err := chat.ReadByTgChatID(db, msg.Chat.ID)
	if err != nil {
		if internalerror.IsRecordNotFoundError(err) {
			return nil, NoChatFoundErr
		}
		return nil, err
	}

	rqq, err := requestquery.ReadLatestByChatID(db, chat.ID)
	if err != nil {
		if internalerror.IsRecordNotFoundError(err) {
			return nil, NoQueryFoundErr
		}
		return nil, err
	}

	msgModel := model.Message{
		TelegramMessageID: int64(msg.MessageID),
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

// Booking is nillable
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

func SendTelegramMessage(
	bot *tgbotapi.BotAPI,
	prompt *tgbotapi.Message,
	content string,
) (*tgbotapi.Message, error) {
	if prompt == nil || prompt.Chat == nil {
		return nil, errors.New("Bad prompt message")
	}

	response := tgbotapi.NewMessage(prompt.Chat.ID, content)
	response.ReplyToMessageID = prompt.MessageID

	msg, err := bot.Send(response)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func broadcast(hub *ws.Hub, t string, v any) error {
	msgStruct := api.WebSocketMessage{
		Type: t,
		Data: v,
	}

	msgBytes, err := json.Marshal(msgStruct)
	if err != nil {
		return err
	}

	hub.Broadcast <- msgBytes
	return nil
}

func broadcastMessage(hub *ws.Hub, msg *model.Message, chatID uint) error {
	msgView := viewmodel.MessageWebSocketView{
		BaseMessageView: viewmodel.BaseMessageView{
			TelegramMessageId: msg.TelegramMessageID,
			By:                string(msg.By),
			MessageBody:       msg.MessageBody,
			Timestamp:         msg.Timestamp.Format(time.RFC3339),
			RequestQueryId:    msg.RequestQueryId,
		},
		ChatID: chatID,
	}

	if err := broadcast(hub, api.MessageReceivedType, msgView); err != nil {
		return err
	}

	return nil
}

func broadcastAuthRequest(hub *ws.Hub, chat *model.Chat, cred string) error {
	tgAuthView := viewmodel.TgAuthView{
		ChatID:      chat.ID,
		Credentials: cred,
	}

	if err := broadcast(hub, api.AuthType, tgAuthView); err != nil {
		return err
	}

	return nil
}
