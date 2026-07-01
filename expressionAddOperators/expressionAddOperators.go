package main

import "fmt"

var result []string
var s string
var t int

func addOperators(num string, target int) []string {
	// save constants to global values to not pass as the stack arguments
	s = num
	t = target

	dfs(0, 0, 0, "")

	return result
}

// todo: current value should be int64 to avoid overflow
func dfs(pos int, prev int, currentValue int, expression string) {
	fmt.Println()
	fmt.Printf("dfs(pos = %v, prev = %v, currentValue = %v, expression = \"%v\") \n", pos, prev, currentValue, expression)

	// reached end of the sting -> check if current value is equal to the target
	if pos >= len(s) {
		fmt.Printf("Reached and of the string \"%v\". Current value: %v, expression: %v", s, currentValue, expression)

		if currentValue == t {
			fmt.Printf("Adding expression %v to the result", expression)
			result = append(result, expression)
		}

		return
	}

	// try all possible lengths of the string starting with the current pos
	for endIndex := pos + 1; endIndex <= len(s); endIndex++ {
		operandString := s[pos:endIndex]

		fmt.Printf("Pos %v, possible operand starting with this pos: %v \n", pos, operandString)

		if (len(operandString) > 0) && (s[0] == '0') { // not just "0" but starts with 0 -> stop further iterations
			fmt.Printf("Operand string \"%v\" starts with 0. Stopping further iterations from pos %v. \n", operandString, pos)
			break
		}
	}

}

func test(s string, target int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: \"%v\", target values = %v \n", s, target)

	result := addOperators(s, target)

	fmt.Println()
	fmt.Printf("Possible solutions of converting string \"%v\" into target %v:\n%v", s, target, result)
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

func main() {
	test1()
	test2()
	test3()
}
