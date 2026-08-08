package main

import (
	"fmt"
	"strings"
)

func maxRepeating(sequence string, word string) int {
	k := 0

	s := word

	for strings.Index(sequence, s) >= 0 {
		s = s + word
		k++
	}

	return k
}

func test(s, word string, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)
	fmt.Printf("Word string: %v \n", s)

	result := maxRepeating(s, word)

	fmt.Printf("Max consecutive repeats of substring \"%v\" in string \"%v\": %v \n", word, s, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("ababc", "ab", 2)
}

func test2() {
	test("ababc", "ba", 1)
}

func test3() {
	test("ababc", "ac", 0)
}

func main() {
	// 1668. Maximum Repeating Substring
	test1()
	test2()
	test3()
}
