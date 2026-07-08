package main

import "fmt"

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		// skip non-letters from left
		for (i < len(s)) && !isAlphanumeric(s[i]) {
			i++
		}

		// skip non-letters from right
		for (j >= 0) && !isAlphanumeric(s[j]) {
			j--
		}

		if i >= j { // if we went out of bounds of the string, it's a palindrome
			return true
		}

		if !letterEqualCaseInsensitive(s[i], s[j]) {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphanumeric(c byte) bool {
	return ('A' <= c && c <= 'Z') ||
		('a' <= c && c <= 'z') ||
		('0' <= c && c <= '9')
}

func letterEqualCaseInsensitive(a byte, b byte) bool {
	return toLowerCase(a) == toLowerCase(b)
}

func toLowerCase(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c + 'a' - 'A'
	}

	return c
}

func test(s string, expectedResult bool) {
	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("String: %v \n", s)

	result := isPalindrome(s)

	fmt.Printf("String is a palindrome: %v  \n", result)
	fmt.Printf("Expected result: %v  \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := ".,"
	expected := true

	test(s, expected)
}

func test2() {
	s := "Was it a car or a cat I saw?"
	expected := true

	test(s, expected)
}

func test3() {
	s := "tab a cat"
	expected := false

	test(s, expected)
}

func test4() {
	s := "0aa"
	expected := false

	test(s, expected)
}

func test5() {
	s := "01a10"
	expected := true

	test(s, expected)
}

func test6() {
	s := "0"
	expected := true

	test(s, expected)
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
