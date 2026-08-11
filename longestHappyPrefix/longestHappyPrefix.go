package main

import "fmt"

func longestPrefix(s string) string {
	// todo: we should comparing using a rolling hash or a KMP algorithm

	// Working solution by comparing substrings.
	// Same brute-force O(n^2) solution as in "214. Shortest Palindrome".
	// It still passes the tests in around 180+ ms
	return longestPrefix_bruteForce(s)
}

func longestPrefix_bruteForce(s string) string {
	if s == "" { // corner-case
		return ""
	}

	n := len(s)
	fmt.Printf("String length: %v \n", n)

	for i := n - 2; i >= 0; i-- { // we exclude the complete string
		// this should not require copying of the substrings, so probably should be faster as reversing every prefix?
		prefix := s[0 : i+1]
		suffix := s[n-i-1 : n]

		match := prefix == suffix

		if match {
			return prefix
		}
	}

	return ""
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := longestPrefix(s)

	fmt.Printf("Longest prefix that is also a suffix: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("level", "l")
}

func test2() {
	test("ababab", "abab") // whole string we don't count, "ababa" is not a suffix
}

func test3() {
	test("abcd", "") // no suffix == prefix
}

func test4() {
	test("a", "") // corner-case -> just 1 char -> we exclude the whole string -> not found
}

func main() {
	// 1392. Longest Happy Prefix
	test1()
	test2()
	test3()
	test4()
}
