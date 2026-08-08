package main

import "fmt"

func longestDupSubstring(s string) string {
	// Binary search for length of the substring
	// since we're searching for "minimum value satisfying the condition",
	// left will be "minimum length where repeating substring does NOT exist"

	left := 1       // even substrings of length 1 can be non-repeating, if all chars in the string are unique
	right := len(s) // Since we're searching for "minimum NOT repeating", we add 1. The sliding window for the repeating substring can have the size up to len(s) - 1.

	for left < right {
		mid := (left + right) / 2
		fmt.Printf("left: %v, right: %v, mid: %v \n", left, right, mid)

		repeatingSubstringIndex := getRepeatingSubstringIndex(s, mid)
		fmt.Printf("Index of longest repeating substring of len %v: %v \n", mid, repeatingSubstringIndex)

		if repeatingSubstringIndex == -1 { // target condition
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left == 1 { // no repeating substrings even with length 1, i.e.
		return ""
	}

	firstIndex := getRepeatingSubstringIndex(s, left-1)

	return s[firstIndex : firstIndex+left-1]
}

// todo: implement this method using the Rabin-Karp algorithm
func getRepeatingSubstringIndex(s string, substringLen int) int { // returns first index of repeating substring or -1
	// this is brute-force

	// substring to first index
	m := make(map[string]int)

	for i := 0; i <= len(s)-substringLen; i++ {
		substring := s[i : i+substringLen]

		if firstIndex, ok := m[substring]; ok { // substring already present -> return the index
			return firstIndex
		}

		// put first index of the new substring to the map
		m[substring] = i
	}

	return -1
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := longestDupSubstring(s)

	fmt.Printf("Longest repeating substring: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("banana", "ana")
}

func test2() {
	test("abcd", "")
}

func test3() {
	test("aa", "a")
}

func test4() {
	// repeating substrings can have length be up to (len(s) -1).
	test(
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	)

}

func main() {
	// 1044. Longest Duplicate Substring

	// to study Rabin-Karp algorithm -> find matching substrings when sliding window on the parent string
	test1()
	test2()
	test3()
	test4()
}
