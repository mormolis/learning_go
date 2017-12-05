package notes

import (
	"fmt"
	"time"
)

type note struct {
	title string
	text  string
	date  time.Time
}

func (note *note) CreateNote(title string, text string) {
	note.text = text
	note.title = title
	note.date = time.Now()
}

func (note *note) FindByTitle(title string) bool {
	return note.title == title
}

func (note *note) DisplayNote() string {
	return note.title + "\n" + note.text + "\n" + fmt.Sprintf("%v", note.date)
}
