package main

import (
	"cmp"
	"fmt"
	"slices"
)

func largestPerimeter(nums []int) int {
	// a >= b >= c
	// It muse be a < b + c
	// We take 3 consecutive elements. If (b + c) are not enough,
	// then any elements smaller than B and C are not enough as well.

	// sort desc
	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	for i := range len(nums) - 2 {
		if nums[i] < nums[i+1]+nums[i+2] { // strict less, to exclude the 0-square triangle
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}

	return 0
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := largestPerimeter(arr)

	fmt.Printf("Largest triangle perimeter: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{2, 1, 2}, 5)
}

func test2() {
	test([]int{1, 2, 1, 10}, 0) // cannot find any non-0-square triangle
}

func main() {
	// 976. Largest Perimeter Triangle
	test1()
	test2()
}
