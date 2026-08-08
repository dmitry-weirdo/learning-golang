package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	// If the string consists of a pattern, it looks like:
	// s = p1 + p2 + ... + pLast
	//
	// Let's concatenate string with self: s + s.
	// s + s = p1 + p2 + ... + pLast + (p1 + p2 + ... pLast).
	//
	// Let's shift from the start of the string by P.
	// s + s = p1 + (p2 + ... pLast + p1) + (p2 + ... pLast)
	//
	// and (p2 + ... pLast + p1) - is equal to the S itself.
	// So if matching of S to itself in (S + S) from index 1 happens before len(s), S is a divisible to a pattern.
	// If it happened only on len(s), the only pattern of S is S itself, so we return false.
	//
	// Basically, for the pattern, we have to divide a string at least by 2.
	// So the match must happen before len(s) / 2.

	t := s + s

	index := strings.Index(t[1:], s) // search for S in S + S, starting with index 1 (skip the beginning match from index 0)

	return (index + 1) < len(s)
}

func test(s string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := repeatedSubstringPattern(s)

	fmt.Printf("String is divisible into a pattern: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("abab", true) // pattern is "ab"
}

func test2() {
	test("aba", false)
}

func test3() {
	test("abcabcabcabc", true) // pattern is "abc" or "abcabc"
}

func test4() {
	test("abcabcabc", true) // pattern is "abc"
}

func test5() {
	test("a", false) // cannot split into _multiple_ patterns
}

func main() {
	// 459. Repeated Substring Pattern
	test1()
	test2()
	test3()
	test4()
	test5()
}
