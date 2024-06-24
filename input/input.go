package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
