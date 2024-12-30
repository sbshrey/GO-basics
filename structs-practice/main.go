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

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("something")
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

	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	printSomething(todo)

	err = outputData(note)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func printSomething(value interface{}) {
	intVal, ok := value.(int) // ok is boolean
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}
	floatVal, ok := value.(float64) // ok is boolean
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}
	stringVal, ok := value.(string) // ok is boolean
	if ok {
		fmt.Println("String:", stringVal)
		return
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed: ", err)
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
}

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