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
}

func myFunc(name string, otherName *string) {
	name = "Name updated" // we can override arguments :(
	*otherName = "Other name updated"

	fmt.Printf("Name in myFunc: %v \n", name)
	fmt.Printf("OtherName in myFunc: %v \n", *otherName)
}
