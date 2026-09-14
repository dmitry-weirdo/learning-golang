package main

import (
	"fmt"
	"strconv"
)

func validWordAbbreviation(word string, abbr string) bool {
	i := 0 // index in word
	j := 0 // index in abbr

	m := len(word)
	n := len(abbr)

	for j < n {
		for (j < n) && !isDigit(abbr[j]) { // all non-digit chars must match between word and abbr
			if i >= m {
				return false
			}

			if word[i] != abbr[j] {
				return false
			}

			//fmt.Printf("Skipping char '%c' in both word and pattern. \n", word[i])

			// move one char in both strings
			i++
			j++
		}

		if j >= n {
			break
		}

		// parse digits
		if abbr[j] == '0' { // number cannot start with 0
			return false
		}

		numberStart := j
		for (j < n) && isDigit(abbr[j]) {
			j++
		}

		numberString := abbr[numberStart:j]
		//fmt.Printf("Number string: \"%v\" \n", numberString)

		number, err := strconv.Atoi(numberString)

		if err != nil { // this must NOT happen, invalid input
			panic(fmt.Sprintf("Cannot convert string \"%v\" to number.", numberString))
		}

		//fmt.Printf("Number parsed: %v \n", number)

		// skip number of chars in word
		i += number
		if i > m { // we can go just 1 char over the word end, else it's too many characters
			return false
		}

		//fmt.Printf("Skipped %v characters in word. \n", number)
	}

	// we must reach end of the word
	return i >= m
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func test(w, a string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Word: %v \n", w)
	fmt.Printf("Abbreviation: %v \n", a)

	result := validWordAbbreviation(w, a)

	fmt.Printf("Word \"%v\" can be abbreviated as \"%v\": %v \n", w, a, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("internationalization", "i12iz4n", true)
}

func test2() {
	test("apple", "a2e", false)
}

func test3() {
	test("substitution", "s10n", true)
	test("substitution", "sub4u4", true)
	test("substitution", "12", true) // whole word is 12 letters
	test("substitution", "su3i1u2on", true)
	test("substitution", "substitution", true) // not abbreviating is valid

	test("substitution", "s55n", false)          // 55 will read as 55, not as 5 + 5
	test("substitution", "s010n", false)         // leading 0 not allowed
	test("substitution", "s0ubstitution", false) // just 0 not allowed (replacing empty string not allowed)
}

func main() {
	// 408. Valid Word Abbreviation
	test1()
	test2()
	test3()
}
