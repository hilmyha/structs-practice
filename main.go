package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Note taker")
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(title)
	fmt.Println(content)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Title: ")

	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Content: ")

	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("Value cannot be empty")
	}
	
	return value, nil
}