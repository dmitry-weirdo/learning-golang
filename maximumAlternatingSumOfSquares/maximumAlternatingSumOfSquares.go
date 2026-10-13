package main

import (
	"fmt"
	"slices"
)

func maxAlternatingSum(nums []int) int64 {
	// sort by squares desc
	slices.SortFunc(nums, func(a, b int) int {
		return b*b - a*a
	})

	//fmt.Printf("Array sorted by squares desc: %v \n", nums)

	n := len(nums)

	if n == 1 {
		return int64(nums[0] * nums[0])
	}

	result := 0

	i := 0

	for i = range n / 2 {
		result += nums[i] * nums[i]
	}

	if (n%2 == 1) && (i < n) {
		i++
		result += nums[i] * nums[i]
	}

	for i < n-1 {
		i++
		result -= nums[i] * nums[i]
	}

	return int64(result)
}

func test(arr []int, expectedResult int64) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := maxAlternatingSum(arr)

	fmt.Printf("Max alternating sum of squares: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{1, 2, 3}, 12)
}

func test2() {
	test([]int{1, -1, 2, -2, 3, -3}, 16)
}

func test3() {
	test([]int{-6}, 36)
}

func test4() {
	test([]int{-1, -2}, 3)
}

func main() {
	// 3727. Maximum Alternating Sum of Squares
	test1()
	test2()
	test3()
	test4()
}
