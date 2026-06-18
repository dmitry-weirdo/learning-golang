package main

import "fmt"

func main() {
	fmt.Println("main 1")

	// will defer to after the execution of main
	defer fmt.Println("deferred 1")

	fmt.Println("main 2")

	// deferred functions are LIFO, so "deferred 2" will be executed before "deferred 1"
	defer fmt.Println("deferred 2")
}
