package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"Vishrut.com/Types/note"
	"Vishrut.com/Types/todo"
)

type str string

type saver interface {
	Save() error
}
type output interface {
	saver
	Display() 
}

func main() {
	title, content := getNodeData()
	todoText := getUserinput("Todo Text: ")
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(errors.New("todo cannot be empty"))
		return
	}

	outputData(userTodo)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(errors.New("title or content cannot be empty"))
		return
	}

	outputData(userNote)

}

func getNodeData() (string, string) {
	title := getUserinput("Enter node title: ")

	content := getUserinput("Enter node content: ")
	return title, content
}

func outputData(data output) error { 
	data.Display()
	return data.Save()
}

// func saveData(data saver) error {
// 	err := data.Save()
// 	if err != nil {
// 		fmt.Println("Error saving note:", err)
// 		return err
// 	}
// 	fmt.Println("Note saved successfully.")
// 	return nil
// }

func getUserinput(prompt string) string {
	fmt.Print(prompt)
	// var input string
	// fmt.Scanln(&input)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(errors.New("HEY its Error"))
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}
