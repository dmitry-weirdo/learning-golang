package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	// todo: implement Z-algorithm
	// todo: implement Boyer-Moore
	// todo: implement KMP algorithm

	// KMP algorithm - O(m + n)
	// O(m) - to build a LPS array on the pattern
	// O(n) - iterate the search string
	return indexOf_kmp(haystack, needle)

	// Built-in Go algorithm - either brute-force or Rabin-Karp
	//return indexOf_trivial(haystack, needle)
}

func indexOf_trivial(s, p string) int {
	return strings.Index(s, p) // will do either brute-force or Rabin-Karp
}

func indexOf_kmp(s, p string) int {
	// see https://www.youtube.com/watch?v=V5-7GzOfADQ

	lps := calculateLPS(p) // we're calculating lps on pattern
	fmt.Printf("LPS for substring pattern \"%v\": \n%v \n", p, lps)

	i := 0 // index in string
	j := 0 // index in pattern (substring), also applies as index in LPS

	for i < len(s) {
		// !! j means that we've already matched characters [0 : j-1] to the substring before s[i]

		// characters match -> go to the next characters in both of the strings.
		if s[i] == p[j] {
			i++
			j++
		} else if j == 0 { // characters do not match
			// no pattern characters matched yet before s[i]
			// and as checked above, the current character does not match -> not a single character matches to pattern prefix
			// -> continues to the next s[i], we will still search from the p[0], 0-th character in the pattern
			i++
		} else {
			// similar to LPS calculation -> how many prefix characters we still have matched with the string before s[i]
			j = lps[j-1]
		}

		if j == len(p) { // we went over the last character in the pattern -> complete match found
			return i - len(p)
		}
	}

	return -1 // no match found -> -1
}

func calculateLPS(s string) []int {
	// we're calculating this normally on the substring (pattern) that we're searching for

	// see https://www.youtube.com/watch?v=V5-7GzOfADQ
	// see https://www.youtube.com/watch?v=JoF0Z7nVSrA

	// LPS is "Longest Prefix Suffix"
	// LPS[i] means
	// -> what is the longest substring ending on index [i] that is a prefix of s,
	// !!! not including the substring [0:i] itself

	/*
		NB: after I spent O(N^2) time while trying to understand the LPS calculation magic, I got that prevLPS should be better named as prefixLen or matchedPrefixLen or nextPrefixCharacterToMatch.
		I would call it PL for brevity.

		Its meaning is:
		- We currently matched (PL - 1) characters from the start of the string, they match these (PL - 1) characters before index i.
		- Currently we're checking to match the next character, i.e. we're checking whether p[i] = p[PL]. I.e. whether the prefix character PL is equal to the current character.

		If they don't match and PL = 0 -> it means we had 0 characters matched before and the current character is also not the 0-th character, so we set LPS[i] = 0 and move on with i++, starting to match from the beginning of the string.

		And then comes the most non-obvious part of all LPS calculation:
		if p[i] != p[PL] and PL != 0, we're checking LPS[PL - 1] and setting PL to it.

		LPS[PL - 1] -- for the previous character of the prefix, we're checking, what is the length of the prefix that we're currently matched before p[i], and we're jumping to this "previous longest matched prefix" for the next iteration.

		An example that makes it better to understand is string ABABAC, i = 5, pL = 3 , LPS = [0 0 1 2 3 x].
		So we've matched ABA, we're checking p[5] ?= p[3] -> C != B.
		Then we're looking at LPS[3 - 1] = LPS[2] = 1.
		We're setting PL =1, that means: prefix A was matched, we're checking the next character p[1] = B against p[i].
		And yes, we see that A is matched before C.

		If we just decreased PL = PL - 1, we would have gotten PL = 2 that means "we've matched the prefix AB, check next character for A".
		But this is a false statement, since we don't have AB before C.

		Hope this helps.
	*/

	lps := make([]int, len(s))

	// current index s[i] where we're filling the LPS
	i := 1 // we're starting from 0, lps[0] is always 0

	// s[pl] points to the index of the prefix that we're trying to current match with s[i].
	// It means that we have already matched pl characters [0:pl-1] (since strings are 0-based).
	pl := 0

	lps[0] = 0 // for the first character, there were no previous strings to match.

	for i < len(s) {
		printLPSState(s, lps, i, pl)

		if s[i] == s[pl] { // we matched the next prefix character -> current sum is (pl + 1), go to the next character in both prefix and i,
			lps[i] = pl + 1 // s[pl] matched -> we've already matched (pl + 1) characters since strings are 0-based.

			pl++ // try to match the next prefix character
			i++  // go to next string character
			continue
		}

		if pl == 0 {
			// pl == 0 means that we haven't matched any characters in the prefix yet.
			// Because of the above check, we already know that the current character s[i] != s[pl],
			// i.e. we don't match even with the 0-th character prefix
			// -> set lps[i] == 0 and move to the next character.

			// pl remains at 0, i.e. for the next character, we will also try to match it to the 0-th prefix character
			lps[i] = 0
			i++
		} else {
			// pl != 0 means that we have matched some pattern characters (all s[0:pl], but the match broke on s[pl] != s[i].

			// LPS[PL - 1] -- for the previous character of the prefix, we're checking, what is the length of the prefix
			// that we're currently matched before p[i], and we're jumping to this "previous longest matched prefix" for the next iteration.

			// So, after this jump, we know that we've matched prefix characters s[0:pl]
			// with the string characters before s[i].

			// To understand this magic trick, see the "ABABAC" example for the last character C above.
			pl = lps[pl-1]
		}
	}

	return lps
}

func printLPSState(s string, lps []int, i int, pl int) {
	fmt.Println()
	fmt.Printf("i = %v, pl = %v, matched %v prefix characters \"%v\". \n", i, pl, pl, s[0:pl])
	fmt.Printf("Trying to match prefix character s[%v] = '%c' to string character s[%v] = '%c'... \n", pl, s[pl], i, s[i])
	fmt.Println()

	// 0 1 2 3 4 5   // indexes
	// _ _ ?         // matched prefix characters, current checking prefix character
	// A B A B A C   // string
	//         i      // i pointer -> current character we're trying to match with a prefix
	// 0 0 1 2 ?      // LPS

	// todo: make better if numbers are > 10
	for j := 0; j < len(s); j++ {
		fmt.Printf("%v ", j)
	}
	fmt.Println(" | index ")

	for j := 0; j < pl; j++ { // we already matched characters [0: pl-1]
		fmt.Printf("_ ")
	}

	fmt.Printf("? ") // testing this prefix character, s[pl] is ?

	for j := pl + 1; j < len(s); j++ {
		fmt.Printf("  ")
	}
	fmt.Printf(" | prefix matched and checking \n")

	// print the string characters with space separator
	for j := 0; j < len(s); j++ {
		fmt.Printf("%c ", s[j])
	}
	fmt.Println(" | string ")

	// print i pointer
	for j := 0; j < i; j++ {
		fmt.Printf("  ")
	}
	fmt.Printf("i ")

	for j := i + 1; j < len(s); j++ {
		fmt.Printf("  ")
	}
	fmt.Printf(" | current index to find the prefix match \n")

	for j := 0; j < i; j++ {
		fmt.Printf("%v ", lps[j])
	}
	fmt.Printf("? ")

	for j := i + 1; j < len(s); j++ {
		fmt.Printf("  ")
	}
	fmt.Printf(" | current LPS \n")
}

func testCalculateLPS(s string, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("String: %v \n", s)

	result := calculateLPS(s) // todo: replace with your function

	fmt.Printf("LPS of \"%v\": %v \n", s, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func testLPS1() {
	s := "ABABAC"
	expected := []int{0, 0, 1, 2, 3, 0}

	testCalculateLPS(s, expected)
}

func testLPS2() {
	s := "AAACAAAA"
	expected := []int{0, 1, 2, 0, 1, 2, 3, 3}

	testCalculateLPS(s, expected)
}

func testLPS3() {
	s := "ababd"
	expected := []int{0, 0, 1, 2, 0}

	testCalculateLPS(s, expected)
}

func testLPS4() {
	s := "aaaa"
	expected := []int{0, 1, 2, 3}

	testCalculateLPS(s, expected)
}

func testLPS5() {
	s := "aabcadaabc"
	expected := []int{0, 1, 0, 0, 1, 0, 1, 2, 3, 4}

	testCalculateLPS(s, expected)
}

func testLPS6() {
	s := "abcdabeabf"
	expected := []int{0, 0, 0, 0, 1, 2, 0, 1, 2, 0}

	testCalculateLPS(s, expected)
}

func testLPSSuite() {
	testLPS1()
	testLPS2()
	testLPS3()
	testLPS4()
	testLPS5()
	testLPS6()
}

func test(s, p string, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)                      // todo: replace with your text if required
	fmt.Printf("Pattern (substring) to find: %v \n", p) // todo: replace with your text if required

	result := strStr(s, p) // todo: replace with your function

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("sadbutsad", "sad", 0)
}

func test2() {
	test("leetcode", "leeto", -1)
}

func main() {
	// 28. Find the Index of the First Occurrence in a String
	test1()
	test2()

	// KMP algorithm test
	//testLPSSuite()
}
