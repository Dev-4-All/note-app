package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if isInputInvalid(title, content) {
		return Note{}, errors.New("Invalid input!")
	}

	return Note{
		title,
		content,
		time.Now(),
	}, nil
}

func isInputInvalid(title, content string) bool {
	return title == "" || content == ""
}

func (note Note) Display() {
	fmt.Printf("Your note titled '%s' has the following content:\n\n%s\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ToLower(note.Title)
	fileName = strings.ReplaceAll(fileName, " ", "_")
	fileName += ".json"

	fileContent, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, fileContent, 0644)
}
