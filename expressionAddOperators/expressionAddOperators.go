package main

import (
	"fmt"
	"strconv"
)

var result []string
var s string
var t int

func addOperators(num string, target int) []string {
	// save constants to global values to not pass as the stack arguments
	result = make([]string, 0)
	s = num
	t = target

	dfs(0, 0, 0, "")

	return result
}

// todo: current value should be int64 to avoid overflow
func dfs(pos int, previousOperand int, currentValue int, expression string) {
	fmt.Println()
	fmt.Printf("dfs(pos = %v, previousOperand = %v, currentValue = %v, expression = \"%v\") \n", pos, previousOperand, currentValue, expression)

	// reached end of the sting -> check if current value is equal to the target
	if pos >= len(s) {
		fmt.Printf("Reached the end of the string \"%v\". Current value: %v, expression: \"%v\" \n", s, currentValue, expression)

		if currentValue == t {
			fmt.Printf("Adding expression \"%v\" to the result. \n", expression)
			result = append(result, expression)
		}

		return
	}

	// try all possible lengths of the string starting with the current pos
	for endIndex := pos + 1; endIndex <= len(s); endIndex++ {
		operandString := s[pos:endIndex]

		fmt.Printf("Pos %v, possible operand starting with this pos: %v \n", pos, operandString)

		if (len(operandString) > 1) && (operandString[0] == '0') { // not just "0" but starts with 0 -> stop further iterations
			fmt.Printf("Operand string \"%v\" starts with 0. Stopping further iterations from pos %v. \n", operandString, pos)
			break
		}

		// this is O(n), we need to procceed the complete string
		operandInt, _ := strconv.Atoi(operandString) // error should never happen

		// todo: exclude handing for pos == 0?

		var newIndex = endIndex
		var newPrev int
		var newValue int
		var newExpression string

		// adding an expression (i.e. string concatenation) is O(n), it will add a new string
		// todo: we could add strings.Builder
		if pos == 0 {
			// since we don't add any operands before position 0, don't try any operators before it
			newPrev = operandInt
			newValue = operandInt         // no operations, just take the current value
			newExpression = operandString // expression will be just "", so we just include the current operand

			dfs(newIndex, newPrev, newValue, newExpression)
		} else {
			// try addition
			newPrev = operandInt
			newValue = currentValue + operandInt
			newExpression = expression + "+" + operandString
			dfs(newIndex, newPrev, newValue, newExpression)

			// try subtraction
			newPrev = -operandInt
			newValue = currentValue - operandInt
			newExpression = expression + "-" + operandString
			dfs(newIndex, newPrev, newValue, newExpression)

			// try multiplication
			newPrev = previousOperand * operandInt // we multiply the previous operand and the current operand

			// we need to negate the previous operation, since multiplication takes precedence
			// tricky: (currentValue - previousOperand) negates the previous operation

			// explanation for the case if previous operation was also a multiplication:
			// previousOperand will be the multiplication of the previous arguments (a1 * a2)
			// so we first subtract the previous multiplication: -(a1 * a2)
			// and then add the new multiplication of all the operands including the new one: (a1 * a2 * a3)
			newValue = (currentValue - previousOperand) + previousOperand*operandInt

			newExpression = expression + "*" + operandString
			dfs(newIndex, newPrev, newValue, newExpression)
		}
	}
}

func test(s string, target int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: \"%v\", target values = %v \n", s, target)

	solutions := addOperators(s, target)

	fmt.Println()
	fmt.Printf("Possible solutions of converting string \"%v\" into target %v:\n%v", s, target, solutions)
}

func test1() {
	s = "123"
	target := 6

	test(s, target)
}

func test2() {
	s = "105"
	target := 5

	test(s, target)
}

func test3() {
	s = "000"
	target := 0

	test(s, target)
}

func test4() {
	s = "00"
	target := 0

	test(s, target)
}

func test5() {
	s = "0"
	target := 0

	test(s, target)
}

func main() {
	// 282. Expression Add Operators
	test1()
	test2()
	test3()
	test4()
	test5()
}
