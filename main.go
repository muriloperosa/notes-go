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
			readNote()
		case 3:
			editNote()
		case 4:
			deleteNote()
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

func readNote() {

	note, err := display.SelectNote()

	if err != nil {
		output.Error(err)
		return
	}

	note.Show()
}

func editNote() {
	note, err := display.SelectNote()

	if err != nil {
		output.Error(err)
		return
	}

	title := input.Text("Enter the new title:")
	content := input.Text("Enter the new content:")

	err = note.Edit(title, content)

	if err != nil {
		output.Error(err)
		return
	}

	output.Notice("Note edited successfully!")

}

func deleteNote() {
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
