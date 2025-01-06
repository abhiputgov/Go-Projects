package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{Text: text}, nil
}

func (todo Todo) DisplayNote() {
	fmt.Printf("%v\n", todo.Text)
}

func (todo Todo) SaveToFile() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644) // will be nil if there is no error. and err is generated if something goes wrong
}
