package main

import "fmt"

func longestPrefix(s string) string {
	// todo: we should comparing using a rolling hash

	// KMP algorithm - we're just calculating LPS, should be O(n)
	// passes in 0-5 ms !!!
	return longestPrefix_kmp_lps(s)

	// Working solution by comparing substrings.
	// Same brute-force O(n^2) solution as in "214. Shortest Palindrome".
	// It still passes the tests in around 180+ ms
	//return longestPrefix_bruteForce(s)
}

func longestPrefix_kmp_lps(s string) string {
	// basically we're asked for LPS[last] by definition -> starting LPS-prefix for the end of the string

	lps := calculateLPS(s)

	lastLps := lps[len(lps)-1] // this is the length of the target substring

	return s[:lastLps]
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
		//printLPSState(s, lps, i, pl)

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

func longestPrefix_bruteForce(s string) string {
	if s == "" { // corner-case
		return ""
	}

	n := len(s)
	fmt.Printf("String length: %v \n", n)

	for i := n - 2; i >= 0; i-- { // we exclude the complete string
		// this should not require copying of the substrings, so probably should be faster as reversing every prefix?
		prefix := s[0 : i+1]
		suffix := s[n-i-1 : n]

		match := prefix == suffix

		if match {
			return prefix
		}
	}

	return ""
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := longestPrefix(s)

	fmt.Printf("Longest prefix that is also a suffix: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("level", "l")
}

func test2() {
	test("ababab", "abab") // whole string we don't count, "ababa" is not a suffix
}

func test3() {
	test("abcd", "") // no suffix == prefix
}

func test4() {
	test("a", "") // corner-case -> just 1 char -> we exclude the whole string -> not found
}

func main() {
	// 1392. Longest Happy Prefix
	test1()
	test2()
	test3()
	test4()
}
