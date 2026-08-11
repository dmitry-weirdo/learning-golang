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

	// Super-stupid finding of the starting palindrome, just by comparing the strings [0:i] and its reverse.
	// Should be O(n^2), but no matrix required.
	// Fails on TLE on Test-case 123 / 126, string length 50000

	// If we iterate from longest starting substring to shortest (to break fast),
	// it fails on TLE on 125 / 126, string length 42000
	// Although, with this variation, a couple of times I passed the tests in the measly 2043ms and 2123ms!
	return longestStartingPalindrome_bruteForce(s)

	// Using my solution for "5. Longest Palindromic Substring"
	// with O(n^2) and collectin a dp for s[i:j] is palindrome or not.
	// Basically, we need to find the longest palindrome from the start of the string s[0:j]

	// It fails on OOM for the matrix for a string with length 40002.
	// And the N limit is [0:50_000], while in "5. Longest Palindromic Substring", it's just [1:1000]
	// Test-case 121/ 126, string length 40002
	//return longestPalindrome_dp(s)
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
