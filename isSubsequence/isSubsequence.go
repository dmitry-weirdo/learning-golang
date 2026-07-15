package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	// for the follow-up, we can parse t into map[char] -> list(pos)
	// Then, for every char found, we can find the position of the next s[i] > currentPos
	// By using a binary search on the list/array on a list for that character.

	if len(s) < 1 {
		return true
	}

	i := 0
	j := 0

	// move t[j] until we find s[i]
	for j < len(t) && i < len(s) {
		if s[i] == t[j] {
			i++
		}

		j++
	}

	return i >= len(s)
}

func test(s, t string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Subsequence to be found: %v \n", s)
	fmt.Printf("Big string: %v \n", t)

	result := isSubsequence(s, t)

	fmt.Printf("Subsequence \"%v\" is contained in string \"%v\": %v \n", s, t, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "abc"
	t := "ahbgdc"
	expected := true

	test(s, t, expected)
}

func test2() {
	s := "axc"
	t := "ahbgdc"
	expected := false

	test(s, t, expected)
}

func test3() {
	s := ""
	t := "ahbgdc"
	expected := true

	test(s, t, expected)
}

func main() {
	// 392. Is Subsequence
	test1()
	test2()
	test3()
}
