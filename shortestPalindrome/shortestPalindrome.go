package main

import (
	"fmt"
	"slices"
)

func shortestPalindrome(s string) string {
	// todo: we need to implement other solutions:
	// KMP - how?
	// Rolling hash - hash algorithm is like in Rabin-Karp's algorithm
	// Manacher's algorithm - O(n) to find the longest palindromic substring -> // todo: also apply to "5. Longest Palindromic Substring" to reduce the time!

	// LPS of KMP of "<originalString>#<reversedString>:
	// O(2*n + 1) = O(2*n) = O(n) - to build the LPS array.
	// + actually O(n) on reverse, O(n) on concatenate, O(n) on constructing the result.
	// But it all is O(n)
	// It passes in 0-6 ms!
	return longestStartingPalindrome_bruteForce_lpsOfReversed(s)

	// ===================================================================================== //
	// Try to match the prefix(originalString) and suffix(reversedString)
	// Still O(n^2), but I hope that substrings will be just pointers and don't require memory copying as for string reversing.
	// Why O(n^2) - for every length from (n-1) down to 2, we're comparing the strings of this length.
	// !!! It passes in 6-7 ms!
	//return longestStartingPalindrome_bruteForce_comparePrefixesWithReverseSuffixes(s)

	// ===================================================================================== //
	// Super-stupid finding of the starting palindrome, just by comparing the strings [0:i] and its reverse.
	// Should be O(n^2), but no matrix required.
	// Fails on TLE on Test-case 123 / 126, string length 50000

	// If we iterate from longest starting substring to shortest (to break fast),
	// it fails on TLE on 125 / 126, string length 42000
	// Although, with this variation, a couple of times I passed the tests in the measly 2043ms and 2123ms!
	//return longestStartingPalindrome_bruteForce(s)

	// ===================================================================================== //
	// Using my solution for "5. Longest Palindromic Substring"
	// with O(n^2) and collectin a dp for s[i:j] is palindrome or not.
	// Basically, we need to find the longest palindrome from the start of the string s[0:j]

	// It fails on OOM for the matrix for a string with length 40002.
	// And the N limit is [0:50_000], while in "5. Longest Palindromic Substring", it's just [1:1000]
	// Test-case 121/ 126, string length 40002
	//return longestPalindrome_dp(s)
}

func longestStartingPalindrome_bruteForce_lpsOfReversed(s string) string {
	// generate string: "<originalString>#<reversedString>",
	// where # is a character not in the string,
	// so that the LPS-prefix for the reversed string can NOT grow beyond "<originalString>".

	// At the end of the reversed string, we will have the starting character of the original string.
	// By definition, LPS is "longest prefix string that is equal to a string ending on the current position".
	// Prefix string -> is the start of the original string.
	// String in the end "<originalString>#<reversedString>" is the reversed prefix of the original string.
	// So LPS[last] will give us the longest prefix that matches with its own reverse. This is the longest starting palindrome.
	// I.e. this is the longest prefix string that matches with its reverse.

	// Example:
	// original string "abade" ->  "abade#edaba"
	// LPS[last] will be length of the "aba" string.

	// If the original string is a palindrome itself:
	// "aba#aba", then the unique separator character "#" guarantees us that the longest prefix match
	// will NOT go after this "#", since the <reversed> part does NOT contain this "#".
	// I.e, after "aba" match, extending the match will stop since "aba#" is not contained in the part after "#".

	reversed := reverseString(s)

	stringForLps := s + "#" + reversed

	lps := calculateLPS(stringForLps)

	longestStaringPalindromeLength := lps[len(lps)-1]

	return getPalindromeByLongestStartingPalindrome(s, longestStaringPalindromeLength-1)
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

func longestStartingPalindrome_bruteForce_comparePrefixesWithReverseSuffixes(s string) string {
	if s == "" { // corner-case
		return ""
	}

	n := len(s)
	fmt.Printf("String length: %v \n", n)

	// todo: the substrings can be calculated with 2 opposite growing hashes (see Solution.java), then it will be O(n) and will pass!

	// super stupid finding of the longest starting palindrome
	longestPalindromeEndIndex := 0 // start with just 1 character

	reversed := reverseString(s)

	for i := n - 1; i >= 0; i-- {
		// this should not require copying of the substrings, so probably should be faster as reversing every prefix?
		prefixOfOriginal := s[0 : i+1]
		suffixOfReversed := reversed[n-i-1 : n]

		substringIsPalindrome := prefixOfOriginal == suffixOfReversed

		if substringIsPalindrome {
			longestPalindromeEndIndex = i
			break
		}
	}

	return getPalindromeByLongestStartingPalindrome(s, longestPalindromeEndIndex)
}

func longestStartingPalindrome_bruteForce(s string) string {
	if s == "" { // corner-case
		return ""
	}

	fmt.Printf("String length: %v \n", len(s))

	// todo: we can also construct by reversing the string and comparing prefixes(s) and suffixes(reversedS). Will it be better time (probably still O(n^2)???

	// super stupid finding of the longest starting palindrome
	longestPalindromeEndIndex := 0

	/*
		// Try substrings from s[0:1] up to s[0:n-1].
		// The latest found palindrome will be the longest.
		for i := 1; i < len(s); i++ {
			substring := s[:i+1]

			substringIsPalindrome := reverseString(substring) == substring

			if substringIsPalindrome {
				longestPalindromeEndIndex = i
			}
		}
	*/

	// Yes, this runs faster, I passed the test-cases in 2000+ ms.
	// Try substrings from s[0:n-1] down to s[0:1].
	// The first found palindrome will be the longest.
	for i := len(s) - 1; i >= 0; i-- {
		substring := s[:i+1]

		substringIsPalindrome := reverseString(substring) == substring

		if substringIsPalindrome {
			longestPalindromeEndIndex = i
			break
		}
	}

	return getPalindromeByLongestStartingPalindrome(s, longestPalindromeEndIndex)
}

func longestPalindrome_dp(s string) string {
	// this is a DP-algorithm from "5. Longest Palindromic Substring"

	if s == "" { // corner-case
		return ""
	}

	n := len(s)

	fmt.Printf("String length: %v \n", n)

	// DP-memo matrix of whether string [i, j] is a palindrome
	// values of i > j (bottom-left) are not used.
	m := make([][]bool, n)

	// put true values to [i][i] (since 1 character is a palindrome), other values to false
	for i := range n {
		m[i] = make([]bool, n)
		m[i][i] = true
	}

	//fmt.Println("Initialized the matrix:")
	//printBooleanMatrix(m)

	// we skip i = n - 1, it's just the last character that is already true
	for i := n - 2; i >= 0; i-- {
		// go from smaller intervals to bigger since we need the values of [i + 1; j - 1]
		for j := i + 1; j <= n-1; j++ {
			if s[i] != s[j] {
				// first and last characters of substring are non-equal -> not a palindrome
				m[i][j] = false // todo: this is not super-necessary since initialized to false anyway
				continue
			}

			isPalindrome := false

			if ((i + 1) >= (j - 1)) || // if i + 1 = j, it is a palindrome of 2 characters!
				m[i+1][j-1] { // substring within i and j is a palindrome
				isPalindrome = true
			}

			if !isPalindrome {
				m[i][j] = false // todo: this is not super-necessary since initialized to false anyway
				continue
			}

			// [i;j] is a palindrome -> check whether it's longer than the current max palindrome
			m[i][j] = true
		}
	}

	//fmt.Println("Matrix after all checks:")
	//printBooleanMatrix(m)

	// find longest s[0:j] palindrome substring
	longestStartingPalindromeEndIndex := 0

	for j := n - 1; j >= 0; j-- {
		if m[0][j] {
			longestStartingPalindromeEndIndex = j
			break
		}
	}

	return getPalindromeByLongestStartingPalindrome(s, longestStartingPalindromeEndIndex)
}

// If we know the longest starting (prefix) palindrome,
// we easily calculate the result by appending the reversed suffix and the original string.
// Suffix is the string after the longest starting palindrome.
func getPalindromeByLongestStartingPalindrome(s string, longestStartingPalindromeEndIndex int) string {
	if longestStartingPalindromeEndIndex == (len(s) - 1) { // string is a palindrome -> no further change
		fmt.Printf("Whole string \"%v\" is a palindrome. Returning it. \n", s)
		return s
	}

	longestStartingPalindrome := s[0 : longestStartingPalindromeEndIndex+1]
	fmt.Printf("Longest starting palindrome: \"%v\" \n", longestStartingPalindrome)

	// end of the string after the palindromic start substring
	suffix := s[longestStartingPalindromeEndIndex+1:]
	reversedSuffix := reverseString(suffix)

	fmt.Printf("End suffix after palindromic substring: \"%v\" \n", suffix)
	fmt.Printf("Reversed suffix: \"%v\" \n", reversedSuffix)

	// we append to the original string, NOT to the starting palindrome
	return reversedSuffix + s
}

func reverseString(s string) string {
	stringAsSlice := []byte(s)
	slices.Reverse(stringAsSlice) // reverses in place

	return string(stringAsSlice)
}

func printBooleanMatrix(mat [][]bool) {
	rows := len(mat)
	columns := len(mat[0])

	for i := range rows {
		for j := range columns {
			var v string

			if i > j { // we don't care about these values
				v = "."
			} else if mat[i][j] {
				v = "T"
			} else {
				v = "F"
			}

			fmt.Printf("%v ", v)
		}

		fmt.Println()
	}
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := shortestPalindrome(s)

	fmt.Printf("Shortest palindrome by appending characters to the start of the string: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("aacecaaa", "aaacecaaa")
}

func test2() {
	test("abcd", "dcbabcd")
}

func test3() {
	test("", "")
}

func test4() {
	test("a", "a") // already a palindrome
}

func main() {
	// 214. Shortest Palindrome
	test1()
	test2()
	test3()
	test4()
}
