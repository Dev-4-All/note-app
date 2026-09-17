package main

import (
	"fmt"

	"example.com/note-app/note"
)

func main() {
	title, content := getNoteData()

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Enter title: ")
	content := getUserInput("Enter content: ")

	return title, content
}

func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)

	return input
}
