package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	// stack keeps the int operands
	stack := list.New()

	value := 0

	for _, s := range tokens {
		switch s {
		case "+", "-", "*", "/":
			// last operand is on the top of the stack
			// pre-last operand is second from the top of the stack
			// todo: throw if there are less than 2 elements in the stack

			operand2 := removeFromStack(stack)
			operand1 := removeFromStack(stack)

			if s == "+" {
				value = operand1 + operand2
			} else if s == "-" {
				value = operand1 - operand2
			} else if s == "*" {
				value = operand1 * operand2
			} else if s == "/" {
				value = operand1 / operand2
			}

			stack.PushFront(value)

		default:
			// todo: handle the parsing error if required
			operand, _ := strconv.Atoi(s)
			stack.PushFront(operand)
		}
	}

	// last value in the stack is the result
	return removeFromStack(stack)
}

func removeFromStack(stack *list.List) int {
	return stack.Remove(stack.Front()).(int)
}

func test(arr []string, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)

	result := evalRPN(arr)

	fmt.Printf("Evaluate result: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
		return
	}
}

func test1() {
	arr := []string{"2", "1", "+", "3", "*"}
	expected := 9

	test(arr, expected)
}

func test2() {
	arr := []string{"4", "13", "5", "/", "+"}
	expected := 6

	test(arr, expected)
}

func test3() {
	arr := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	expected := 22

	test(arr, expected)
}

func main() {
	// 150. Evaluate Reverse Polish Notation
	test1()
	test2()
	test3()

	// e.g. 6 / -132 should be 0, but will it be in Go/Java? Yes, 6 / -132 == 0
	//fmt.Printf("6 / -132: %v \n", 6/-132)
}
