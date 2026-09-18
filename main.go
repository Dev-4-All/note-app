package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	fmt.Print(prompt)
	
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}
