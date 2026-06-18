package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{
			name: "Coffee",
			prices: map[string]float64{
				"small":  1.65,
				"medium": 1.80,
				"large":  1.95,
			},
		},
		{
			name: "Espresso",
			prices: map[string]float64{
				"single": 1.90,
				"double": 2.25,
				"triple": 2.55,
			},
		},
	}

loop: // label
	for { // infinite loop of actions
		fmt.Println("")
		fmt.Println("=======================")
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) Quit")

		in := bufio.NewReader(os.Stdin)

		const delimiter = '\n'
		choice, _ := in.ReadString(delimiter)
		choice = strings.TrimSpace(choice) // required to trim the \n delimiter that will be read

		// see https://www.reddit.com/r/golang/comments/1oktsft/the_indentation_of_switch_statements_really/
		switch choice {
		case "1": // super-ugly formatting by gofmt, no indentation within the "switch"
			for _, item := range menu { // ignore the array index
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))

				// iterate the prices
				for size, price := range item.prices {
					fmt.Printf("\t%10s: %10.2f\n", size, price) // %10 is "fill with spaces up to 10 characters"
				}
			}

		case "2":
			fmt.Println("Please enter the name of the new item")

			name, _ := in.ReadString(delimiter)
			name = strings.TrimSpace(name) // required to trim the \n delimiter that will be read

			newItem := menuItem{
				name:   name,
				prices: make(map[string]float64), // you can also provide size to make
			}

			menu = append(menu, newItem)

		case "q":
			// break // this will just break from `switch`, not the outer `for`
			break loop // break from the outer `for` by label

		default:
			fmt.Printf("Unknown option: %v \n", choice)
		}

		//fmt.Println(menu)
	}
}
