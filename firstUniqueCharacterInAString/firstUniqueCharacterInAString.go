package main

import (
	"fmt"
)

// for just lower-case english chars, we can use an array of 26 elements instead of map
func firstUniqChar(s string) int {
	m := make([]int, 26)

	// todo: to avoid rune and use just bytes, we can iterate as s[i]
	for _, ch := range s {
		index := getIndex(ch)
		m[index]++
	}

	for i, ch := range s {
		index := getIndex(ch)
		if m[index] == 1 {
			return i
		}
	}

	return -1
}

func getIndex(ch rune) int {
	return int(ch - 'a')
}

func firstUniqCharMap(s string) int {
	m := make(map[rune]int)

	for _, ch := range s {
		if _, ok := m[ch]; ok {
			m[ch]++
		} else {
			m[ch] = 1
		}
	}

	for i, ch := range s {
		if m[ch] == 1 {
			return i
		}
	}

	return -1
}

func test(s string, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("s: %v \n", s)

	result := firstUniqChar(s)

	fmt.Printf("First unique character index: %v  \n", result)
	fmt.Printf("Expected result: %v  \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "neetcodeislove"
	expected := 0

	test(s, expected)
}

func test2() {
	s := "baab"
	expected := -1

	test(s, expected)
}

func test3() {
	s := "neetcodeneet"
	expected := 4

	test(s, expected)
}

func main() {
	test1()
	test2()
	test3()
}
