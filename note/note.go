package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
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
	fmt.Printf("Your note titled '%s' has the following content:\n\n%s\n", note.title, note.content)
}
