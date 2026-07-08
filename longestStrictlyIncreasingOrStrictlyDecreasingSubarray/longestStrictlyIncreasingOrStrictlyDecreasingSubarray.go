package main

import "fmt"

func longestMonotonicSubarray(nums []int) int {
	n := len(nums)

	// find max increasing length
	maxIncreasingLength := 1

	currentLength := 1
	i := 1

	for i < n {
		for (i < n) && (nums[i] > nums[i-1]) {
			i++
			currentLength++
		}

		maxIncreasingLength = max(currentLength, maxIncreasingLength)

		// reset increasing sequence start
		currentLength = 1
		i++
	}

	fmt.Printf("Max increasing subarray length: %d \n", maxIncreasingLength)

	// find max decreasing length
	maxDecreasingLength := 1

	currentLength = 1
	i = 1

	for i < n {
		for (i < n) && (nums[i] < nums[i-1]) {
			i++
			currentLength++
		}

		maxDecreasingLength = max(currentLength, maxDecreasingLength)

		// reset decreasing sequence start
		currentLength = 1
		i++
	}

	fmt.Printf("Max decreasing subarray length: %d \n", maxIncreasingLength)

	return max(maxIncreasingLength, maxDecreasingLength)
}

func longestMonotonicSubarrayNaive(nums []int) int {
	n := len(nums)

	// find max increasing length
	maxIncreasingLength := 1

	start := 0
	i := 1

	for i < n {
		for (i < n) && (nums[i] > nums[i-1]) {
			i++
		}

		newLength := i - start // i is already an element that broke the sequence or is out of an array

		maxIncreasingLength = max(newLength, maxIncreasingLength)

		// reset increasing sequence start
		start = i
		i++
	}

	fmt.Printf("Max increasing subarray length: %d \n", maxIncreasingLength)

	// find max decreasing length
	maxDecreasingLength := 1

	start = 0
	i = 1

	for i < n {
		for (i < n) && (nums[i] < nums[i-1]) {
			i++
		}

		newLength := i - start // i is already an element that broke the sequence or is out of an array

		maxDecreasingLength = max(newLength, maxDecreasingLength)

		// reset decreasing sequence start
		start = i
		i++
	}

	fmt.Printf("Max decreasing subarray length: %d \n", maxIncreasingLength)

	return max(maxIncreasingLength, maxDecreasingLength)
}

func test(a []int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array 1: %v \n", a)

	result := longestMonotonicSubarray(a)

	fmt.Printf("Max monotonic subarray length: %v \n", result)
	fmt.Printf("Expected max monotonic subarray length: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{1, 4, 3, 3, 2}
	expected := 2

	test(arr, expected)
}

func test2() {
	arr := []int{3, 3, 3, 3}
	expected := 1

	test(arr, expected)
}

func main() {
	// 3105. Longest Strictly Increasing or Strictly Decreasing Subarray
	test1()
	test2()
}
