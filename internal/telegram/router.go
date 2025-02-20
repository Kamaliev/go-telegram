package telegram

import (
	"TelegramBot/internal/telegram/models/response"
	"TelegramBot/internal/telegram/options"
	"sync"
)

type AbstractRouter interface {
	AddMessageHandler(handler func(ctx *ContextMessage), opts ...options.Option)
	AddCallbackQueryHandler(handler func(ctx *ContextCallbackQuery), opts ...options.Option)
}

type Router struct {
	mu       sync.RWMutex
	handlers map[response.UpdateType][]Handler
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[response.UpdateType][]Handler),
	}
}

func (router *Router) AddMessageHandler(handler func(ctx *ContextMessage), opts ...options.Option) {
	router.mu.Lock()
	defer router.mu.Unlock()
	var opt options.Options
	for _, o := range opts {
		o(&opt)
	}
	router.handlers[response.MessageHandler] = append(router.handlers[response.MessageHandler],
		MessageHandler{
			handle:     handler,
			matchSting: opt.MatchString(),
			state:      opt.State(),
		})
}

func (router *Router) AddCallbackQueryHandler(handler func(ctx *ContextCallbackQuery), opts ...options.Option) {
	router.mu.Lock()
	defer router.mu.Unlock()

	var opt options.Options

	for _, o := range opts {
		o(&opt)
	}
	router.handlers[response.CallbackQueryHandler] = append(router.handlers[response.CallbackQueryHandler],
		CallbackQueryHandler{
			handle:     handler,
			matchSting: opt.MatchString(),
			state:      opt.State(),
		})
}
