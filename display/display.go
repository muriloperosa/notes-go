package display

import "github.com/muriloperosa/notes-go/output"

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
	output.TextLn("[3] ...")
}

func Goodbye() {
	output.Notice("Bye! Come back anytime.")
}
