package todo

import (
	"encoding/json"
	"errors"
	"os"
)

func (n Todo) Display() {
	println("Todo Content: ", n.Text)
}

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Save() error {
	fileName := "Todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return errors.New("Json error")
	}

	return os.WriteFile(fileName, json, 0644)

}
 
func New(text string) (Todo, error) {

	if text == "" {
		return Todo{}, errors.New("Failed to create todo: text cannot be empty")
	}
	return Todo{
		Text: text,
	}, nil
}
