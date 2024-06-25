package display

import (
	"errors"

	"github.com/muriloperosa/notes-go/input"
	"github.com/muriloperosa/notes-go/note"
	"github.com/muriloperosa/notes-go/output"
)

func Greeting() {
	output.Notice("Welcome to your Go Notes app!")
}

func Menu() {
	output.BlankLine()
	output.TextLn("Choose an option to continue:")
	output.BlankLine()
	output.TextLn("[0] Exit")
	output.TextLn("[1] Create a new note")
	output.TextLn("[2] Read note")
	output.TextLn("[3] Edit note")
	output.TextLn("[4] Delete note")
}

func Goodbye() {
	output.Notice("Bye! Come back anytime.")
}

func SelectNote() (*note.Note, error) {
	notes, err := note.GetAll()

	if err != nil {
		return &note.Note{}, err
	}

	if len(notes) == 0 {
		return &note.Note{}, errors.New("no notes found")
	}

	err = note.ListAll(notes)

	if err != nil {
		return &note.Note{}, err
	}

	output.BlankLine()
	opt := input.Int("Enter your the note code: ")

	if opt > len(notes) || opt <= 0 {
		return &note.Note{}, errors.New("invalid option selected")
	}

	return &notes[opt-1], nil
}
