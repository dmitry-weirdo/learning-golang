package main

import (
	"fmt"
)

func numDistinct(s string, t string) int {
	// O(m * n) time
	// O(n) space - just save one row of the matrix since we only need a previous row for every calculation
	// is faster because we don't need the M * N matrix initialization
	// passes in 0-1 ms
	return numDistinct_dp_oneArray(s, t)

	// O(m * n) time
	// O(m * n) space
	// passes in same 6-10 ms
	//return numDistinct_dp_iterative(s, t)

	// O(m * n) time
	// O(m * n) space
	// passes in 7-9 ms
	//return numDistinct_recursive(s, t)
}

func numDistinct_recursive(s string, t string) int {
	m := len(s)
	n := len(t)

	if m < n {
		return 0
	}

	if m == n {
		if s == t {
			return 1
		}

		return 0
	}

	// todo: this requires O(m * n) operations. Probably using a hashmap with Pair<i, j> key will be faster
	memo := createIntMatrixWithDefaultValues(m, n, -1)

	var dfs func(i, j int) int // returns a count of subsequences of s[i...M] and t[j...N]

	dfs = func(i, j int) int {
		if j == n { // base-case -> reached the end of T -> success
			return 1
		}

		if (m - i) < (n - j) { // not enough characters in S left -> stop iteration
			return 0
		}

		if memo[i][j] >= 0 { // already pre-calculated
			return memo[i][j]
		}

		if s[i] == t[j] {
			// characters s[i] and t[j] match -> 2 choices:
			// - use both characters (i + 1, j + 1)
			// - skip character in S (i + 1, j)
			memo[i][j] = dfs(i+1, j+1) + dfs(i+1, j)
		} else { // no match -> skip a character in S
			memo[i][j] = dfs(i+1, j)
		}

		return memo[i][j]
	}

	return dfs(0, 0)
}

func numDistinct_dp_iterative(s string, t string) int {
	m := len(s)
	n := len(t)

	if m < n {
		return 0
	}

	if m == n {
		if s == t {
			return 1
		}

		return 0
	}

	dp := createIntMatrix(m+1, n+1)

	// if we ran over S end, it's no solution
	for j := range n + 1 {
		dp[m][j] = 0
	}

	// if we ran over T end, it's a solution
	for i := range m + 1 {
		dp[i][n] = 1
	}

	//fmt.Printf("Initialized DP matrix: \n")
	//matrixCommon.PrintIntMatrix(dp)

	// iterate in reverse
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				// characters s[i] and t[j] match -> 2 choices:
				// - use both characters (i + 1, j + 1)
				// - skip character in S (i + 1, j)
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else { // no match -> skip a character in S
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	//fmt.Printf("Calculated DP matrix: \n")
	//matrixCommon.PrintIntMatrix(dp)

	return dp[0][0]
}

func numDistinct_dp_oneArray(s string, t string) int {
	m := len(s)
	n := len(t)

	if m < n {
		return 0
	}

	if m == n {
		if s == t {
			return 1
		}

		return 0
	}

	dp := make([]int, n)

	// iterate in reverse
	for i := m - 1; i >= 0; i-- {
		// if we ran over T end, it's a solution
		prev := 1 // dp[i+1][n], generally prev is dp[i+1][j+1]

		for j := n - 1; j >= 0; j-- {
			oldDpJ := dp[j] // dp[i+1][j]

			if s[i] == t[j] {
				// characters s[i] and t[j] match -> 2 choices:
				// - use both characters (i + 1, j + 1)
				// - skip character in S (i + 1, j)
				//dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
				dp[j] = prev + oldDpJ
			} else { // no match -> skip a character in S
				//dp[i][j] = dp[i+1][j]
				dp[j] = oldDpJ
			}

			// prev is dp[i+1][j+1]
			prev = oldDpJ
		}
	}

	return dp[0]
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func createIntMatrixWithDefaultValues(rows, columns int, defaultValue int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)

		for j := range columns { // !!! note that this is slow, will take O(m * n) additional operations :(
			m[i][j] = defaultValue
		}
	}

	return m
}

func test(s, t string, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String S: %v \n", s)
	fmt.Printf("String T: %v \n", t)

	result := numDistinct(s, t)

	fmt.Printf("Count of distinct subsequences of \"%v\" that are equal to \"%v\": %v \n", s, t, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("rabbbit", "rabbit", 3)
}

func test2() {
	test("babgbag", "bag", 5)
}

func main() {
	// 115. Distinct Subsequences
	test1()
	test2()
}
