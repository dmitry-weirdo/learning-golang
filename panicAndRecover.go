package main

import (
	"fmt"
)

func main() {
	fmt.Println("main 1")

	func1()

	fmt.Println("main 2")

	dividend, divisor := 10, 5
	divideResult := divide(dividend, divisor)
	fmt.Printf("%v divided by %v is %v \n", dividend, divisor, divideResult)

	// panic: runtime error: integer divide by zero
	dividend, divisor = 10, 0
	divideResult = divide(dividend, divisor)

	// todo: this will still called and printed if we recovered from an error in `divide`
	fmt.Printf("%v divided by %v is %v \n", dividend, divisor, divideResult)
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

func divide(dividend, divisor int) int {

	// anonymous function
	defer func() { // will be called also on a successful function completion
		message := recover() // nil if no panic was raised

		if message != nil {
			fmt.Printf("Recovered from: %v \n", message)
		}
	}()

	return dividend / divisor
}
