package response

type Message struct {
	MessageID         int    `json:"message_id"`
	MessageThreadID   int    `json:"message_thread_id"`
	From              *User  `json:"from"`
	SenderChat        *Chat  `json:"sender_chat"`
	SenderBoostCount  int    `json:"sender_boost_count"`
	SenderBusinessBot int    `json:"sender_business_bot"`
	Date              int    `json:"date"`
	Text              string `json:"text,omitempty"`
	Chat              *Chat  `json:"chat"`
}
