package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	// 0 - qwertyuiop
	// 1 - asdfghjkl
	// 2 - zxcvbnm
	rows := []int{
		1, // a
		2, // b
		2, // c
		1, // d
		0, // e
		1, // f
		1, // g
		1, // h
		0, // i
		1, // j
		1, // k
		1, // l
		2, // m
		2, // n
		0, // o
		0, // p
		0, // q
		0, // r
		1, // s
		0, // t
		0, // u
		2, // v
		0, // w
		2, // x
		0, // y
		2, // z
	}

	result := make([]string, 0)

	for _, v := range words {
		s := strings.ToLower(v)

		mask := 0
		noRepeats := true

		for _, ch := range s {
			row := rows[ch-'a']

			maskBit := 1 << row

			if maskBit|mask != maskBit {
				noRepeats = false
				break
			}

			mask = maskBit
		}

		if noRepeats {
			result = append(result, v)
		}
	}

	return result
}

func test(arr []string, expectedResult []string) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := findWords(arr)

	fmt.Printf("Words that can be typed with one keyboard row (ignore case): %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	test(
		[]string{"Hello", "Alaska", "Dad", "Peace"},
		[]string{"Alaska", "Dad"},
	)
}

func test2() {
	test(
		[]string{"omk"},
		[]string{},
	)
}

func test3() {
	test(
		[]string{"adsdf", "sfd"},
		[]string{"adsdf", "sfd"},
	)
}

func main() {
	// 500. Keyboard Row
	test1()
	test2()
	test3()
}
