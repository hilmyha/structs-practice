package main

import (
	"fmt"

	"github.com/hilmyha/structs-practice/note"
)

func main() {
	fmt.Println("Note taker")
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Title: ")
	content := getUserInput("Content: ")

	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	
	return value
}