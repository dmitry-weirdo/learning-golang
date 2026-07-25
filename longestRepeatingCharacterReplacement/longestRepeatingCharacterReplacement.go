package main

import "fmt"

func characterReplacement(s string, k int) int {
	// the main heuristic here is that we don't need to update the maxFrequency.
	// it should be treated as "maxFrequency seen so far", NOT as "maxFrequency in the current window".

	// Basically, the bigger maxFrequency, the bigger the size of the window.
	// If we encountered some maxFrequency, we can just add any other K characters and get the maxWindowSize.
	// The invariant is:
	// windowSize - maxFrequency <= k
	// -> windowSize <= maxFrequency + k
	// i.e. maxFrequency defines the current maxWindowSize.

	// !!! But the K additional characters must be within the window.
	// I.e. the maxFrequency is only valid within the window of some size.

	// When maxFrequency is stale, invalid windows will validate as valid,
	// !!! but maxWindowSize will not increase until maxFrequency increases.

	// Note that frequencyMap is NOT stale.

	// see https://leetcode.com/problems/longest-repeating-character-replacement/editorial/
	// see https://www.youtube.com/watch?v=gqXU1UyA8pk

	// frequency map: character to count in the current windows
	freq := make(map[byte]int)

	maxWindowSize := 1

	maxFrequency := 0

	left := 0

	for right := 0; right < len(s); right++ {
		rightChar := s[right]

		freq[rightChar]++

		if freq[rightChar] > maxFrequency {
			maxFrequency = freq[rightChar]
		}

		windowsSize := right - left + 1
		charactersToReplace := windowsSize - maxFrequency

		if charactersToReplace <= k { // valid window -> check whether the window size is the new maximum
			maxWindowSize = max(maxWindowSize, windowsSize)
		} else { // shrink window from left by 1 character
			// !!! we do NOT decrease maxFrequency here

			leftChar := s[left]
			freq[leftChar]--
			left++
		}
	}

	return maxWindowSize
}

func test(s string, k int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("String: %v \n", s)
	fmt.Printf("K (maximum replacements): %v \n", k)

	result := characterReplacement(s, k)

	fmt.Printf("Longest substring with <= %v replacements to get one letter string: %v \n", k, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "ABAB"
	k := 2
	expected := 4

	test(s, k, expected)
}

func test2() {
	s := "AABABBA"
	k := 1
	expected := 4

	test(s, k, expected)
}

func main() {
	// 424. Longest Repeating Character Replacement
	test1()
	test2()
}
