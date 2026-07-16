package main

import "fmt"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// calculate frequencies required for t
	// count frequencies of characters
	freq := make(map[rune]int)

	for _, v := range t {
		freq[v]++ //  wow, it works even without initializing to 0
	}

	fmt.Printf("Target frequencies: %v \n", freq)

	totalCharactersRequired := len(t)
	fmt.Printf("Total characters required for \"%v\": %v \n", t, totalCharactersRequired)

	windowFreq := make(map[rune]int)

	left := 0
	right := 0

	// how many characters of are found in the current [left; right] window
	charactersFound := 0

	minWindowStart := -1
	minWindowLength := len(s) + 1 // any found valid window will be shorter than this

	for right < len(s) {
		rightChar := rune(s[right])
		windowFreq[rightChar]++ // will also increase for non-valid chars

		fmt.Printf("right: %v, s[right] = s[%v] = %c \n", right, right, rightChar)
		fmt.Printf("Window after expanding to right: %v \n", s[left:right+1])

		// window is not the fixed size (unlike the 567. Permutation in String) -> it may contain more characters from t than required:
		// "BAAAAAXXXXXXX" for "BAC" -> A satisfied many times, but C still not reached while expanding
		// we only increase charaсtersFound if window character frequency does NOT exceed the required character frequency
		if windowFreq[rightChar] <= freq[rightChar] {
			// for non-present characters in t, freq[rightChar] will be 0
			// so by increasing windowFreq[notRequiredChar], we'll set it to 1 and this is NOT a charactersFound increase
			charactersFound++

			fmt.Printf("windowFreq[%c] = %v still contributed to freq[%c] = %v. Increased charactersFound to %v. \n", rightChar, windowFreq[rightChar], rightChar, freq[rightChar], charactersFound)
		}

		for charactersFound == totalCharactersRequired { // we reached the valid window -> try to shrink from left until window becomes invalid
			currentWindowLength := right - left + 1 // right is inclusive

			if currentWindowLength < minWindowLength {
				minWindowStart = left
				minWindowLength = currentWindowLength

				fmt.Printf("Current valid window length = %v < current min window length = %v. Saving the new min window. \n", currentWindowLength, minWindowLength)
				fmt.Printf("New min window: [%v; %v] = \"%v\" \n", left, right, s[left:right+1])
			}

			leftChar := rune(s[left])
			windowFreq[leftChar]--

			fmt.Printf("left: %v, s[left] = s[%v] = %c \n", left, left, leftChar)
			fmt.Printf("Window after shrinking from left: %v \n", s[left:right+1])

			// we only decrease charactersFound if window character frequency is less than the required character frequency
			// i.e. removing the surplus characters from t does not hurt the charactersFound
			if windowFreq[leftChar] < freq[leftChar] {
				charactersFound--

				fmt.Printf("windowFreq[%c] = %v still contributed to freq[%c] = %v. Decreased charactersFound to %v. \n", leftChar, windowFreq[leftChar], leftChar, freq[rightChar], charactersFound)
			}

			left++
		}

		right++
	}

	if minWindowStart < 0 { // not a single target window found
		return ""
	}

	// window found -> return it
	return s[minWindowStart : minWindowStart+minWindowLength]
}

func test(s, t string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String to contain: %v \n", s)
	fmt.Printf("String to be included: %v \n", t)

	result := minWindow(s, t)

	fmt.Printf("Minimum window: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "ADOBECODEBANC"
	t := "ABC"
	expected := "BANC"

	test(s, t, expected)
}

func test2() {
	s := "a"
	t := "a"
	expected := "a"

	test(s, t, expected)
}

func test3() {
	s := "a"
	t := "aa"
	expected := ""

	test(s, t, expected)
}

func main() {
	// 76. Minimum Window Substring
	test1()
	test2()
	test3()
}
