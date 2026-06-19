package main

import (
	"bufio"
	"demo/menu"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

const DELIMITER = '\n'

func main() {

loop: // label
	for { // infinite loop of actions
		fmt.Println("")
		fmt.Println("=======================")
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) Quit")

		choice, _ := in.ReadString(DELIMITER)
		choice = strings.TrimSpace(choice) // required to trim the \n delimiter that will be read

		// see https://www.reddit.com/r/golang/comments/1oktsft/the_indentation_of_switch_statements_really/
		switch choice {
		case "1": // super-ugly formatting by gofmt, no indentation within the "switch"
			menu.Print()

		case "2":
			err := menu.Add()
			if err != nil {
				fmt.Println(fmt.Errorf("invalid input: %w", err))
			}

		case "q":
			// break // this will just break from `switch`, not the outer `for`
			break loop // break from the outer `for` by label

		default:
			fmt.Printf("Unknown option: %v \n", choice)
		}

		//fmt.Println(menu)
	}
}
