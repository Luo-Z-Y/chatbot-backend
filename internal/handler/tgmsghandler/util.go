package tgmsghandler

import (
	"backend/internal/dataaccess/message"
	"backend/internal/dataaccess/requestquery"
	"backend/internal/model"
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
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

func GetAIResponse(chatID int64) string {
	// --- Perform query to AI to determine the type of query --- //
	return "Placeholder AI response"
}

func SendTextMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, text string) error {
	response := tgbotapi.NewMessage(msg.Chat.ID, text)
	response.ReplyToMessageID = msg.MessageID
	_, err := bot.Send(response)
	return err
}

func createRequestQueryTransaction(
	db *gorm.DB,
	tgMsg *tgbotapi.Message,
	chat *model.Chat,
	booking *model.Booking,
	queryType model.Type,
) error {
	result := db.Transaction(func(tx *gorm.DB) error {
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

		msg := model.Message{
			TelegramMessageId: int64(tgMsg.MessageID),
			By:                model.ByGuest,
			MessageBody:       tgMsg.Text,
			Timestamp:         tgMsg.Time(),
			RequestQueryId:    query.ID,
		}

		if err := message.Create(db, &msg); err != nil {
			return err
		}

		return nil
	})

	return result
}
