package main

import (
	"github.com/muriloperosa/notes-go/display"
	"github.com/muriloperosa/notes-go/input"
	"github.com/muriloperosa/notes-go/note"
	"github.com/muriloperosa/notes-go/output"
)

func main() {
	display.Greeting()

	for {
		display.Menu()
		output.BlankLine()
		opt := input.Int("Enter your choice:")

		switch opt {
		case 1:
			createNote()
		case 2:
			ReadNote()
		case 3:
			EditNote()
		case 4:
			DeleteNote()
		default:
			display.Goodbye()
			return
		}
	}
}

func createNote() {
	output.Notice("New Note")
	output.BlankLine()

	title := input.Text("Enter the title:")
	content := input.Text("Enter the note:")

	note, err := note.New(title, content)

	if err != nil {
		output.Error(err)
		return
	}

	err = note.Save()

	if err != nil {
		output.Error(err)
		return
	}

	output.Notice("New note created successfully!")
}

func ReadNote() {

	note, err := display.SelectNote()

	if err != nil {
		output.Error(err)
	}

	note.Show()
}

func EditNote() {}

func DeleteNote() {
	note, err := display.SelectNote()

	if err != nil {
		output.Error(err)
		return
	}

	confirmed := input.Confirmation("Are you sure to delete the note?")

	if !confirmed {
		output.Notice("Operation cancelled!")
		return
	}

	err = note.Delete()

	if err != nil {
		output.Error(err)
		return
	}

	output.Notice("Note deleted successfully!")
}
