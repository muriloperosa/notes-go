package main

import (
	"github.com/muriloperosa/notes-go/input"
	"github.com/muriloperosa/notes-go/note"
	"github.com/muriloperosa/notes-go/output"
)

func main() {
	output.TextLn("[NEW NOTE]")
	output.TextLn(" ")

	title := input.Text("Note Title:")
	content := input.Text("Note Content:")

	note, err := note.New(title, content)

	if err != nil {
		output.Error(err)
		return
	}

	output.TextLn("=======")
	note.Show()

	err = note.Save()

	if err != nil {
		output.Error(err)
		return
	}

	output.TextLn("Note saved!")
}
