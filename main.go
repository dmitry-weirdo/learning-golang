package main

import "fmt"

func main() {
	fmt.Println("Hello, Gophers!")

	fmt.Println("\nInterpreted string:\n - This will be at the new line.\n")
	fmt.Println(`Raw string:\n - This will NOT be at the new line, "\n" is not interpreted.`)
}
