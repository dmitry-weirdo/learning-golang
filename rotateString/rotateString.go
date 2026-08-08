package main

import (
	"fmt"
	"strings"
)

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) { // safety measure, substring still can be found
		return false
	}

	// All the rotations of S will be present in S + S.
	// "abcde" -> "abcdeabcde" - will contain all rotations "bcdea", "cdeab", "deabc", "eabcd", "abcde"
	t := s + s

	return strings.Index(t, goal) >= 0
}

func test(s, r string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Original string: %v \n", s)
	fmt.Printf("Potential rotated string: %v \n", r)

	result := rotateString(s, r)

	fmt.Printf("\"%v\" can be rotated to \"%v\": %v \n", s, r, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("abcde", "cdeab", true)
}

func test2() {
	test("abcde", "abcde", true)
}

func test3() {
	test("abcde", "abc", false)
}

func test4() {
	test("abcde", "abcdea", false)
}

func test5() {
	test("abcde", "abced", false)
}

func test6() {
	test("a", "a", true)
}

func test7() {
	test("a", "b", false)
}

func main() {
	// 796. Rotate String
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
}
