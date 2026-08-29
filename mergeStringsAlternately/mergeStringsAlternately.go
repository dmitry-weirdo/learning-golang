package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	n := min(len(word1), len(word2))

	var sb strings.Builder

	for i := range n {
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[i])
	}

	if len(word1) > n {
		sb.WriteString(word1[n:])
	}

	if len(word2) > n {
		sb.WriteString(word2[n:])
	}

	return sb.String()
}

func test(s1, s2 string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String 1: %v \n", s1)
	fmt.Printf("String 2: %v \n", s2)

	result := mergeAlternately(s1, s2)

	fmt.Printf("Merged string: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("abc", "pqr", "apbqcr")
}

func test2() {
	test("ab", "pqrs", "apbqrs")
}

func test3() {
	test("abcd", "pq", "apbqcd")
}

func main() {
	// 1768. Merge Strings Alternately
	test1()
	test2()
	test3()
}
