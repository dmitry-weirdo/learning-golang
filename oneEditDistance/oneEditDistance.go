package main

import (
	"fmt"
)

func isOneEditDistance(s string, t string) bool {
	var short, long string

	if len(s) <= len(t) {
		short = s
		long = t
	} else {
		short = t
		long = s
	}

	if len(long)-len(short) >= 2 { // length diff of 2 and more is more than 1 edit distance
		return false
	}

	// if same lengths, we need to replace exactly 1 character

	// if len diff is 1, one edit must be the insertion into short

	left := 0
	right := 0

	// update left and right to the first diff
	for (left < len(short)) && (short[left] == long[right]) {
		left++
		right++
	}

	// skip the difference
	if len(short) == len(long) {
		if left == len(short) {
			// reached the end of the strings -> strings are equal -> 0 edit distance -> return false
			return false
		}

		// skip a char in both strings
		return short[left+1:] == long[right+1:]
	} else { // len(short) = len(long) - 1
		// skip the character in the long string
		return short[left:] == long[right+1:]
	}
}

func isOneEditDistanceNaive(s string, t string) bool {
	var short, long string

	if len(s) <= len(t) {
		short = s
		long = t
	} else {
		short = t
		long = s
	}

	if len(long)-len(short) >= 2 { // length diff of 2 and more is more than 1 edit distance
		return false
	}

	// if same lengths, we need to replace exactly 1 character

	// if len diff is 1, one edit must be the insertion into short

	left := 0
	right := 0

	// update left and right to the first diff
	for (left < len(short)) && (short[left] == long[right]) {
		left++
		right++
	}

	// todo: we can just compare the remaining substrings :)
	// skip the difference
	if len(short) == len(long) {
		if left == len(short) {
			// reached the end of the strings -> strings are equal -> 0 edit distance -> return false
			return false
		}

		// skip a char in both strings
		left++
		right++

		for (left < len(short)) && (short[left] == long[right]) {
			left++
			right++
		}

		if left != len(short) { // found the 2nd diff
			return false
		}

		// reached the end of the strings and not found the 2nd diff -> success
		return true
	} else { // len(short) = len(long) - 1
		// skip the character in the long string
		right++

		for (left < len(short)) && (short[left] == long[right]) {
			left++
			right++
		}

		if left != len(short) { // found the 2nd diff
			return false
		}

		// reached the end of the strings and not found the 2nd diff -> success
		return true
	}
}

func test(s, t string, expectedResult bool) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("S (first string): %v \n", s)
	fmt.Printf("T (second string): %v \n", t)

	result := isOneEditDistance(s, t)

	fmt.Printf("String \"%v\" is exactly within 1 edit distance of string \"%v\": %v \n", s, t, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "ab"
	t := "abc"
	expected := true // insertion of "c" is one edit operation

	test(s, t, expected)
}

func test2() {
	s := ""
	t := ""
	expected := false // 0 edit distance < 1 edit distance

	test(s, t, expected)
}

func main() {
	// 161. One Edit Distance
	test1()
	test2()
}
