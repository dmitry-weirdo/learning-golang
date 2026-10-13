package main

import "fmt"

func minAddToMakeValid(s string) int {
	openingClosingBalance := 0
	nonPairedClosingCount := 0

	for _, ch := range s {
		if ch == '(' {
			openingClosingBalance++
		} else { // closing bracket
			if openingClosingBalance > 0 { // closing bracket with the preceding opening bracket
				openingClosingBalance--
			} else { // closing bracket without the preceding opening bracket
				nonPairedClosingCount++
			}
		}
	}

	return openingClosingBalance + nonPairedClosingCount
}

func test(s string, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := minAddToMakeValid(s)

	fmt.Printf("Min added brackets to make the string valid: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("())", 1) // (
}

func test2() {
	test("(((", 3) // )))
}

func test3() {
	// Test-case 45 / 116
	test("()))((", 4) // (())
}

func main() {
	// 921. Minimum Add to Make Parentheses Valid
	test1()
	test2()
	test3()
}
