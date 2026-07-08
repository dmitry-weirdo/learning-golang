package main

import (
	"fmt"
)

func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] == s[right] { // happy-case -> go to next characters
			left++
			right--
			continue
		}

		// mismatch found -> try skipping from left or from right
		palindromeSkippingLeft := isPalindrome(s, left+1, right)
		if palindromeSkippingLeft { // do not check skipping from right if skipping from left is already a valid palindrome
			return true
		}

		palindromeSkippingRight := isPalindrome(s, left, right-1)
		return palindromeSkippingRight
	}

	// no mismatches found -> is a palindrome
	return true
}

func isPalindrome(s string, left, right int) bool {
	i := left
	j := right

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func test(s string, expectedResult bool) {
	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("String: %v \n", s)

	result := validPalindrome(s)

	fmt.Printf("Valid palindrome (can skip 1 character): %v  \n", result)
	fmt.Printf("Expected result: %v  \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "aba"
	expected := true

	test(s, expected)
}

func test2() {
	s := "abca"
	expected := true

	test(s, expected)
}

func test3() {
	s := "abc"
	expected := false

	test(s, expected)
}

func main() {
	//test1()
	//test2()
	test3()
}
