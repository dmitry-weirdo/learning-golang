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

func Print() {
	for _, item := range menu { // ignore the array index
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))

		// iterate the prices
		for size, price := range item.prices {
			fmt.Printf("\t%10s: %10.2f\n", size, price) // %10 is "fill with spaces up to 10 characters"
		}
	}
}

func Add() {
	fmt.Println("Please enter the name of the new item")

	name, _ := in.ReadString(DELIMITER)
	name = strings.TrimSpace(name) // required to trim the \n delimiter that will be read

	newItem := menuItem{
		name:   name,
		prices: make(map[string]float64), // you can also provide size to make
	}

	menu = append(menu, newItem)
}
