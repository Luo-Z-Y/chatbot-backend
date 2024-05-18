package tgcmd

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const helpText = "I'm a bot that can help you with your tasks. Here are the available commands:\n\n" +
	"/help - Show this message\n" +
	"/new - Start a new chat\n" +
	"/query - Make a simple query regarding the hotel\n" +
	"/request - Request a logistical request\n"

func HandleHelpCommand(msg *tgbotapi.Message) {
	msg.Text = helpText
}
