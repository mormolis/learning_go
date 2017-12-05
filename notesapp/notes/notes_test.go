package notes

import "testing"

func TestAddNote(t *testing.T) {
	notes := NotesList{}

	result := notes.AddNote("Note1", "some text")
	if result != 1 {
		t.Error("array length should increase")
	}
}

func TestRetrieveNotePositionByTitle(t *testing.T) {
	notes := NotesList{}
	notes.AddNote("Note1", "some text")
	notes.AddNote("Note2", "some text 2")
	notes.AddNote("Note3", "some text 3")
	result := notes.RetrieveNotePositionByTitle("Note2")
	if result != 1 {
		t.Error("It should return the position no 1")
	}

	result = notes.RetrieveNotePositionByTitle("Note33")
	if result != -1 {
		t.Error("There is no Note with that title, function should return -1")
	}
}

func TestRetrieveNoteByPosition(t *testing.T) {
	notes := NotesList{}
	notes.AddNote("Note1", "some text")
	notes.AddNote("Note2", "some text 2")
	notes.AddNote("Note3", "some text 3")

	result := notes.RetrieveNoteByPosition(1)
	if result != notes.notes[1] {
		t.Error("It should return the note object in position one")
	}

	result = notes.RetrieveNoteByPosition(4)
	n := note{}
	if result != n {
		t.Error("It should return the note object in position one")
	}

}
