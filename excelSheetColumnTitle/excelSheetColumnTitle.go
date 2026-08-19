package main

import (
	"fmt"
	"slices"
	"strings"
)

func convertToTitle(columnNumber int) string {
	// corner-case -> 0 input
	if columnNumber == 1 {
		return "A"
	}

	chars := []byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G',
		'H', 'I', 'J', 'K', 'L', 'M', 'N',
		'O', 'P', 'Q', 'R', 'S', 'T', 'U',
		'V', 'W', 'X', 'Y', 'Z',
	}

	base := len(chars)
	char := byte('A')

	var sb strings.Builder

	for columnNumber > 0 {
		// given index starts at 1, not at 0
		// and this propagates to every digit
		columnNumber--

		char = chars[columnNumber%base]
		sb.WriteByte(char)

		columnNumber /= base
	}

	return reverseString(sb.String())
}

func reverseString(s string) string {
	stringAsSlice := []byte(s)
	slices.Reverse(stringAsSlice) // reverses in place

	return string(stringAsSlice)
}

func test(x int, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Excel column number: %v \n", x)

	result := convertToTitle(x)

	fmt.Printf("Excel column name: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(1, "A")
}

func test2() {
	test(26, "Z")
}

func test3() {
	test(27, "AA")
}

func test4() {
	test(28, "AB")
}

func test5() {
	test(701, "ZY")
}

func main() {
	// 168. Excel Sheet Column Title
	test1()
	test2()
	test3()
	test4()
	test5()
}
