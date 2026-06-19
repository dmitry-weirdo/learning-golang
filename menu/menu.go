package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

const DELIMITER = '\n'

type menuItem struct {
	name   string
	prices map[string]float64
}

// do NOT use the = sign here
// type menu = []menuItem
type menu []menuItem

func (m menu) print() {
	for _, item := range m { // ignore the array index
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))

		// iterate the prices
		for size, price := range item.prices {
			fmt.Printf("\t%10s: %10.2f\n", size, price) // %10 is "fill with spaces up to 10 characters"
		}
	}
}

func (m *menu) add() {
	fmt.Println("Please enter the name of the new item")

	name, _ := in.ReadString(DELIMITER)
	name = strings.TrimSpace(name) // required to trim the \n delimiter that will be read

	newItem := menuItem{
		name:   name,
		prices: make(map[string]float64), // you can also provide size to make
	}

	*m = append(*m, newItem)
}

func Print() {
	data.print()
}

func Add() {
	data.add()
}
