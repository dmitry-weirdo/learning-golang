package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	// count the frequencies in a string
	freq := lowercaseEnglishLettersStringToFrequencies(magazine)

	// For every char in the node, subtract the frequency of the character.
	// If any char goes below 0 frequency, return false
	for _, ch := range ransomNote {
		freq[ch-'a']--

		if freq[ch-'a'] < 0 {
			return false
		}
	}

	return true
}

func lowercaseEnglishLettersStringToFrequencies(s string) []int {
	freq := make([]int, 26)

	for _, ch := range s {
		freq[ch-'a']++
	}

	return freq
}

func test(ransomNote, magazine string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Ransom note string: %v \n", ransomNote)
	fmt.Printf("Magazine string: %v \n", magazine)

	result := canConstruct(ransomNote, magazine)

	fmt.Printf("Ransom note can be constructed from magazine: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("a", "b", false)
}

func test2() {
	test("aa", "ab", false)
}

func test3() {
	test("aa", "aab", true)
}

func main() {
	// 383. Ransom Note
	test1()
	test2()
	test3()
}
