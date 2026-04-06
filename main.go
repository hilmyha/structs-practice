package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hilmyha/structs-practice/note"
	"github.com/hilmyha/structs-practice/todo"
)

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


	userTodo.Display()
	err = userTodo.Save()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Todo saved successfully")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Note saved successfully")
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