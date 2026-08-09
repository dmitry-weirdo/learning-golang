package main

import "fmt"

func stoneGameII(piles []int) int {
	// prefixSums help us track the remaining sum of any range (namely from current index included in the sum to the end of the array).
	prefixSums := getPrefixSums(piles)
	//fmt.Printf("Prefix sums: %v \n", pS)

	n := len(piles)

	// memo matrix, i = index, j = m
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	var dfs func(i, m int) int

	dfs = func(i, m int) int {
		// last array index is [n - 1]
		// last prefix sum (whole array) = prefixSums[n]
		// we take the current a[i] as well,
		// i.e. we need to subtract the sum without a[i], i.e. prefixSums[i]
		remainingSum := prefixSums[n] - prefixSums[i]

		// Base case -> if we can take all remaining stones (2*M or fewer), take them.
		// Actually, it is (i + 2*m - 1) >= (n - 1),
		// since i is inclusive in the current turn, and array index goes up to (N - 1)
		// But we can remove -1 from both sides.
		if i+2*m >= n {
			return remainingSum
		}

		// if value is already calculated -> return it
		if dp[i][m] != 0 {
			return dp[i][m]
		}

		// we try all the values of X in range [1; 2*M] and get the maximum stones for the current player
		maxSum := 0

		for x := 1; x <= 2*m; x++ { // try to take x piles -> next index will be (i + x)
			// Main trick:
			// of the remaining sum in range [i; n-1] = (prefixSums[n] - prefixSums[i]),
			// we subtract the optimal game of the opponent for current X and accordingly updated M
			nextIndex := i + x
			nextM := max(x, m)
			opponentScore := dfs(nextIndex, nextM)

			ourScore := remainingSum - opponentScore

			maxSum = max(maxSum, ourScore)
		}

		dp[i][m] = maxSum
		return maxSum
	}

	return dfs(0, 1) // m starts with 1, i starts with 0 (start of the array)
}

func getPrefixSums(a []int) []int {
	// with prefixSums[0] = 0, sum up to a[i] will be in prefixSums[i+1]
	prefixSums := make([]int, len(a)+1)

	prefixSums[0] = 0

	for i, v := range a {
		prefixSumIndex := i + 1
		prefixSums[prefixSumIndex] = prefixSums[i] + v
	}

	return prefixSums
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Piles of stones: %v \n", arr)

	result := stoneGameII(arr)

	fmt.Printf("Alice (1st player) max stones: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{2, 7, 9, 4, 4}
	expected := 10

	test(arr, expected)
}

func test2() {
	arr := []int{1, 2, 3, 4, 5, 100}
	expected := 104

	test(arr, expected)
}

func test3() {
	arr := []int{1}
	expected := 1

	test(arr, expected)
}

func main() {
	// 1140. Stone Game II
	test1()
	test2()
	test3()
}
