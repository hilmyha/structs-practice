package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hilmyha/structs-practice/note"
	"github.com/hilmyha/structs-practice/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputter interface {
	saver
	displayer
}

func main() {
	fmt.Println("Note taker")
	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")

	userTodo, err :=	todo.New(todoText)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	// userTodo.Display()
	// err = saveData(userTodo)
	err = saveDataAndDisplay(userTodo)

	if err != nil {
		return
	}

	// userNote.Display()
	// err = saveData(userNote)
	err = saveDataAndDisplay(userNote)

	if err != nil {
		return
	}
}

func getTodoData() string {
	text := getUserInput("Todo text: ")
	
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Title: ")
	content := getUserInput("Content: ")

	return title, content
}

func saveDataAndDisplay(data outputter) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Data saved successfully")
	
	return nil
}

func getUserInput(prompt string) (string) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	
	return text
}