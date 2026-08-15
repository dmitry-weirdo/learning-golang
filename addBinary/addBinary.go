package main

import (
	"fmt"
	"slices"
	"strings"
)

func addBinary(a string, b string) string {
	// implement as parsing bits, adding int values ant returning the string of bits
	// (this might not work since the strings can be of 10^4 length
	// Bits level comparison summing is actually a separate task "371. Sum of Two Integers".

	return addBinary_naive(a, b)
}

func addBinary_naive(a string, b string) string {
	long, short := getLongAndShortStrings(a, b)

	var sb strings.Builder

	j := len(short) - 1 // index in short

	longDigit := 0
	shortDigit := 0

	overflow := 0

	for i := len(long) - 1; i >= 0; i-- {
		if long[i] == '0' {
			longDigit = 0
		} else {
			longDigit = 1
		}

		if j >= 0 { // there are still digits in shorter string
			if short[j] == '0' {
				shortDigit = 0
			} else {
				shortDigit = 1
			}
		} else { // no more digits from the short string -> assume its bits as 0
			shortDigit = 0
		}

		// !!! we can have more than 1 as overflow to the next digit
		overflow += longDigit
		overflow += shortDigit

		// current digit
		if overflow%2 == 0 {
			sb.WriteByte('0')
		} else {
			sb.WriteByte('1')
		}

		// (overflow / 2) carries to the next digit
		overflow /= 2

		// also go down in the short string
		j--
	}

	// add the remaining overflow if there is one
	if overflow > 0 {
		sb.WriteByte('1')
	}

	// we're calculating from the last digit to the first, so the result must be reversed.
	s := sb.String()
	return reverseString(s)
}

func getLongAndShortStrings(a, b string) (long, short string) {
	if len(a) > len(b) {
		return a, b
	}

	return b, a
}

func reverseString(s string) string {
	stringAsSlice := []byte(s)
	slices.Reverse(stringAsSlice) // reverses in place

	return string(stringAsSlice)
}

func test(a, b string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Binary string A: %v \n", a)
	fmt.Printf("Binary string B: %v \n", b)

	result := addBinary(a, b)

	fmt.Printf("Binary sum of A + B: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("11", "1", "100") // 3 + 1 = 4
}

func test2() {
	test("1010", "1011", "10101") // 10 + 11 = 21
}

func test3() {
	test("11", "11", "110") // 3 + 3 = 6
}

func test4() {
	test("1111101000", "1111101000", "11111010000") // 1000 + 1000 = 2000
}

func main() {
	// 67. Add Binary
	test1()
	test2()
	test3()
	test4()
}
