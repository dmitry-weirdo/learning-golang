package main

import "fmt"

func kConcatenationMaxSum(arr []int, k int) int {
	// we should have a max of:
	// 0 (if all elements are negative) -> we take the empty subarray
	// k * sum(array)
	// maxSubarray (usual Kadane / prefixSum)
	// maxPrefixSum + maxSuffixSum - end of one array + start of the copy  - !!! only if K > 1
	// max(maxPrefixSum, 0) + (k - 2) * totalSum + max(maxSuffixSum, 0) !!! only if K > 2

	maxSubarraySum, maxPrefixSum, totalSum := prefixSums(arr)
	maxPrefixSum = max(0, maxPrefixSum) // we can take 0 values from the start

	maxSuffixSum := getMaxSuffixSum(arr)
	maxSuffixSum = max(0, maxSuffixSum) // we can take 0 values from the end

	if maxSubarraySum < 0 { // all array values are negative -> selecting an empty subarray is the best option
		return 0
	}

	sumOfKArrays := k * totalSum

	prefixAndSuffix := -1 // less than 0
	if k > 1 {
		prefixAndSuffix = maxPrefixSum + maxSuffixSum
	}

	prefixAndSuffixAndCenter := -1
	if k > 2 {
		prefixAndSuffixAndCenter = maxPrefixSum + maxSuffixSum + (k-2)*totalSum
	}

	maxValue := getMaxValue(0, sumOfKArrays, maxSubarraySum, prefixAndSuffix, prefixAndSuffixAndCenter)
	const mod = 1_000_000_007

	return maxValue % mod
}

func prefixSums(nums []int) (maxSubarraySum int, maxPrefixSum int, totalSum int) { // extended version of a function from "53. Maximum Subarray"
	// prefixSums[i+1] - prefixSums[j] = sum of subarray[i; j]
	prefixSum := 0
	minPrefixSum := 0
	totalSum = 0

	maxSubarraySum = nums[0]
	maxPrefixSum = nums[0]

	for _, v := range nums {
		prefixSum += v
		maxPrefixSum = max(maxPrefixSum, prefixSum)
		totalSum += v // we can also use prefixSum

		maxSubarraySum = max(maxSubarraySum, prefixSum-minPrefixSum)

		minPrefixSum = min(minPrefixSum, prefixSum)
	}

	return maxSubarraySum, maxPrefixSum, totalSum
}

func getMaxSuffixSum(a []int) int {
	n := len(a)

	suffixSum := 0

	maxSuffixSum := a[n-1]

	for i := n - 1; i >= 0; i-- {
		suffixSum += a[i]

		maxSuffixSum = max(maxSuffixSum, suffixSum)
	}

	return maxSuffixSum
}

func getMaxValue(values ...int) int { // wrapper for varargs instead of array
	return maxInArray(values)
}

func maxInArray(arr []int) int {
	if len(arr) == 0 {
		panic("maxInArray: no values provided")
	}

	// we assume the array is non-empty
	m := arr[0]

	for _, v := range arr {
		m = max(m, v)
	}

	return m
}

func test(arr []int, k int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("K (count of repeating the array): %v \n", k)

	result := kConcatenationMaxSum(arr, k)

	fmt.Printf("Max sub-array sum of array repeated %v times: %v \n", k, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{1, 2}, // (1 + 2) * 3 times
		3,
		9,
	)
}

func test2() {
	test(
		[]int{1, -2, 1},
		5,
		2, // just 1 + 1
	)
}

func test3() {
	test(
		[]int{-1, -2},
		7,
		0, // we can have an empty sub-array
	)
}

func test4() {
	// failing test-case 28/44
	test(
		[]int{-5, 4, -4, -3, 5, -3},
		3,
		5, // 5 (subarray)
	)
}

func test5() {
	// failing test-case 30/44
	test(
		[]int{-5, -2, 0, 0, 3, 9, -2, -5, 4},
		5,
		20,
	)
}

func main() {
	// 1191. K-Concatenation Maximum Sum
	test1()
	test2()
	test3()
	test4()
	test5()
}
