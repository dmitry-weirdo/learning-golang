package main

import (
	"fmt"
	"slices"
)

func threeSumSmaller(nums []int, target int) int {
	if len(nums) < 3 { // we need 3 distinct indexes!
		return 0
	}

	// external is similar to normal 3 sum
	slices.Sort(nums)

	sum := 0 // count of pairs

	for i := 0; i < len(nums)-2; i++ { // 2 remaining values are required to search values 2 and 3
		// we do NOT skip duplicates in this ticket, since we search for index pairs, NOT value pairs

		//if (i > 0) && (nums[i] == nums[i-1]) {
		//	// skip duplicate a[i] values
		//	continue
		//}

		// search with twoSum O(n) on the remaining array
		remainingTarget := target - nums[i]

		// search from the next element
		sum += twoSumSmaller(nums, i+1, remainingTarget)
	}

	return sum
}

func twoSumSmaller(arr []int, firstElementIndex int, target int) int {
	left := firstElementIndex
	right := len(arr) - 1

	// we're counting array index pairs that have sum < target
	sum := 0

	for left < right {
		if arr[left]+arr[right] >= target {
			// we cannot increase left since it will increase the sum
			// -> decrease right

			right--
		} else {
			// we will have (right - left) pairs that start with arr[left]
			// and end with value arr[left + 1] ... arr[right]

			// example: 1 2 3 5 8, target = 7
			// left = 0, right = 3, arr[left] = 1, arr[right] = 5
			// we will have (3 - 0) = 3 pairs that start with arr[left] = 1 and have sum < target:
			// [1, 2], [1, 3], [1, 5]

			sum += right - left
			left++
		}
	}

	return sum
}

func test(arr []int, target int, expectedResult int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Target sum: %v \n", target) // will always be 0

	result := threeSumSmaller(arr, target)

	fmt.Printf("Count of i, j, k index triplets that a[i] + a[j] + a[k] <  %v: \n%v \n", target, result)
	fmt.Printf("Expected result: \n%v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{-2, 0, 1, 3}
	target := 2
	expected := 2

	test(arr, target, expected)
}

func test2() {
	arr := []int{}
	target := 0
	expected := 0

	test(arr, target, expected)
}

func test3() {
	arr := []int{0}
	target := 0
	expected := 0

	test(arr, target, expected)
}

func test4() {
	arr := []int{0}
	target := 1
	expected := 0 // we need 3 different indexes

	test(arr, target, expected)
}

func main() {
	// 259. 3Sum Smaller
	test1()
	test2()
	test3()
	test4()
}
