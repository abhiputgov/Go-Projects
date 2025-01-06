package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	SaveToFile() error
}

type displayer interface {
	saver
	DisplayNote()
}

func main() {
	title, content := getNoteData()
	text := getTodoData()
	UserNote, errNote := note.New(title, content)
	UserTodo, errTodo := todo.New(text)
	if errNote != nil {
		fmt.Println(errNote)
	}
	if errTodo != nil {
		fmt.Printf("%v ", errTodo)
	}

	outputData(UserNote)
	outputData(UserTodo)
	printSomething(1)
	printSomething(1.5)
	printSomething("hey")
}

func printSomething(data interface{}) {
	// switch data.(type) {
	// case int:
	// case float64:
	// case string:
	// }

	// another alternative to the above code
	typedVal, isInt := data.(int)
	if isInt {
		fmt.Println("integer value of:\t", typedVal)
	}
	typedVal2, isFloat := data.(float64)
	if isFloat {
		fmt.Println("float64 value of:\t", typedVal2)
	}
	typedVal3, isString := data.(string)
	if isString {
		fmt.Println("float64 value of:\t", typedVal3)
	}

}
func outputData(data displayer) error {
	data.DisplayNote()
	err := saveData(data)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func saveData(data saver) error {
	err := data.SaveToFile()
	if err != nil {
		fmt.Println("todo not saved")
		return err
	}
	fmt.Println("Saved as file!")
	return nil
}

func getNoteData() (string, string) {
	title, err := getUserInput("Note title:")
	if err != nil {
		fmt.Println(err)
	}
	content, err := getUserInput("Note content:")
	if err != nil {
		fmt.Println(err)
	}
	return title, content
}

func getTodoData() string {
	text, err := getUserInput("Todo text:")
	if err != nil {
		fmt.Printf("%v ", err)
	}
	return text
}

func getUserInput(prompt string) (string, error) {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n') // rune dont forget
	if err != nil {
		return "", err
	}
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return value, nil

}
