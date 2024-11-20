package main

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleUpdate(update tgbotapi.Update) tgbotapi.MessageConfig {
	if update.Message.Text == "/start" {
		return tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Напиши /card, чтобы получить метафорическую карту.")
	}
	return tgbotapi.MessageConfig{}
}

func TestHandleStartCommand(t *testing.T) {

	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "/start",
			Chat: &tgbotapi.Chat{
				ID: 1561345883,
			},
		},
	}

	msg := handleUpdate(update)

	expectedMessage := tgbotapi.NewMessage(1561345883, "Привет! Напиши /card, чтобы получить метафорическую карту.")

	if msg.ChatID != expectedMessage.ChatID || msg.Text != expectedMessage.Text {
		t.Errorf("expected message %v, got %v", expectedMessage, msg)
	}
}
