package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	// todo: this can be solved using O(n) / O(m) space, if we're using a 1D-memo optimized DP.

	// O(m * n) time - all indices in S and T
	// O(m * n) space to store memo - non-optimal
	return isInterleave_recursive(s1, s2, s3)
}

func isInterleave_recursive(s string, t string, r string) bool {
	if len(s)+len(t) != len(r) {
		return false
	}

	if len(s) <= 0 { // S is empty
		return t == r
	}

	if len(t) <= 0 { // T is empty
		return s == r
	}

	// i - index in S
	// j - index in T
	// k - index in R

	// memo -> caches result for [i][j] pairs
	// 0 - undefined
	// 1 - solution impossible for [i][j]
	// 2 - solution possible for [i][j]
	const MEMO_UNDEFINED = 0
	const MEMO_FALSE = 1
	const MEMO_TRUE = 2

	memo := createIntMatrix(len(s), len(t)) // will be filed with 0

	var dfs func(i, j, k int) bool

	dfs = func(i, j, k int) bool {
		if i >= len(s) { // exhausted S -> check if the remaining part of T matches the remaining part of R
			return t[j:] == r[k:]
		}

		if j >= len(t) { // exhausted T -> check if the remaining part of S matches the remaining part of R
			return s[i:] == r[k:]
		}

		if memo[i][j] != MEMO_UNDEFINED { // result for [i][j] within the string lengths already calculated -> return it
			return memo[i][j] == MEMO_TRUE
		}

		result := false

		if ((s[i] == r[k]) && dfs(i+1, j, k+1)) || // next S char matches s[i] == r[k] -> +1 char in S and R
			((t[j] == r[k]) && dfs(i, j+1, k+1)) { // next T char matches t[j] == r[k] -> +1 char in T and R
			result = true
		}

		// cache the result in memo matrix
		if result {
			memo[i][j] = MEMO_TRUE
		} else {
			memo[i][j] = MEMO_FALSE
		}

		return result
	}

	dfs(0, 0, 0)

	return memo[0][0] == MEMO_TRUE
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func test(s, t, r string, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String S: %v \n", s)
	fmt.Printf("String T: %v \n", t)
	fmt.Printf("Result string R: %v \n", r)

	result := isInterleave(s, t, r)

	fmt.Printf("Result string is interleave of strings S and T: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("aabcc", "dbbca", "aadbbcbcac", true) // aa + dbbc + bc + a + c
}

func test2() {
	test("aabcc", "dbbca", "aadbbbaccc", false)
}

func test3() {
	test("", "", "", true)
}

func main() {
	// 97. Interleaving String
	test1()
	test2()
	test3()
}
