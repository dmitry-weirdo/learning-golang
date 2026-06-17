package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What would you me to scream?")

	var delimiter byte = '\n' // character syntax

	// os.Stdin super low-level
	// bufio decorates the os.Stdin
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString(delimiter) // we ignore the error from the tuple returned

	// trim the string
	s = strings.TrimSpace(s)

	sUpperCase := strings.ToUpper(s)

	fmt.Println(sUpperCase + "!")
}
