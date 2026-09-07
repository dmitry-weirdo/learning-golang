package main

import "fmt"

func distinctSubseqII(s string) int {
	const mod = 1_000_000_007
	n := len(s)

	// dp[i] is count of subsequences using characters s[0]...s[i]
	dp := createIntArrayWithDefaultValues(n+1, -1)
	dp[0] = 1 // just empty string

	// link to the previous dp[i] for s[i] = char
	last := createIntArrayWithDefaultValues(26, -1)

	for i := range n {
		ch := s[i] - 'a'

		dp[i+1] = (dp[i] * 2) % mod

		if last[ch] >= 0 {
			// Remove duplicated subsequences: previous subsequences that end with current character.
			dp[i+1] -= dp[last[ch]]
		}

		dp[i+1] %= mod

		// Save the index of last "Number of subsequences that end with current character"
		// Index in dp array.
		last[ch] = i
	}

	dp[n]-- // we do not count empty string

	if dp[n] < 0 { // avoid negative values
		dp[n] += mod
	}

	return dp[n]
}

func createIntArrayWithDefaultValues(n int, defaultValue int) []int {
	a := make([]int, n)

	for i := range n {
		a[i] = defaultValue
	}

	return a
}

func test(s string, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := distinctSubseqII(s)

	fmt.Printf("Count of distinct subsequences: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("abc", 7) // a, b, c, ab, ac, bc, abc
}

func test2() {
	test("aba", 6) // a, b, ab, aa, ba, aba
}

func test3() {
	test("aaa", 3) // a, aa, aaa
}

func main() {
	// 940. Distinct Subsequences II
	test1()
	test2()
	test3()
}
