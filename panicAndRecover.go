package main

import (
	"fmt"
)

func main() {
	fmt.Println("main 1")

	func1()

	fmt.Println("main 2")
}

func func1() {
	defer func() {
		recoverResult := recover()

		fmt.Printf("Recover result: %v \n", recoverResult)
	}() // expression in defer must be a function call

	fmt.Println("func1 1")
	panic("panic from func1")
	fmt.Println("func1 2")
}
