package options

import "TelegramBot/internal/telegram/fsm"

type Options struct {
	matchString string
	state       fsm.State
}

func (ho Options) MatchString() *string {
	if ho.matchString == "" {
		return nil
	}
	return &ho.matchString
}

func (ho Options) State() *fsm.State {
	return &ho.state
}

type Option func(*Options)

func WithMatchString(matchString string) Option {
	return func(o *Options) {
		o.matchString = matchString
	}
}
func WithState(state fsm.State) Option {
	return func(o *Options) {
		o.state = state
	}
}
