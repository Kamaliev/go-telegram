package domain

type AbstractContext interface {
	Answer(text string)
	UserID() int64
	FSM() AbstractFsm
}
