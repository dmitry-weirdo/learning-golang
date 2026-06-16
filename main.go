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

	// error is an interface and it can be nil
	var x error = nil
	fmt.Println(x) // will print <nil>

	// !!! unused variables won't even compile!
	// different ways to declare a variable
	println("\nDifferent way of declaring string variables:")

	var s1 string // declare, this will be an empty string, NOT nil
	println(s1)

	var s2 string = "initial value" // declare and initialize
	println(s2)

	var s3 = "inferred type of string" // init with inferred type
	println(s3)

	s4 := "short declaration string" // short declaration syntax
	println(s4)

	var i1 int // will be 0
	println(i1)
}
