package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	return slidingWindow(target, nums)
}

func slidingWindow(target int, nums []int) int {
	// typical sliding window, similar to "76. Minimum Window Substring"
	// O(n)
	left := 0

	sum := 0

	minWindowSize := len(nums) + 1

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		// shrink from left
		for sum >= target {
			windowSize := right - left + 1
			minWindowSize = min(minWindowSize, windowSize)

			sum -= nums[left]
			left++
		}
	}

	// no matching window found -> return 0
	if minWindowSize > len(nums) {
		return 0
	}

	return minWindowSize
}

func test(arr []int, target int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Target sum: %v \n", target)

	result := minSubArrayLen(target, arr)

	fmt.Printf("Length of minimum subarray whose sum >= %v: %v \n", target, result)
	fmt.Printf("Expected result:                           %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{2, 3, 1, 2, 4, 3}
	target := 7
	expected := 2 // {4, 3}

	test(arr, target, expected)
}

func test2() {
	arr := []int{1, 4, 4}
	target := 4
	expected := 1 // {4}

	test(arr, target, expected)
}

func test3() {
	arr := []int{1, 1, 1, 1, 1, 1, 1, 1}
	target := 11
	expected := 0 // no solution, array is not enough

	test(arr, target, expected)
}

func main() {
	// 209. Minimum Size Subarray Sum
	test1()
	test2()
	test3()
}
