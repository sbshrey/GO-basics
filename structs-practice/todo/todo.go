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

	return Todo{
		Text: text,
	}, nil
}

func (t Todo) Display() {
	fmt.Println("Text:", t.Text)
}

func (t Todo) Save() error {
	file_name := "todo.json"
	jsonText, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return os.WriteFile(file_name, jsonText, 0644)
}
