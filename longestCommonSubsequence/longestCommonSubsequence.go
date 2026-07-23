package main

import (
	"fmt"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	// row i - represents substring s[i; len(s) - 1]
	// column j - represents substring t[j; len(j) - 1]
	// DP table shows the longest subsequence for these substrings

	// if the characters at s[i] and t[j] match -> we add 1 to s[i + 1][j + 1]
	// (i.e. we go diagonally to bottom-right)
	// It basically means that we increase the existing subsequence from next characters by 1

	// if the characters at s[i] and t[j] do NOT match ->
	// we select the max of:
	// - s[i + 1][j] - one character less in s, same string in t
	// - s[i][j + 1] - same string in s, one character less in t
	// i.e. we move

	// example on "abcde" and "ace" strings
	// i = 4, j = 2 -> "e" matches "e" -> we add 1 to the 0 (out of bounds) -> set 1

	// i = 4, j = 1 -> "e" vs "ce" -> "e" does not match "c" -> this does not add + 1
	// we select max of
	// - i = 5, j = 1 -> "" vs "ce" -> out of range -> 0
	// - i = 4, j = 2 -> "e" vs "e" -> already calculated above to be 1 -> 1
	// max (0, 1) = 1
	// i.e. "e" vs "ce" produces a max common subsequence of len 1 which is "e"

	// we're going from bottom-right to left then up
	// (each row from right to left,
	// each column from bottom to top)

	m := len(text1)
	n := len(text2)

	fmt.Printf("m: %v, n: %v \n", m, n)

	// we're adding additional zeroes, so the table is (m + 1) * (n + 1)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// todo: we can actually use only two rows [i] and [i + 1], this can be optimized to previousRow and currentRow, to reduce the space complexity from O(m * n) to O(2 * n)
	// calculate the DP table
	for i := m - 1; i >= 0; i-- { // rows go bottom to top
		for j := n - 1; j >= 0; j-- { // columns go right to left
			if text1[i] == text2[j] {
				// go diagonally -> add the current character (1 to length) to the substrings of s[i + 1:] and t[j + 1:]
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				// current character [i][j] does not add to the matching subsequence
				// We skip the current character, and the next subsequence is the best (longest) of the following substrings:
				// - s[i + 1:] and t[j:]
				// - s[i:] and t[j + 1:]
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	fmt.Printf("DP matrix: \n")
	printMatrix(dp, text1, text2)

	// return dp[0][0], i.e. the subsequence length of s[0:] and t[0:]
	return dp[0][0]
}

func printMatrix(mat [][]int, s, t string) {
	rows := len(mat)
	columns := len(mat[0])

	// print the header, string T
	for j := 0; j < columns; j++ {

		if j == 0 {
			fmt.Printf("  ")
		} else {
			fmt.Printf("%c ", t[j-1])
		}
	}

	fmt.Println()

	for i := 0; i < rows; i++ {
		if i < len(s) {
			fmt.Printf("%c ", s[i])
		} else {
			fmt.Printf("  ")
		}

		for j := 0; j < columns; j++ {
			fmt.Printf("%v ", mat[i][j])
		}

		fmt.Println()
	}
}

func test(s, t string, expectedResult int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("S: %v \n", s)
	fmt.Printf("T: %v \n", t)

	result := longestCommonSubsequence(s, t)

	fmt.Printf("Longest common subsequence of \"%v\" and \"%v\" length: %v \n", s, t, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "abcde"
	t := "ace"
	expected := 3

	test(s, t, expected)
}

func test2() {
	s := "abc"
	t := "abc"
	expected := 3

	test(s, t, expected)
}

func test3() {
	s := "abc"
	t := "def"
	expected := 0

	test(s, t, expected)
}

func main() {
	// 1143. Longest Common Subsequence
	test1()
	test2()
	test3()
}
