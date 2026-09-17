package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData() (string, string, error) {
	title := getUserInput("Enter title: ")
	content := getUserInput("Enter content: ")

	if isInvalid(title) || isInvalid(content) {
		return "", "", errors.New("Invalid input!")
	}

	return title, content, nil
}

func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)

	return input
}

func isInvalid(input string) bool {
	if input == "" {
		return true
	}

	return false
}
