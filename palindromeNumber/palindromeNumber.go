package main

import (
	"fmt"
	"slices"
	"strconv"
)

func isPalindrome(x int) bool {
	return isPalindrome_number(x) // naive, works perfectly, O(log10 x) space for string
	//return isPalindrome_string(x) // naive, works perfectly, O(log10 x) space for string
}

func isPalindrome_number(x int) bool {
	if x < 0 { // negative numbers are NOT a palindrome
		return false
	}

	if x == 0 {
		return true
	}

	if x%10 == 0 { // ends with 0 and is not 0 -> not a palindrome, since 0 at the start would be required
		return false
	}

	reverted := 0

	for x > reverted { // we're reverting until there are more or same amount of digits in the remaining number then in the reverted
		lastDigit := x % 10
		reverted = reverted*10 + lastDigit

		x = x / 10
	}

	// for odd digits, reverted will be one number (one last digit) longer
	return (x == reverted) || x == (reverted/10)
}

func isPalindrome_string(x int) bool {
	s := strconv.Itoa(x)
	sb := []byte(s)

	slices.Reverse(sb)

	s2 := string(sb)

	return s == s2
}

func test(x int, expectedResult bool) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Number: %v \n", x)

	result := isPalindrome(x)

	fmt.Printf("%v is a palindrome number: %v \n", x, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(121, true)
}

func test2() {
	test(-121, false)
}

func test3() {
	test(10, false)
}

func test4() {
	test(0, true)
}

func test5() {
	test(1221, true)
}

func main() {
	// 9. Palindrome number
	test1()
	test2()
	test3()
	test4()
	test5()
}
