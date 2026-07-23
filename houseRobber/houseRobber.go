package main

import "fmt"

func rob(nums []int) int {
	return dpFlat(nums)
	//return dpWithMemoization(nums)
}

func dpFlat(nums []int) int {
	// todo: we can also probably go from houses -1 and 0
	// dp[len] = 0
	// dp[len-1] = nums[len-1]
	l := len(nums)

	// dp will just keep i + 1, i + 2 values
	memo := [2]int{nums[l-1], 0}

	for i := l - 2; i >= 0; i-- {
		// take current house money - current and memo[i+2] -> memo[1]
		takeCurrent := nums[i] + memo[1]

		// skip current house - memo[i+1] -> memo[0]
		skipCurrent := memo[0]

		memo[1] = memo[0]                       // i + 1
		memo[0] = max(takeCurrent, skipCurrent) // i
	}

	return memo[0]
}

func dpWithMemoization(nums []int) int {
	l := len(nums)
	memo := make([]int, l)

	for i := range memo {
		memo[i] = -1 // be safe since nums[i] can be 0
	}

	// just 1 house left -> it's optimal to take it
	memo[l-1] = nums[l-1]

	return dp(nums, memo, 0)
}

func dp(a []int, memo []int, i int) int {
	if i >= len(a) {
		return 0
	}

	if memo[i] >= 0 {
		return memo[i]
	}

	// take current house money
	takeCurrent := a[i] + dp(a, memo, i+2)

	// skip current house
	skipCurrent := dp(a, memo, i+1)

	memo[i] = max(takeCurrent, skipCurrent)
	return memo[i]
}

func test(arr []int, expectedResult int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array of houses: %v \n", arr)

	result := rob(arr)

	fmt.Printf("Max robbery sum: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{1, 2, 3, 1}
	expected := 4

	test(arr, expected)
}

func test2() {
	arr := []int{2, 1, 1, 2}
	expected := 4

	test(arr, expected)
}

func test3() {
	arr := []int{2, 7, 9, 3, 1}
	expected := 12

	test(arr, expected)
}

func main() {
	// 198. House Robber
	test1()
	test2()
	test3()
}
