package response

type Update struct {
	UpdateId      int64          `json:"update_id"`
	Message       *Message       `json:"message,omitempty"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
}

type Updates []Update

type UpdateType string

const (
	MessageHandler       UpdateType = "message"
	CallbackQueryHandler            = "callback_query"
)

func (u Update) Type() UpdateType {
	if u.Message != nil {
		return MessageHandler
	} else if u.CallbackQuery != nil {
		return CallbackQueryHandler
	}
	panic("invalid Update type")
}
