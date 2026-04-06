package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"content"`
}

func (todo Todo) Display() {
	fmt.Printf("Your todo %v has the following content\n\n%v\n\n", todo.Text, todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Value cannot be empty")
	}
	
	return Todo{
		Text: text,
	}, nil
}