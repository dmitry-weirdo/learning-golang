package main

import "fmt"

func reverseVowels(s string) string {
	// just 2 pointers from both ends, replacing these 2 vowels, until the pointers meet
	b := []byte(s)

	n := len(b)

	left := 0
	right := n - 1

	for left < right {
		// skip to vowel from left
		for (left < n) && !isVowel(b[left]) {
			left++
		}

		// skip to vowel from right
		for (right >= 0) && !isVowel(b[right]) {
			right--
		}

		if left < right { // if left went over right, do not swap
			b[left], b[right] = b[right], b[left]

			// skip the swapped characters
			left++
			right--
		}
	}

	return string(b)
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true

	default:
		return false
	}
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := reverseVowels(s)

	fmt.Printf("String with reversed vowels: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("IceCreAm", "AceCreIm")
}

func test2() {
	test("leetcode", "leotcede")
}

func main() {
	// 345. Reverse Vowels of a String
	test1()
	test2()
}
