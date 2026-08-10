package main

import "fmt"

func numberOfSubstrings(s string) int {
	// since we're only using characters abc, we're using an array instead of map, since it's faster
	freq := make([]int, 3)

	count := 0

	requiredChars := 3
	totalFoundChars := 0 // control how many characters are found in the current [left:right] window

	left := 0

	// for every right position, we find the last left containing all chars
	// and then all substrings on this left and before should be counted (for this right position)
	for right := 0; right < len(s); right++ {
		chRight := s[right]
		chRightIndex := getCharIndex(chRight)
		freq[chRightIndex]++

		// Freq changed from 0 to 1 -> increase totalFoundChars
		if freq[chRightIndex] == 1 {
			totalFoundChars++
		}

		// note whether the window contained enough chars
		enoughCharsBeforeShrinking := totalFoundChars >= requiredChars

		// shrink window from left while it has all chars
		for totalFoundChars >= requiredChars {
			chLeft := s[left]
			chLeftIndex := getCharIndex(chLeft)
			freq[chLeftIndex]--

			// Freq changed from 1 to 0 -> decrease totalFoundChars
			if freq[chLeftIndex] == 0 {
				totalFoundChars--
			}

			left++
		}

		if enoughCharsBeforeShrinking { // enough chars -> jump back one char back, so that [left:right] will contain all characters
			left--

			chLeft := s[left]
			chLeftIndex := getCharIndex(chLeft)
			freq[chLeftIndex]++

			totalFoundChars++
		}

		if enoughCharsBeforeShrinking {
			// for the current right, we add substrings from all indexes from 0 to left inclusive,
			// i.e. all substrings from [0:right] to [left:right] will count.
			count += left + 1 // left - 0 + 1
		}
	}

	return count
}

func getCharIndex(ch byte) int { // to map to array index
	return int(ch - 'a')
}

func test(s string, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s) // todo: replace with your text if required

	result := numberOfSubstrings(s)

	fmt.Printf("Count of substrings containing all characters `a`, `b`, `c`: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	// pos 0: abc, abca, abcab, abcabc - 4
	// pos 1: bca, bcab, bcabc - 3
	// pos 2: cab, cabc - 2
	// pos 3: abc - 1
	test("abcabc", 10)
}

func test2() {
	// aaacb, aacb, acb
	test("aaacb", 3)
}

func test3() {
	test("abc", 1)
}

func test4() {
	test("abb", 0)
}

func main() {
	// 1358. Number of Substrings Containing All Three Characters
	// Very similar to "76. Minimum Window Substring"
	test1()
	test2()
	test3()
	test4()
}
