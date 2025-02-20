package registrate

import (
	"TelegramBot/internal/telegram"
	"TelegramBot/internal/telegram/fsm"
	"fmt"
)

type UserData struct {
	FirstName string
	LastName  string
}

func start(context *telegram.ContextMessage) {
	context.Answer("Привет напиши имя")
	context.Bot.FSM().Set(context.UserID(), firstNameState)
}

func getFirstName(context *telegram.ContextMessage) {
	defer context.Bot.FSM().Set(context.UserID(), lastNameState)

	storage := fsm.GetFSM[UserData](context)
	storage.Set(UserData{FirstName: context.Update.Message.Text})
	context.Answer("Теперь фамилию")

}

func getLastName(context *telegram.ContextMessage) {
	defer context.Bot.FSM().Finish(context.UserID())

	storage := fsm.GetFSM[UserData](context)
	data := storage.Get()
	data.LastName = context.Update.Message.Text

	context.Answer(fmt.Sprintf("Все я тебя записал %s %s", data.FirstName, data.LastName))
}
