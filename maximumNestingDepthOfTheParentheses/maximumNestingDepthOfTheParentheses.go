package main

import "fmt"

func maxDepth(s string) int {
	maxBrackets := 0
	level := 0

	for _, v := range s {
		if v == '(' {
			level++

			maxBrackets = max(maxBrackets, level)
		} else if v == ')' {
			level--
		}
	}

	return maxBrackets
}

func test(s string, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := maxDepth(s)

	fmt.Printf("Max brackets depth: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("", 0)
}

func test2() {
	test("0123", 0)
}

func test3() {
	test("1+2-3", 0)
}

func test4() {
	test("(1+(2*3)+((8)/4))+1", 3)
}

func test5() {
	test("(1)+((2))+(((3)))", 3)
}

func test6() {
	test("()(())((()()))", 3)
}

func main() {
	// 1614. Maximum Nesting Depth of the Parentheses
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
