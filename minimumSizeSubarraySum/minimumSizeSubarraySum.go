package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	return calculateWithPrefixSums(target, nums)
	//return slidingWindow(target, nums)
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

func calculateWithPrefixSums(target int, nums []int) int {
	// this is O(n * log n) solution:
	// O(n) - calculate prefix sums
	// For every i in [0; n- 1], we execute the binary search that is O(log n)
	// -> O(n * log n)

	n := len(nums)

	// calculate prefixSums[0; n]
	prefixSums := make([]int, n+1)

	prefixSums[0] = 0

	for i, v := range nums {
		prefixSums[i+1] = prefixSums[i] + v
	}

	fmt.Printf("Prefix sums: %v \n", prefixSums)

	// !!! since all values in arr are positive, we know that prefixSums is strictly increasing,
	// i.e. sorted, and we can binary-search in it with O(log n)

	// to calculate the sum of the subarray [i; j] (inclusive), the formula is ps[j + 1] - ps[i]

	minSubarraySize := len(nums) + 1

	for i := 0; i < n; i++ {
		// for every a[i], we search the smallest a[j] so that
		// prefixSums[j + 1] - prefixSums[i] >= target

		// this is a binary search starting from i
		// !!! Note that we return j that is in [i; n] range, not the prefixSum range
		minJ := binarySearch(prefixSums, i, target)

		// if j >= n then insert position is after the array -> j is not found
		if minJ < n { // j = n - 1 -> prefixSums[n] is still valid
			subarraySize := minJ - i + 1
			minSubarraySize = min(minSubarraySize, subarraySize)
		}
	}

	// no matching subarray found -> return 0
	if minSubarraySize > len(nums) {
		return 0
	}

	return minSubarraySize
}

func binarySearch(ps []int, leftStart int, target int) int {
	left := leftStart

	// but since we're using ps[mid + 1], this is set to (N + 1), NOT (N + 2)
	// I.e. in case of insertion after the array, we will return N + 1 that is out of bounds for the original array of size N
	right := len(ps) - 1 // insertion point can be after the array, len is (N + 1)

	for left < right {
		mid := (left + right) / 2

		// !!! Notice mid + 1, so that we'll return j, NOT j + 1
		if ps[mid+1]-ps[leftStart] >= target { // target condition
			right = mid
		} else {
			left = mid + 1
		}
	}

	// returned is the leftmost value where condition is satisfied
	return left
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
	//test1()
	//test2()
	test3()
}
