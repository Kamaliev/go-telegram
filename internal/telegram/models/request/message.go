package request

type Message struct {
	BusinessConnectionID int64  `json:"business_connection_id,omitempty"`
	ChatID               int64  `json:"chat_id,required"`
	MessageThreadID      int64  `json:"message_thread_id,omitempty"`
	Text                 string `json:"text,required"`
	ParseMode            string `json:"parse_mode,omitempty"`
}
