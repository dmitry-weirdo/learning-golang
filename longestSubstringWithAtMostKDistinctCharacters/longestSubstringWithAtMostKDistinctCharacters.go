package main

import "fmt"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	return slidingWindowNonShrinking(s, k)
	//return slidingWindowShrinking(s, k)
}

func slidingWindowShrinking(s string, k int) int {
	// O(n)
	// my obvious solution, left is shrinking to become valid even if the window becomes smaller then maxWindowSize
	n := len(s)
	if k >= n {
		// there are more allowed characters than the length of the string -> the complete string will fit
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

func slidingWindowNonShrinking(s string, k int) int {
	// O(n)
	// slightly optimized -> when a window is invalid, we only shrink it from left by 1 -> to the current maxWindowSize
	n := len(s)
	if k >= n {
		// there are more allowed characters than the length of the string -> the complete string will fit
		return n
	}

	// window validity is freqMap.size <= k
	freq := make(map[byte]int)

	// string of size K cannot contain more than K distinct characters
	maxWindowSize := k

	left := 0

	for right := 0; right < n; right++ {
		fmt.Println()
		fmt.Printf("left: %v, right: %v, windowSize: %v, maxWindowSize: %v \n", left, right, right-left+1, maxWindowSize)
		fmt.Printf("Current window: \"%v\" \n", s[left:right+1])

		// add right character to the map
		rightChar := s[right]
		freq[rightChar]++

		if len(freq) <= k {
			fmt.Printf("Window is valid. \n")

			windowSize := right - left + 1
			maxWindowSize = max(maxWindowSize, windowSize)
		} else {
			fmt.Printf("Window is invalid. Shrinking 1 from the left. \n")

			// Shrink the window from left just by 1 character
			// (it will decrease from maxWindowSize + 1 (failed attempt) to maxWindowSize),
			// so that we keep the window on the current size (which would be maxWindowSize).
			// The window can remain invalid, but it will be of size maxWindowSize.
			// So we'll try to continue right-expansion into (maxWindowSize + 1) until it's valid, or we reach the end of the string.

			//for len(freq) > k {
			leftChar := s[left]

			freq[leftChar]--

			if freq[leftChar] <= 0 {
				delete(freq, leftChar)
			}

			left++
			//}
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
