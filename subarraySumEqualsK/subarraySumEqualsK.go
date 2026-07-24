package main

import "fmt"

func subarraySum(nums []int, k int) int {
	return checkSubarraySum_optimized(nums, k)
}

func checkSubarraySum_optimized(nums []int, k int) int {
	// this is an O(n) solution, just filling in the prefix sums

	// map of "sum to count"
	m := make(map[int]int)

	// !!! Any prefixSum = k should pair with sum = 0,
	// including the sums starting at a[0].
	// E.g. in array { 2 }, prefixSum[1] must count for k = 2
	m[0] = 1

	// !!! we don't need the prefixSums array, just the previous and the current sum
	currentSum := 0

	// total of subarrays summing up to K
	count := 0

	for _, v := range nums {
		currentSum = currentSum + v

		// we search for prefixSum[i] - prefixSum[prev] = k
		// -> prefixSum[prev] = prefixSum[i] - k
		targetEarlierSum := currentSum - k

		if earlierPrefixSumsEqualToTarget, ok := m[targetEarlierSum]; ok { // target sum already exists
			count += earlierPrefixSumsEqualToTarget
		}

		m[currentSum]++
	}

	return count
}

func test(arr []int, k int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("K: %v \n", k)

	result := subarraySum(arr, k)

	fmt.Printf("Total subarray that sum up to %v: %v \n", k, result)
	fmt.Printf("Expected result:                 %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{1, 1, 1}
	k := 2
	expected := 2

	test(arr, k, expected)
}

func test2() {
	arr := []int{2}
	k := 2
	expected := 1

	test(arr, k, expected)
}

func test3() {
	arr := []int{1, 2, 3}
	k := 3
	expected := 2

	test(arr, k, expected)
}

func main() {
	// 560. Subarray Sum Equals K
	test1()
	test2()
	test3()
}
