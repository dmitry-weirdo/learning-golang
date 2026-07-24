package main

import (
	"fmt"
)

func checkSubarraySum(nums []int, k int) bool {
	// this is O(n^2) solution
	// it gives TLE on big test-case 95

	n := len(nums)
	prefixSums := make([]int, n+1)

	// let's make correct prefix sums, with 0 and index = i + 1
	prefixSums[0] = 0

	for i, v := range nums {
		prefixSums[i+1] = prefixSums[i] + v
	}

	fmt.Printf("Prefix sums: %v \n", prefixSums)

	// Full comparison of prefix sums -O(n^2)
	for i := 0; i < n; i++ { // we leave 1 space for j, but prefix sums have N + 1 elements
		for j := i + 2; j < n+1; j++ { // we start with (i + 2) since subarrays of size 1 are not fitting
			subarraySum := prefixSums[j] - prefixSums[i]
			if subarraySum%k == 0 {
				fmt.Printf("prefixSums[%v] - prefixSums[%v] = %v - %v = %v. Returning true. \n", j, i, prefixSums[j], prefixSums[i], subarraySum)
				return true
			}
		}
	}

	return false
}

func test(arr []int, k int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("K: %v \n", k)

	result := checkSubarraySum(arr, k)

	fmt.Printf("There is a \"good\" subarray: %v \n", result)
	fmt.Printf("Expected result:            %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{0, 0}
	k := 1
	expected := true // 0 % 1 == 0

	test(arr, k, expected)
}

func main() {
	// 523. Continuous Subarray Sum
	test1()
}
