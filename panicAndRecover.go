package main

import (
	"errors"
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

	dividend, divisor = 10, 0
	//result, err := divide1(dividend, divisor)
	result, err := divide2(dividend, divisor)

	if err != nil {
		fmt.Println(err)
		return // we won't work with result anymore
	}

	fmt.Printf("%v divided by %v is %v \n", dividend, divisor, result)
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

func divide1(l, r int) (int, error) {
	if r == 0 {
		// default return 0 (default value) for the result
		return 0, errors.New("invalid divisor: must not be zero")
	}

	return l / r, nil
}

func divide2(l, r int) (result int, err error) {
	defer func() { // handle the panic thrown
		if msg := recover(); msg != nil {
			// !!! we can access the function arguments if they're named
			result = 0
			err = fmt.Errorf("%v", msg)

			// this won't work
			//return 0, errors.New("invalid divisor: must not be zero")
		}
	}()

	return l / r, nil
}
