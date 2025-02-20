package telegram

import (
	"TelegramBot/internal/telegram/models/request"
	"TelegramBot/internal/telegram/models/response"
)

func (b *Bot) getUpdates(params request.Update) (response.Updates, error) {
	method := "getUpdates"
	url := b.apiUrl + "/" + method
	var updates response.Updates
	client := Request{b.httpClient}
	err := client.Post(url, ContentTypeJSON, params, &updates)
	if err != nil {
		return response.Updates{}, err
	}

	return updates, nil
}

func (b *Bot) SendMessage(params request.Message) (response.Message, error) {
	method := "sendMessage"
	url := b.apiUrl + "/" + method
	var message response.Message
	client := Request{b.httpClient}
	err := client.Post(url, ContentTypeJSON, params, &message)
	if err != nil {
		return response.Message{}, err
	}
	return message, nil
}
