package telegram

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("Ошибка маршалинга JSON:", err)
		return
	}
	fmt.Println(string(b))
}
