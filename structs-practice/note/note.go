package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Println("Title:", n.Title)
	fmt.Println("Content:", n.Content)
}

func (n Note) Save() error {
	file_name := strings.ReplaceAll(n.Title, " ", "_")
	file_name = strings.ToLower(file_name) + ".json"
	jsonText, err := json.Marshal(n)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return os.WriteFile(file_name, jsonText, 0644)
}
