package main

import (
	"fmt"
)

func maxPalindromes(s string, k int) int {
	if k == 1 { // use every character as palindrome -> no need for complex calculations
		return len(s)
	}

	// First, we calculate the isPalindrome[i; j], same as in "5. Longest Palindromic Substring"
	m := getPalindromeMatrix(s)

	//fmt.Printf("Is palindrome matrix: \n")
	//printMatrix(s, m)

	// Then, we do DP on selecting the palindromic substring on next index.
	// At index [i], we can:
	// - Skip the current index -> just continue from [i + 1]
	// - Use the current index -> select the shortest palindrome starting from [i] and having length at least k

	n := len(s)

	// max palindromes starting with index [i]
	memo := createIntArrayWithDefaultValues(n, -1) // 0 might be a valid value

	var dfs func(i int) int

	dfs = func(i int) int {
		if i >= n { // went over the last char
			return 0
		}

		if memo[i] >= 0 { // value already pre-calculated -> return it
			return memo[i]
		}

		// case 1 - skip using the current value as palindrome
		skipCurrent := dfs(i + 1)

		// case 2 - use the first position -> find the first palindrome starting with [i] and length >= k
		useCurrent := -1
		for j := i + k - 1; j < n; j++ {
			if m[i][j] { // palindrome from i found -> count this palindrome and continue DFS from [j + 1]
				useCurrent = 1 + dfs(j+1)
				break
			}
		}

		memo[i] = max(skipCurrent, useCurrent)
		return memo[i]
	}

	return dfs(0)
}

func getPalindromeMatrix(s string) [][]bool {
	n := len(s)

	// DP-memo matrix of whether string [i, j] is a palindrome
	// values of i > j (bottom-left) are not used.
	m := make([][]bool, n)

	// put true values to [i][i] (since 1 character is a palindrome), other values to false
	for i := range n {
		m[i] = make([]bool, n)
		m[i][i] = true
	}

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

	return m
}

func printMatrix(s string, mat [][]bool) {
	rows := len(mat)
	columns := len(mat[0])

	// row with string char on each column
	fmt.Printf("  ")

	for i := range rows {
		fmt.Printf("%c ", s[i])
	}

	fmt.Println()

	for i := range rows {
		fmt.Printf("%c ", s[i])

		for j := range columns {
			var v string

			// this is very specific for this task, so we use this method instead of matrixCommon.PrintBoolMatrix
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

func createIntArrayWithDefaultValues(n int, defaultValue int) []int {
	a := make([]int, n)

	for i := range n {
		a[i] = defaultValue
	}

	return a
}

func test(s string, k int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)
	fmt.Printf("Min length of palindrome: %v \n", k)

	result := maxPalindromes(s, k)

	fmt.Printf("Max substrings of length >= %v: %v \n", k, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("abaccdbbd", 3, 2) // aba, dbbd
}

func test2() {
	test("adbcda", 2, 0) // no palindrome substrings of len >= 2
}

func main() {
	// 2472. Maximum Number of Non-overlapping Palindrome Substrings
	test1()
	test2()
}
