package options

import (
	"TelegramBot/internal/telegram/domain"
)

type Options struct {
	matchString string
	state       domain.State
}

func (ho Options) MatchString() *string {
	if ho.matchString == "" {
		return nil
	}
	return &ho.matchString
}

func (ho Options) State() *domain.State {
	return &ho.state
}

type Option func(*Options)

func WithMatchString(matchString string) Option {
	return func(o *Options) {
		o.matchString = matchString
	}
}
func WithState(state domain.State) Option {
	return func(o *Options) {
		o.state = state
	}
}
