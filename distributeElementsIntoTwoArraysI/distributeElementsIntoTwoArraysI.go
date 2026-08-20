package main

import (
	"fmt"
	"slices"
)

func resultArray(nums []int) []int {
	// allocates just a single result array, no
	// so it's O(n) space if we're counting the result array,
	// or O(1) space if we're NOT counting the result array.
	// Time in this task is always O(n)
	return resultArray_singleResultArray(nums)
}

func resultArray_singleResultArray(nums []int) []int {
	n := len(nums)

	result := make([]int, n)

	result[0] = nums[0]
	result[n-1] = nums[1]

	// left goes left -> right
	left := 0

	// right goes right -> left
	right := n - 1

	for i := 2; i < n; i++ {
		if result[left] > result[right] {
			left++
			result[left] = nums[i]
		} else {
			right--
			result[right] = nums[i]
		}
	}

	// reverse the part [result; n-1]
	slices.Reverse(result[right:])

	return result
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)

	result := resultArray(arr)

	fmt.Printf("Transformed distributed array: %v \n", result)
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
		[]int{2, 1, 3},
		[]int{2, 3, 1},
	)
}

func test2() {
	test(
		[]int{5, 4, 3, 8},
		[]int{5, 3, 4, 8},
	)
}

func main() {
	// 3069. Distribute Elements Into Two Arrays I
	test1()
	test2()
}
