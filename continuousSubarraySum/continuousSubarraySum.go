package main

import (
	"fmt"
)

func checkSubarraySum(nums []int, k int) bool {
	return checkSubarraySum_optimized_2(nums, k) // O(n), O(k) space, just on counting the prefixSum values, storing just the "remainder to first position" map
	//return checkSubarraySum_optimized(nums, k) // O(n), O(n) + O(k) space just on counting the prefixSum values
	//return checkSubarraySum_naive(nums, k) // O(n^2), fails on TLE
}

func checkSubarraySum_optimized_2(nums []int, k int) bool {
	// this is an O(n) solution, just filling in the prefix sums
	// this is space-optimized to O(k) just for the map -> we don't need to keep the prefixSums array, we're just keeping the remainders

	// todo: for the faster/less space in Go, we can use an int[] of size k for all the remainder in [0; k - 1] range
	// map of "remainder by K" to "first prefixSum position with this remainder"
	m := make(map[int]int)

	// let's make correct prefix sums, with 0 and index = i + 1
	// !!! we don't need the prefixSums array, just the previous and the current sum
	currentSum := 0
	m[0] = 0

	for i, v := range nums {
		currentSum = currentSum + v

		remainder := currentSum % k

		if remainderFirstIndex, ok := m[remainder]; ok { // remainder already exists
			if ((i + 1) - remainderFirstIndex) >= 2 {
				return true
			}
		} else { // remainder found the first time -> put the first index into the map
			m[remainder] = i + 1
		}
	}

	return false
}

func checkSubarraySum_optimized(nums []int, k int) bool {
	// this is an O(n) solution, just filling in the prefix sums

	n := len(nums)
	prefixSums := make([]int, n+1)

	// map of "remainder by K" to "first prefixSum position with this remainder"
	m := make(map[int]int)

	// let's make correct prefix sums, with 0 and index = i + 1
	prefixSums[0] = 0
	m[0] = 0

	for i, v := range nums {
		prefixSums[i+1] = prefixSums[i] + v

		remainder := prefixSums[i+1] % k

		if remainderFirstIndex, ok := m[remainder]; ok { // remainder already exists
			if ((i + 1) - remainderFirstIndex) >= 2 {
				return true
			}
		} else { // remainder found the first time -> put the first index into the map
			m[remainder] = i + 1
		}
	}

	fmt.Printf("Prefix sums: %v \n", prefixSums)

	return false
}

func checkSubarraySum_naive(nums []int, k int) bool {
	// this is an O(n^2) solution
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

func test2() {
	arr := []int{23, 2, 4, 6, 7}
	k := 6
	expected := true

	test(arr, k, expected)
}

func test3() {
	arr := []int{23, 2, 6, 4, 7}
	k := 6
	expected := true

	test(arr, k, expected)
}

func test4() {
	arr := []int{23, 2, 6, 4, 7}
	k := 13
	expected := false

	test(arr, k, expected)
}

func test5() {
	arr := []int{5, 0, 0, 0}
	k := 3
	expected := true

	test(arr, k, expected)
}

func main() {
	// 523. Continuous Subarray Sum
	test1()
	test2()
	test3()
	test4()
	test5()
}
