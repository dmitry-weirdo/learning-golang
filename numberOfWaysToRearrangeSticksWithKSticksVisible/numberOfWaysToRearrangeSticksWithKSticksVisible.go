package main

import "fmt"

const MOD = 1_000_000_007

func rearrangeSticks(n int, k int) int {
	// (n + 1) * (k + 1
	dp := make([][]int, n+1)

	for i := range n + 1 {
		dp[i] = make([]int, k+1)
	}

	var dfs func(n, k int) int

	dfs = func(n, k int) int {
		if n == 0 || k == 0 { // no sticks left or no positions left
			return 0
		}

		if n == k { // all sticks must be seen -> we need to rearrange ascending: 1, 2, 3 ... k=n
			return 1
		}

		// already pre-calculated -> return
		if dp[n][k] != 0 {
			return dp[n][k]
		}

		// case 1 - biggest stick is the last -> then for the (n - 1) previous sticks,
		// we solve (k - 1), since they won't be blocked by the biggest
		biggestStickIsLast := dfs(n-1, k-1) % MOD

		// case 1 - biggest stick is NOT  last
		// -> (n - 1) to select the last stick (excluding the biggest)
		// * solve for (n - 1) previous sticks, still k sticks, since the current last will be hidden by the current biggest that is in the first (n - 1) positions
		//biggestStickNotLast := (((n - 1) % mod) * (dfs(n-1, k) % mod)) % mod
		biggestStickNotLast := ((n - 1) * dfs(n-1, k)) % MOD // this is also working

		dp[n][k] = (biggestStickIsLast + biggestStickNotLast) % MOD
		return dp[n][k]
	}

	return dfs(n, k)
}

func test(n, k int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - number of sticks: %v \n", n)
	fmt.Printf("K - sticks to be visible: %v \n", k)

	result := rearrangeSticks(n, k)

	fmt.Printf("Count of put sticks [1 ; %v] so that %v sticks are visible (mod  %v): %v \n", n, k, MOD, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(3, 2, 3) // [1, 3, 2], [2, 3, 1], [2, 1, 3]
}

func test2() {
	test(5, 5, 1) // only [1, 2, 3, 4, 5]
}

func test3() {
	test(20, 11, 647427950)
}

func main() {
	// 1866. Number of Ways to Rearrange Sticks With K Sticks Visible
	test1()
	test2()
	test3()
}
