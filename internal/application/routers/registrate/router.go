package registrate

import (
	"TelegramBot/internal/telegram"
	"TelegramBot/internal/telegram/options"
)

func RegisterRoutes(r *telegram.Router) {
	r.AddMessageHandler(start, options.WithMatchString("/start"))
	r.AddMessageHandler(getFirstName, options.WithState(firstNameState))
	r.AddMessageHandler(getLastName, options.WithState(lastNameState))
}
