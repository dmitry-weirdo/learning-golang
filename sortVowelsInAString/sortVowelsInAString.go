package main

import "fmt"

func sortVowels(s string) string {
	b := []byte(s)

	//n := len(b)

	freq := make([]int, 10)

	// todo: for even faster time, we can also collect the vowel indices, to not iterate the complete string when replacing

	// count frequencies of vowels
	totalVowels := 0

	for _, v := range b {
		index := getVowelIndex(v)
		if index != -1 {
			freq[index]++
			totalVowels++
		}
	}

	//fmt.Printf("Frequencies of vowels: %v \n", freq)

	if totalVowels <= 0 { // no vowels -> nothing to replace -> return the source string
		return s
	}

	vowels := "AEIOUaeiou"

	freqIndex := getNextFrequencyIndex(freq, 0)

	for i, v := range b {
		if !isVowel(v) {
			continue
		}

		b[i] = vowels[freqIndex]
		freq[freqIndex]-- // decrease the frequency for the used vowel

		// go to next non-empty vowel can be the current)
		freqIndex = getNextFrequencyIndex(freq, freqIndex)

		if freqIndex < 0 { // all vowels used -> no need for further iteration
			break
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

func getVowelIndex(b byte) int {
	switch b {
	case 'A':
		return 0
	case 'E':
		return 1
	case 'I':
		return 2
	case 'O':
		return 3
	case 'U':
		return 4

	case 'a':
		return 5
	case 'e':
		return 6
	case 'i':
		return 7
	case 'o':
		return 8
	case 'u':
		return 9

	default:
		return -1
	}
}

func getNextFrequencyIndex(freq []int, i int) int {
	for j := i; j < len(freq); j++ {
		if freq[j] > 0 {
			return j
		}
	}

	return -1
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := sortVowels(s)

	fmt.Printf("String with sorted vowels: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("lEetcOde", "lEOtcede")
}

func test2() {
	test("lYmpH", "lYmpH") // Y is not a vowel!
}

func main() {
	// 2785. Sort Vowels in a String
	test1()
	test2()
}
