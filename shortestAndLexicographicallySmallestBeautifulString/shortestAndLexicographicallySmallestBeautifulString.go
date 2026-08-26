package main

import (
	"fmt"
)

func shortestBeautifulSubstring(s string, k int) string {
	const BIGGER_THAN_MAX_LENGTH = 100_000 // len(s) <= 100
	minLength := BIGGER_THAN_MAX_LENGTH
	result := "" // if no substring found -> return empty substring

	left := 0

	freq := make([]int, 2) // just 0 and 1
	length := 0

	for right := 0; right < len(s); right++ {
		rightValue := charToInt(s[right])
		freq[rightValue]++

		for freq[1] >= k {
			length = right - left + 1
			if length < minLength { // found shorter substring
				minLength = length
				result = s[left : right+1]
			} else if length == minLength {
				if s[left:right+1] < result { // same length but lexicographically smaller
					result = s[left : right+1]
				}
			}

			leftValue := charToInt(s[left])
			freq[leftValue]--
			left++
		}
	}

	return result
}

func charToInt(ch byte) int {
	if ch == '0' {
		return 0
	} else {
		return 1
	}
}

func test(s string, k int, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)
	fmt.Printf("K (count of 1-s in the substring): %v \n", k)

	result := shortestBeautifulSubstring(s, k)

	fmt.Printf("Lexicographically smallest substring with %v 1-s: %v \n", k, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("100011001", 3, "11001")
}

func test2() {
	test("1011", 2, "11")
}

func test3() {
	test("000", 1, "") // no substring found
}

func main() {
	// 2904. Shortest and Lexicographically Smallest Beautiful String
	test1()
	test2()
	test3()
}
