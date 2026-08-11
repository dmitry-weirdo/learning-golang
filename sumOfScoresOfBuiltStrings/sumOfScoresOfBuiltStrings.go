package main

import "fmt"

func sumScores(s string) int64 {
	// todo: write an O(n) solution using Z-algorithm and KMP-algorithm

	// Brute-force O(n^2) like in "214. Shortest Palindrome".
	// However, since we don't break fast on finding 1 solution,
	// it's falling on TLE.
	// Test-case 43 / 150 testcases passed
	// String length: 39239
	return sumScores_bruteForce(s)
}

func sumScores_bruteForce(s string) int64 {
	if s == "" { // corner-case
		return 0
	}

	n := len(s)
	fmt.Printf("String length: %v \n", n)

	result := int64(0)

	for i := n - 1; i >= 0; i-- { // we include the whole string
		// this should not require copying of the substrings, so probably should be faster as reversing every prefix?
		suffix := s[n-i-1 : n] // build suffix substring

		// we're comparing all the prefix with length up to i - 1
		for j := len(suffix) - 1; j >= 0; j-- {
			suffixPrefix := suffix[0 : j+1]
			prefix := s[0 : j+1]

			if prefix == suffixPrefix {
				//fmt.Printf("Suffix string = \"%v\". Longest common prefix = \"%v\". Adding its length = %v to the result. \n", suffix, prefix, j + 1)

				result += int64(j + 1)
				break
			}
		}
	}

	return result
}

func test(s string, expectedResult int64) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := sumScores(s)

	fmt.Printf("Sum of longest prefixes lengths between original string and all suffix substrings: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	// "b" -> prefix "b" -> 1
	// "bab" -> prefix "bab" -> 3
	// "babab" -> prefix "babab" -> 5
	test("babab", 9)
}

func test2() {
	// "az" -> prefix "az" -> 2
	// "azbzaz" -> prefix "azb" -> 3
	// "azbazbzaz" -> prefix "azbazbzaz" -> 9
	test("azbazbzaz", 14)
}

func main() {
	// 2223. Sum of Scores of Built Strings
	test1()
	test2()
}
