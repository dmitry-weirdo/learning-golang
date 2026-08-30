package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// todo: for non-just-ascii characters, use a hashMap with rune key
	freq := make([]int, 26)

	for _, v := range s {
		freq[v-'a']++
	}

	for _, v := range t {
		freq[v-'a']--
	}

	// if any of the diffs != 0, we return false
	for _, v := range freq {
		if v != 0 {
			return false
		}
	}

	return true
}

func test(s, t string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("S - String 1: %v \n", s)
	fmt.Printf("T - String 2: %v \n", t)

	result := isAnagram(s, t) // todo: replace with your function

	fmt.Printf("S is anagram of T: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("anagram", "nagaram", true)
}

func test2() {
	test("rat", "car", false)
}

func test3() {
	test("ab", "a", false)
}

func test4() {
	test("a", "ab", false)
}

func main() {
	// 242. Valid Anagram
	test1()
	test2()
	test3()
	test4()
}
