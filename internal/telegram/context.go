package telegram

import (
	"TelegramBot/internal/telegram/domain"
	"TelegramBot/internal/telegram/models/request"
	"TelegramBot/internal/telegram/models/response"
)

type Context struct {
	Bot    *Bot
	Update *response.Update
}

func (ctx *Context) FSM() domain.AbstractFsm {
	return ctx.Bot.fsm
}

func NewContext(bot *Bot, update *response.Update) *Context {
	return &Context{Bot: bot, Update: update}
}

// region Message
type ContextMessage struct {
	Context
}

func NewContextMessage(bot *Bot, update *response.Update) *ContextMessage {
	return &ContextMessage{Context: Context{Bot: bot, Update: update}}
}

func (cm ContextMessage) Answer(text string) {
	_, _ = cm.Bot.SendMessage(
		request.Message{
			ChatID: cm.Update.Message.From.UserID,
			Text:   text,
		})
}

func (cm ContextMessage) UserID() int64 {
	return cm.Update.Message.From.UserID
}

// endregion

// region CallbackQuery
type ContextCallbackQuery struct {
	Context
}

func NewContextCallbackQuery(bot *Bot, update *response.Update) *ContextCallbackQuery {
	return &ContextCallbackQuery{Context: Context{Bot: bot, Update: update}}
}

func (cb ContextCallbackQuery) Answer(text string) {
	_, _ = cb.Bot.SendMessage(
		request.Message{
			ChatID: cb.Update.CallbackQuery.From.UserID,
			Text:   text,
		},
	)
}

func (cb ContextCallbackQuery) UserID() int64 {
	return cb.Update.CallbackQuery.Message.From.UserID
}

// endregion
