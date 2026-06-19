package main

import "fmt"

func main() {
	var i myInt = 666
	even := i.isEven()

	fmt.Printf("MyInt %v is even: %v \n", i, even)
}

// method receiver cannot be a built-in type like int
//func (i int) isEven() bool {
//}

type myInt int

func (i myInt) isEven() bool {
	return i%2 == 0 // super-ugly standard formatting, no space around arithmetic operators
}
