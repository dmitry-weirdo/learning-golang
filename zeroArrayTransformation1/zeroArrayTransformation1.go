package main

import (
	"fmt"
	"unsafe"
)

func isZeroArray(nums []int, queries [][]int) bool {
	return isZeroArray_optimized(nums, queries)
	//return isZeroArray_old(nums, queries)
}

func isZeroArray_optimized(nums []int, queries [][]int) bool {
	// we don't need an array of prefixSums, we just need the current sum
	l := len(nums)

	diff := make([]int, l+1) // can be + 1, after the last element

	// fill the diff array ( +1 at [query start], - 1 at [query end + 1] )
	// diff will affect +1 for all prefix sums of nums[left] to nums[right]
	// and then we remove this effect starting from nums[right + 1]
	for _, v := range queries {
		left := v[0]
		right := v[1]

		diff[left]++
		diff[right+1]-- // note it's right + 1, because the range is inclusive and ends only on the next index
	}

	// prefixSums[0] = 0
	prefixSum := 0

	for i, v := range nums {
		// prefixSum[i + 1] contains diff[i]
		prefixSum += diff[i]

		if prefixSum < v {
			return false
		}
	}

	return true
}

func isZeroArray_old(nums []int, queries [][]int) bool {
	fmt.Printf("Int size: %v \n ", unsafe.Sizeof(int(0))*8)

	l := len(nums)

	diff := make([]int, l+1) // can be + 1, after the last element

	// fill the diff array ( +1 at [query start], - 1 at [query end + 1] )
	for _, v := range queries {
		left := v[0]
		right := v[1]

		diff[left]++
		diff[right+1]-- // note it's right + 1, because the range is inclusive and ends only on the next index
	}

	fmt.Printf("Range array: %v \n", diff)

	prefixSums := make([]int, l+2) // 0th + all sums of diffs
	prefixSums[0] = 0
	//prefixSums[l+1] = 666 // test that the last value is also calculated

	for i, v := range diff {
		prefixSums[i+1] = prefixSums[i] + v

		// we can break at every position before calculating the complete prefixSums array
		if (i < l) && (nums[i] > prefixSums[i+1]) {
			fmt.Printf("nums[%v] = %v > prefixSums[%v] = %v. Returning false. \n", i, nums[i], i+1, prefixSums[i+1])
			return false
		}
	}

	fmt.Printf("Prefix sum array: %v \n", prefixSums)

	/*
		for i, n := range nums {
			if n > prefixSums[i+1] {
				fmt.Printf("nums[%v] = %v > prefixSums[%v] = %v. Returning false. \n", i, n, i+1, prefixSums[i+1])
				return false
			}
		}
	*/

	return true
}

func test(nums []int, queries [][]int, expectedResult bool) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", nums)
	fmt.Printf("Queries: %v \n", queries)

	result := isZeroArray(nums, queries)

	fmt.Printf("Possible to make 0-array with given queries: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	nums := []int{4, 3, 2, 1}

	queries := [][]int{
		{1, 3},
		{0, 2},
	}

	expected := false

	test(nums, queries, expected)
}

func test2() {
	nums := []int{3, 2, 1, 2}

	queries := [][]int{
		{0, 2},
		{1, 3},
		{2, 3},
	}

	expected := false

	test(nums, queries, expected)
}

func test3() {
	nums := []int{1, 0, 1}

	queries := [][]int{
		{0, 2},
	}

	expected := true

	test(nums, queries, expected)
}

func main() {
	// 3355. Zero Array Transformation I
	test1()
	test2()
	test3()
}
