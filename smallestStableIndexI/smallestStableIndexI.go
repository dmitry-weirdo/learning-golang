package main

import "fmt"

func firstStableIndex(nums []int, k int) int {
	// O(3 * n) = O(n)
	// n <= 100, so this all is very fast and 0 ms

	n := len(nums)

	// count max from left to right
	maxValue := nums[0]
	maxValues := make([]int, n)

	for i := 0; i < n; i++ {
		maxValue = max(maxValue, nums[i])
		maxValues[i] = maxValue
	}

	// count min from right to left
	minValue := nums[n-1]
	minValues := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		minValue = min(minValue, nums[i])
		minValues[i] = minValue
	}

	// find the first stable index
	for i := range n {
		if maxValues[i]-minValues[i] <= k {
			return i
		}
	}

	return -1
}

func test(arr []int, k int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("K: %v \n", k)

	result := firstStableIndex(arr, k)

	fmt.Printf("First stable index: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{5, 0, 1, 4},
		3,
		3,
	)
}

func test2() {
	test(
		[]int{3, 2, 1},
		1,
		-1,
	)
}

func test3() {
	test(
		[]int{0},
		0,
		0,
	)
}

func main() {
	// 3903. Smallest Stable Index I
	test1()
	test2()
	test3()
}
