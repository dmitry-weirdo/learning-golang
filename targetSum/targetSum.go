package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	return findTargetSumWays_dp(nums, target)
}

func findTargetSumWays_dp(nums []int, target int) int {
	// all values are >= 0
	// Therefore, the range of the possible sums is [-totalSum; totalSum]

	// Target can be negative

	// dp[i][j] - using numbers [0..i], count of ways to achieve sum j, using these numbers
	// j is in range [-totalSum; totalSum], i.e. size is (2 * totalSum + 1)

	// j index for some sum is gotten like this:
	// sum = -totalSum -> dp[i][0]
	// sum = 0 -> dp[i][totalSum]
	// sum = totalSum -> dp[i][2 * totalSum]
	// i.e. we have to add totalSum to the index [j]

	// For the first value, we can only get sums of a[0] and -a[0]

	// in the last

	totalSum := sumOfArray(nums)

	if abs(target) > totalSum { // impossible to reach a too big target sum
		return 0
	}

	n := len(nums)

	dp := createIntMatrix(n, 2*totalSum+1)

	// fill for i = 0 -> we can only use 1 number a[0]
	dp[0][nums[0]+totalSum] += 1  // 0 should be counted as both +0 and -0
	dp[0][-nums[0]+totalSum] += 1 // 0 should be counted as both +0 and -0

	// fill for indices from 1 to (n - 1)
	for i := 1; i < n; i++ {
		for j := range 2*totalSum + 1 {
			// if there was a possibility of some sum in [i-1], we can add or subtract the nums[i]
			// so the DP will increase if we can get the same sum with different ways
			if dp[i-1][j] > 0 {
				dp[i][j+nums[i]] += dp[i-1][j] // we can add nums[i]
				dp[i][j-nums[i]] += dp[i-1][j] // we can subtract nums[i]
			}
		}
	}

	return dp[n-1][target+totalSum]
}

func sumOfArray(arr []int) int {
	// we assume the array is non-empty
	sum := 0

	for _, v := range arr {
		sum += v
	}

	return sum
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func test(arr []int, target int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Target sum: %v \n", target)

	result := findTargetSumWays(arr, target)

	fmt.Printf("Number of (+/-) expressions evaluating to %v: %v \n", target, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{1, 1, 1, 1, 1}, 3, 5)
}

func test2() {
	test([]int{1}, 1, 1)
}

func test3() {
	// failing test-case 137/146
	test([]int{0}, 0, 2) // we count +0 and -0 as different results!
}

func main() {
	// 494. Target Sum
	test1()
	test2()
	test3()
}
