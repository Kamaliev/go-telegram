package application

import (
	"TelegramBot/internal/application/routers/registrate"
	"TelegramBot/internal/telegram"
	"os"
)

type App struct{}

func NewRouter() *telegram.Router {
	router := telegram.NewRouter()
	registrate.RegisterRoutes(router)
	return router
}

func (app App) Run() {
	token := os.Getenv("TELEGRAM_TOKEN")
	bot := telegram.NewBot(token, NewRouter())
	bot.Start()
}
