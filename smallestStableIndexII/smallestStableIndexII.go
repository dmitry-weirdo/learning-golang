package main

import "fmt"

func firstStableIndex(nums []int, k int) int {
	// O(2*n) = O(n)
	// Let's do 2 * N instead of 3 * N because of bigger constraint.
	n := len(nums)

	// count max from left to right
	maxValue := nums[0]
	maxValues := make([]int, n)

	// count min from right to left
	minValue := nums[n-1]
	minValues := make([]int, n)

	for i := range n {
		maxValue = max(maxValue, nums[i])
		maxValues[i] = maxValue

		minValue = min(minValue, nums[n-1-i])
		minValues[n-1-i] = minValue
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
	// 3904. Smallest Stable Index II
	// Same as "3903. Smallest Stable Index I", just with bigger N constraint: 10^5 instead of 100.
	test1()
	test2()
	test3()
}
