package output

import "fmt"

func Error(err error) {
	fmt.Println("[ERROR] " + err.Error())
}

func TextLn(txt string) {
	fmt.Println(txt)
}

func BlankLine() {
	TextLn("")
}

func Notice(txt string) {
	BlankLine()
	fmt.Printf("### %v ###", txt)
	BlankLine()
}
