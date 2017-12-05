package main

import (
	"fmt"
	"learning_go/notesapp/notes"
)

func main() {
	notebook := notes.NotesList{}

	notebook.AddNote("Test title", "sometext")
	notebook.AddNote("Test title2", "sometext2")
	note := notebook.RetrieveNoteByPosition(1)
	fmt.Println(note.DisplayNote())
}
