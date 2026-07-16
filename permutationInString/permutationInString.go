package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	// todo: many solutions just use array[26] and do not exclude the non-s1 characters.
	// todo: in that case, non-s1 characters just have frequencies decreasing from 0 to negative, and never go above 0, i.e. they won't affect both decreasing and increasing cases hurting the charactersRequired variable
	// count frequencies of characters
	freq := make(map[rune]int)

	for _, v := range s1 {
		freq[v]++ //  wow, it works even without initializing to 0
	}

	fmt.Printf("Map of frequencies of string \"%v\": \n%v \n", s1, freq)

	// this is the main trick of not checking the complete freq[i] to have 0 values.
	charactersRequired := len(freq)
	fmt.Printf("Total different characters in string \"%v\": %v \n", s1, charactersRequired)

	windowSize := len(s1)

	// i points to the end of the window
	for i, char := range s2 {
		fmt.Println()
		fmt.Printf("i: %v, s2[i] = %c \n", i, char)

		// add the new character
		if _, ok := freq[char]; ok { // only handle characters from s1
			freq[char]--

			fmt.Printf("Decreased frequency of char %c to %v \n", char, freq[char])

			if freq[char] == 0 {
				charactersRequired--
				fmt.Printf("Frequency of char %c set to 0. Decreased charactersRequired to %v. \n", char, charactersRequired)
			}
		}

		if i >= windowSize { // remove the last character of the previous window
			removedCharPosition := i - windowSize
			removedChar := rune(s2[removedCharPosition])

			fmt.Printf("Window moved. Decreasing frequency of s2[%v] = %c. \n", removedCharPosition, removedChar)

			if _, ok := freq[removedChar]; ok { // only handle characters from s1
				freq[removedChar]++

				fmt.Printf("Increased frequency of removed char %c to %v \n", char, freq[removedChar])

				if freq[removedChar] == 1 {
					// we moved from 0 required to 1 required -> charactersRequired has to be increased
					// !!! Notably, when we move from -1 to 0, it's ok, no need to increase charactersRequired, character frequency still satisfied

					// s1 = "adc"
					// s2 = "dcda"
					// By i = 2, freq[d] = -1, freq[a] = 1
					// By i = 3, freq[a] will go from 1 to 0, freq[d] will go from -1 to 0, and this doesn't hurt the charactersRequired

					charactersRequired++
					fmt.Printf("Frequency of removed char %c set to 1. Increased charactersRequired to %v. \n", removedChar, charactersRequired)
				}
			}
		}

		fmt.Printf("Frequencies map: %v \n", freq)

		if charactersRequired == 0 { // all frequencies are exactly matched -> we found a required substring
			return true
		}
	}

	return false
}

func test(s1, s2 string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String to be included: %v \n", s1)
	fmt.Printf("String include: %v \n", s2)

	result := checkInclusion(s1, s2)

	fmt.Printf("\"%v\" is permutation-included in \"%v\": %v \n", s1, s2, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s1 := "ab"
	s2 := "eidbaooo"
	expected := true

	test(s1, s2, expected)
}

func test2() {
	s1 := "ab"
	s2 := "eidboaoo"
	expected := false

	test(s1, s2, expected)
}

func test3() {
	// tough test case from commends
	s1 := "adc"
	s2 := "dcda"
	expected := true

	test(s1, s2, expected)
}

func main() {
	// 567. Permutation in String
	test1()
	test2()
	test3()
}
