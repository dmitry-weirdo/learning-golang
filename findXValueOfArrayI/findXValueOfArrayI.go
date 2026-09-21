package main

import (
	"fmt"
)

func resultArray(nums []int, k int) []int64 {

	n := len(nums)

	// todo: we can use just one array (instead of matrix) of dp for dp[i - 1] values

	// dp[i][r]
	// i - subarray ending on [i]
	// r - remainder of K
	// dp[i][r] - count of non-empty subarrays ending on [i] that have a (product of subarray values) % K == R
	dp := createIntMatrix(n, k)

	// we collect the sums of all remainders of K
	result := make([]int64, k)

	for i, v := range nums {
		// one subarray consists just of a[i]
		dp[i][v%k]++

		if i > 0 {
			// other subarrays ending on [i] are an addition of a[i] to arrays ending on dp[i - 1]
			for r := range k {
				// find the new remainder when we multiply every remainder R on a[i]
				newRemainder := (r * v) % k

				dp[i][newRemainder] += dp[i-1][r]
			}
		}

		// accumulate dp[i] for all R values into the result
		for r := range k {
			result[r] += int64(dp[i][r])
		}
	}

	return result
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func test(arr []int, k int, expectedResult []int64) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("K (divisor of the product): %v \n", k)

	result := resultArray(arr, k)

	fmt.Printf("Count of sub-arrays with products mod %v: %v \n", k, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	test(
		[]int{1, 2, 3, 4, 5},
		3,
		[]int64{9, 2, 4},
	)
}

func test2() {
	test(
		[]int{1, 2, 4, 8, 16, 32},
		4,
		[]int64{18, 1, 2, 0},
	)
}

func test3() {
	test(
		[]int{1, 1, 2, 1, 1},
		2,
		[]int64{9, 6},
	)
}

func main() {
	// 3524. Find X Value of Array I
	test1()
	test2()
	test3()
}
