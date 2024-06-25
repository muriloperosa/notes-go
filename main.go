package main

import (
	"errors"

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
		opt := input.Int("Enter your choice: ")

		switch opt {
		case 1:
			createNote()
		case 2:
			ReadNote()
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
	notes, err := note.GetAll()

	if err != nil {
		output.Error(err)
		return
	}

	if len(notes) == 0 {
		output.Notice("No notes found!")
		return
	}

	err = note.ListAll(notes)

	if err != nil {
		output.Error(err)
		return
	}

	output.BlankLine()
	opt := input.Int("Enter your the note code: ")

	if opt > len(notes) || opt <= 0 {
		output.Error(errors.New("invalid option selected"))
		return
	}

	notes[opt-1].Show()
}
