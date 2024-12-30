package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func getNoteDate() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")
	return title, content
}

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteDate()
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := getUserInput("Todo Text: ")
	todo, err := todo.New(text)
	if err != nil {
		fmt.Print(err)
		return
	}

	todo.Display()
	err = todo.Save()
	if err != nil {
		fmt.Println("Saving the todo failed: ", err)
		return
	}
	fmt.Println("Saving the todo succeeded!")

	note.Display()
	err = note.Save()
	if err != nil {
		fmt.Println("Saving the note failed: ", err)
		return
	}
	fmt.Println("Saving the note succeeded!")
}

func saveData

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	//does not work for long string input containing spaces
	// var value string
	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
