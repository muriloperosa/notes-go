package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/muriloperosa/notes-go/output"
)

func Text(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}

func Int(prompt string) int {
	fmt.Printf("%v ", prompt)

	var value int
	_, err := fmt.Scan(&value)

	if err != nil {
		output.Error(err)
	}

	return value
}

func Confirmation(question string) bool {
	output.BlankLine()
	fmt.Printf("%v [y/n] ", question)
	var value string
	fmt.Scan(&value)

	return strings.ToLower(value) == "y"
}
