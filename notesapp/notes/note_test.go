package notes

import (
	"fmt"
	"testing"
)

func TestFindNoteByTitle(t *testing.T) {
	note := note{}
	note.CreateNote("Note1", "some text")
	result := note.FindByTitle("Note1")
	if result != true {
		t.Error("FindByTitle should return true when title given is the same as note")
	}
}

func TestDisplayNote(t *testing.T) {
	note := note{}
	note.CreateNote("Note1", "some text")
	result := note.DisplayNote()

	if result != "Note1\nsome text\n"+fmt.Sprintf("%v", note.date) {
		t.Error("should return the note formated as text")
	}

}
