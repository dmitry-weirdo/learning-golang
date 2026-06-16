package main

import "fmt"

func main() {
	fmt.Println("Hello, Gophers!")

	fmt.Println("\nInterpreted string:\n - This will be at the new line.\n")
	fmt.Println(`Raw string:\n - This will NOT be at the new line, "\n" is not interpreted.`)

	fmt.Println(`
Raw string line one
Raw string line two
Raw string line three.`)

	// try Unicode
	unicodeInline := "Hello, 世界"
	fmt.Println(unicodeInline)

	unicodeAsCodes := "\U0001F600" // 😀
	fmt.Println(unicodeAsCodes)
}
