package main

import (
	"container/list"
	"fmt"
)

func calculate(s string) int {
	v := 0 // current operand

	previousOperator := '+'

	// stack contains just the values
	stack := list.New()

	for i, ch := range s {
		fmt.Printf("char: %c, v: %d, previous operator: %c \n", ch, v, previousOperator)

		if '0' <= ch && ch <= '9' { // collect digits of the current operand
			v = v*10 + int(ch-'0')

			fmt.Printf("Updated current operand to %v \n", v)
		}

		// reached end of the operand -> make a decision based on the previous operand
		if i >= len(s)-1 || // also proceed with the operand end when we're at the end of the string
			ch == '+' ||
			ch == '-' ||
			ch == '*' ||
			ch == '/' {

			if previousOperator == '+' {
				fmt.Printf("Reached operator \"%c\". Previous operator = \"%c\". Pushed operand %v to the stack. \n", ch, previousOperator, v)
				stack.PushBack(v)
			} else if previousOperator == '-' {
				fmt.Printf("Reached operator \"%c\". Previous operator = \"%c\". Pushed operand %v to the stack. \n", ch, previousOperator, -v)
				stack.PushBack(-v)
			} else if previousOperator == '*' {
				fmt.Printf("Reached operator \"%c\". Previous operator = \"%c\".\n", ch, previousOperator)

				previousOperand := stack.Remove(stack.Back()).(int)
				fmt.Printf("Extracted previous operand = %v from the stack. \n", previousOperand)

				newOperand := previousOperand * v
				stack.PushBack(newOperand)

				fmt.Printf("Added new operand = %v * %v = %v from the stack. \n", previousOperand, v, newOperand)
			} else if previousOperator == '/' {
				fmt.Printf("Reached operator \"%c\". Previous operator = \"%c\".\n", ch, previousOperator)

				previousOperand := stack.Remove(stack.Back()).(int)
				fmt.Printf("Extracted previous operand = %v from the stack. \n", previousOperand)

				newOperand := previousOperand / v
				stack.PushBack(newOperand)

				fmt.Printf("Added new operand = %v / %v = %v from the stack. \n", previousOperand, v, newOperand)
			}

			// this will be set to the last character in case of string end,
			// but it doesn't matter since we're not handling the string any further
			previousOperator = ch

			// reset the operand
			// this will be set to 0 in case of string end,
			// but it doesn't matter since we're not handling the string any further
			v = 0
		}
	}

	// just sum the values. The operands after - (minus) sign will be stored as negative values
	sum := 0

	for e := stack.Front(); e != nil; e = e.Next() {
		sum += e.Value.(int)
	}

	return sum
}

func test(s string, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := calculate(s)

	fmt.Printf("Result: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "3+2*2"
	expected := 7

	test(s, expected)
}

func test2() {
	s := " 3+5 / 2 "
	expected := 5

	test(s, expected)
}

func main() {
	// 227. Basic Calculator II
	test1()
	test2()
}
