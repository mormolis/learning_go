package notes

type NotesList struct {
	notes []note
}

func (n *NotesList) AddNote(title string, text string) int {
	n.notes = append(n.notes, note{})
	n.notes[len(n.notes)-1].CreateNote(title, text)
	return len(n.notes)
}

func (n *NotesList) RetrieveNotePositionByTitle(title string) int {
	for i := 0; i < len(n.notes); i++ {
		if n.notes[i].FindByTitle(title) {
			return i
		}
	}
	return -1
}

func (n *NotesList) RetrieveNoteByPosition(position int) note {
	var note note

	if position >= 0 && position < len(n.notes) {
		note = n.notes[position]
	}

	return note
}
