package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")

	in := bufio.NewReader(os.Stdin)

	const delimiter = '\n'
	choice, _ := in.ReadString(delimiter)
	choice = strings.TrimSpace(choice)

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

	fmt.Println(menu)
}
