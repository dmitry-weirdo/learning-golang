package main

import "fmt"

func main() {

	name, otherName := "Name", "Other name"
	fmt.Printf("Name before function call: %v \n", name)
	fmt.Printf("OtherName before function call: %v \n", otherName)

	const nameUnchangeable = "constant name"
	// &nameUnchangeable // you cannot pass constant by reference, even get its address

	myFunc(name, &otherName)

	fmt.Printf("Name after function call: %v \n", name)
	fmt.Printf("OtherName after function call: %v \n", otherName)

	handleDivideSafe(10, 3)
	handleDivideSafe(12, 0)
}

func myFunc(name string, otherName *string) {
	name = "Name updated" // we can override arguments :(
	*otherName = "Other name updated"

	fmt.Printf("Name in myFunc: %v \n", name)
	fmt.Printf("OtherName in myFunc: %v \n", *otherName)
}

func handleDivideSafe(divided int, divisor int) {
	result, ok := divideSafe2(divided, divisor)

	if ok {
		fmt.Printf("%v / %v = %v.\n", divided, divisor, result)
	} else {
		fmt.Printf("Dividing %v / %v failed.\n", divided, divisor)
	}
}

func divideSafe(l, r int) (int, bool) { // multiple return types
	if r == 0 {
		return 0, false
	}

	return l / r, true
}

func divideSafe2(l, r int) (result int, ok bool) { // will default to 0 and false
	if r == 0 {
		return // "naked return" will return the default value
	}

	result = l / r
	ok = true // if not set, the default will be used -> very inexplicit
	return    // "naked return" will return the updated values
}
