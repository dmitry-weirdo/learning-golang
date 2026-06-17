package main

import (
	"fmt"
	"strings"
)

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

	var i int32 = 32
	var f float32

	// f = i // implicit conversion will NOT work
	f = float32(i) // explicit conversion

	fmt.Printf("Int64 value: %d, converted to float32 value: %f \n", i, f)

	// overflow conversion test
	// yes, you can use built-in type names as variable names - it's just a warning
	var int16 int16 = 32767
	var int8 int8 = int8(int16) // won't fail, will be set to -1
	fmt.Printf("Int16 value: %d, converted to int8 value: %d \n", int16, int8)

	// multiple var declaration in one line
	// formatting with %s will format to string with writing the variable type
	a, b, c := 10, 5.3, "string"
	fmt.Printf("a: %s, b: %s, c: %s \n", a, b, c)

	// String comparison
	str1 := "String value äöü"
	str2 := "String VaLuE ÄÖÜ"

	caseSensitiveEquals := (str1 == str2)
	caseInsensitiveEquals := strings.EqualFold(str1, str2) // can convert UTF-8 characters

	fmt.Println()
	fmt.Printf("str1: \"%s\", str2: \"%s\", case-sensitive equals: %t \n", str1, str2, caseSensitiveEquals)
	fmt.Printf("str1: \"%s\", str2: \"%s\", case-insensitive equals: %t \n", str1, str2, caseInsensitiveEquals)

	// constants
	const aConst = 42

	const bConst string = "constant string value"
	const b2Const string = `Raw string constant \n won't be escaped.'`
	const b3Const = bConst

	// multiple constants declaration in a group
	const (
		boolConst = true
		eConst    = 3.14
	)

	const (
		q = "foo"
		w // unassigned constant receives the previous value
		e
		ec = "zoo"
		fc
	)

	// iota
	const ii = iota // will be 0

	const (
		int0 = iota     // iota = 0
		int1            // iota = 1
		int2 = 3 * iota // iota = 2
	)

	const (
		newInt0 = iota // iota resets in a new constant block
	)

	const (
		_    = iota // 0 ignored
		ONE         // 1
		TWO         // 2
		_           // 3 ignored
		FOUR        // 4
	)

	const (
		ten0 = iota * 10 // set step to 10
		ten1             // 1 * 10 = 10
		ten2             // 2 * 10 = 20
	)

	// use iota for bit flags
	const (
		Read    = 1 << iota // 0001
		Write               // 0010
		Execute             // 0100
		Delete              // 1000
	)
}
