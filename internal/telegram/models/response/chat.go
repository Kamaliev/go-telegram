package response

import (
	"encoding/json"
	"errors"
)

type ChatType string

const (
	Private    ChatType = "private"
	Group               = "group"
	Channel             = "channel"
	SuperGroup          = "supergroup"
)

func (c *ChatType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch ChatType(str) {
	case Private, Group, Channel, SuperGroup:
		*c = ChatType(str) // Присваиваем значение
		return nil
	}

	return errors.New("invalid ChatType value")
}

type Chat struct {
	ID        int64    `json:"id"`
	FirstName string   `json:"first_name"`
	Username  string   `json:"username"`
	Type      ChatType `json:"type"`
}
