package main

import "fmt"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	return slidingWindowShrinking(s, k)
}

func slidingWindowShrinking(s string, k int) int {
	// O(n)
	// my obvious solution, left is shrinking to become valid even if the window becomes smaller then maxWindowSize
	n := len(s)
	if k >= n {
		// there are more allowed characters then the length of the string -> the complete string will fit
		return n
	}

	// window validity is freqMap.size <= k
	freq := make(map[byte]int)

	// string of size K cannot contain more than K distinct characters
	maxWindowSize := k

	left := 0

	for right := 0; right < n; right++ {
		// add right character to the map
		rightChar := s[right]
		freq[rightChar]++

		if len(freq) <= k {
			windowSize := right - left + 1
			maxWindowSize = max(maxWindowSize, windowSize)
		} else { // shrink window from left to valid state

			for len(freq) > k {
				leftChar := s[left]

				freq[leftChar]--

				if freq[leftChar] <= 0 {
					delete(freq, leftChar)
				}

				left++
			}
		}
	}

	return maxWindowSize
}

func test(s string, k int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("String: %v \n", s)
	fmt.Printf("K (most distinct characters in window): %v \n", k)

	result := lengthOfLongestSubstringKDistinct(s, k)

	fmt.Printf("Maximum window with <= K = %v distinct characters: %v \n", k, result)
	fmt.Printf("Expected result:                                   %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "eceba"
	k := 2
	expected := 3 // ece

	test(s, k, expected)
}

func test2() {
	s := "aa"
	k := 1
	expected := 2 // aa

	test(s, k, expected)
}

func main() {
	// 340. Longest Substring with At Most K Distinct Characters
	test1()
	test2()
}
