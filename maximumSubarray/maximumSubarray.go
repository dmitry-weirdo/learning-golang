package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	return kadane(nums)
	// todo: divide and conquer solution of O(n * log n) ???
}

func kadane(nums []int) int {
	// O(n) time
	if len(nums) < 1 { // safety-measure, this must never happen
		return 0
	}

	maxSum := nums[0]

	// currentSum is the maximum of the arrays that end on the previous element
	// if the best currentSum before the current element is negative,
	// we're just using the current element
	currentSum := 0

	for _, v := range nums {
		// if previous values sum up to negative, we don't take them, and previous sum is 0
		currentSum = max(currentSum, 0)

		// add the current number
		currentSum += v

		// previous 2 lines can be written as this (either take current value and previous array ending on the previous element, or just take the current element)
		//currentSum = max(v, v + currentSum)

		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func kadaneWithBorders(nums []int) (maxLeft int, maxRight int, maxSum int) {
	maxSum = nums[0]
	currentSum := 0

	maxLeft = 0
	maxRight = 0

	left := 0

	for right := 0; right < len(nums); right++ {
		if currentSum < 0 { // new optimal subarray will start with just the current element, i.e. we move the left
			currentSum = 0
			left = right
		}

		currentSum += nums[right]

		if currentSum > maxSum {
			maxSum = currentSum
			maxLeft = left
			maxRight = right
		}
	}

	return maxLeft, maxRight, maxSum
}

func test(arr []int, expectedResult int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array: %v \n", arr)

	result := maxSubArray(arr)

	left, right, _ := kadaneWithBorders(arr)
	subArrayWithMaxSum := arr[left : right+1]

	fmt.Printf("Maximum subarray sum:   %v \n", result)
	fmt.Printf("Maximum subarray:       %v \n", subArrayWithMaxSum)
	fmt.Printf("Expected result:        %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
		return
	}
}

func test1() {
	arr := []int{-10, -1, -3}
	expected := -1

	test(arr, expected)
}

func test2() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6

	test(arr, expected)
}

func test3() {
	arr := []int{1}
	expected := 1

	test(arr, expected)
}

func test4() {
	arr := []int{5, 4, -1, 7, 8}
	expected := 23

	test(arr, expected)
}

func main() {
	// 53. Maximum Subarray
	test1()
	test2()
	test3()
	test4()
}
