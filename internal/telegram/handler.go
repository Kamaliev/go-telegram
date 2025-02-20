package telegram

import (
	"TelegramBot/internal/telegram/domain"
)

type Handler interface {
	Filter(context *Context) bool
	HandleMessage(context *ContextMessage)
	HandleCallbackQuery(context *ContextCallbackQuery)
}

type defaultHandler struct{}

func (handler defaultHandler) HandleMessage(context *ContextMessage)             {}
func (handler defaultHandler) HandleCallbackQuery(context *ContextCallbackQuery) {}

type MessageHandler struct {
	handle     func(context *ContextMessage)
	matchSting *string
	state      *domain.State
	defaultHandler
}

func (m MessageHandler) Filter(context *Context) bool {
	userID := context.Update.Message.From.UserID
	state, hasState := context.Bot.FSM().Current(userID)

	if hasState {
		if m.matchSting != nil {
			return (m.state != nil && *m.state == state) && *m.matchSting == context.Update.Message.Text
		}
		return m.state != nil && *m.state == state
	}

	return m.matchSting != nil && *m.matchSting == context.Update.Message.Text
}

func (m MessageHandler) HandleMessage(context *ContextMessage) {
	m.handle(context)
}

type CallbackQueryHandler struct {
	handle     func(context *ContextCallbackQuery)
	matchSting *string
	state      *domain.State
	defaultHandler
}

func (c CallbackQueryHandler) Filter(context *Context) bool {
	userID := context.Update.Message.From.UserID
	state, hasState := context.Bot.FSM().Current(userID)

	if hasState {
		if c.matchSting != nil {
			return (c.state != nil && *c.state == state) && *c.matchSting == context.Update.Message.Text
		}
		return c.state != nil && *c.state == state
	}

	return c.matchSting != nil && *c.matchSting == context.Update.Message.Text
}

func (c CallbackQueryHandler) HandleCallbackQuery(context *ContextCallbackQuery) {
	c.handle(context)
}
