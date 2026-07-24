package main

import "fmt"

func maxSubarraySumCircular(nums []int) int {
	return prefixSums(nums)
}

func prefixSums(nums []int) int {
	// heuristic is: if we have a circular sum, the remaining internal subarray has sum of totalArraySum - maxCircularSum
	// Therefore, to get maxCircularSum, we need to minimize the internal subarray sum

	// prefixSums[i] - prefixSums[j] = sum of subarray[i; j]
	prefixSum := 0
	minPrefixSum := 0
	maxPrefixSum := 0

	minSubarraySum := nums[0]
	maxSubarraySum := nums[0]

	for _, v := range nums {
		prefixSum += v

		// calcular the minimum internal subarray to get the circularMaxSum
		minSubarraySum = min(minSubarraySum, prefixSum-maxPrefixSum)

		// usual Kadane sum, no circular
		maxSubarraySum = max(maxSubarraySum, prefixSum-minPrefixSum)

		minPrefixSum = min(minPrefixSum, prefixSum)
		maxPrefixSum = max(maxPrefixSum, prefixSum)
	}

	// prefixSum is now the sum
	totalArraySum := prefixSum
	circularMaxSum := totalArraySum - minSubarraySum

	fmt.Printf("Total array sum: %v \n", totalArraySum)
	fmt.Printf("Min subarray sum (non-circular): %v \n", minSubarraySum)
	fmt.Printf("Max subarray sum (non-circular): %v \n", maxSubarraySum)

	if minSubarraySum == totalArraySum {
		// corner-case -> if minSubarraySum is the complete array (i.e. the complete array is negative),
		// then the circularMaxSum will be 0. We shouldn't count this circular result.
		// Just return the single element solution of usual non-circular subarray
		return maxSubarraySum
	}

	return max(maxSubarraySum, circularMaxSum)
}

func test(arr []int, expectedResult int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array: %v \n", arr)

	result := maxSubarraySumCircular(arr)

	fmt.Printf("Maximum circular subarray sum:   %v \n", result)
	fmt.Printf("Expected result:                 %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
		return
	}
}

func test1() {
	arr := []int{-3, -2, -3} // all-negatives case
	expected := -2           // circular does not apply in the non-negative case

	test(arr, expected)
}

func test2() {
	arr := []int{1, -2, 3, -2}
	expected := 3 // non-circular wins

	test(arr, expected)
}

func test3() {
	arr := []int{5, -3, 5}
	expected := 10 // circular wins

	test(arr, expected)
}

func main() {
	// 918. Maximum Sum Circular Subarray
	test1()
	test2()
	test3()
}
