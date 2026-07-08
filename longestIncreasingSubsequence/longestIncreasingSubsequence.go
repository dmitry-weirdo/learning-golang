package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	// heuristic solution of O(n * log n)

	// we can just keep a sorted list of last elements of a sequence, so that we keep the lowest possible value on every position
	// if a new element is bigger than all previous, we append it to the sequence
	// if new element is smaller than some element[i], we replace this element with the new value.
	// At the end, the list will be the longest increasing sequence with the smallest values.

	l := make([]int, 0)

	for i, v := range nums {
		fmt.Printf("i = %v, list = %v \n", i, l)

		pos := binarySearchFirstBiggerValue(l, v)

		fmt.Printf("Binary search of %v to value %v returned %v. \n", l, v, pos)

		if pos >= len(l) {
			l = append(l, v)
		} else {
			l[pos] = v
		}
	}

	fmt.Printf("Longest increasing subsequence: %v \n", l)

	return len(l)
}

func binarySearchFirstBiggerValue(arr []int, value int) int { // returns index
	left := 0
	right := len(arr) // for "result can be after last element", we initialize right wit len(arr), and NOT with len(arr) - 1

	for left < right {
		mid := (left + right) / 2

		if arr[mid] < value {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func lengthOfLIS_dp(nums []int) int {
	// solution with DP, time O(n^2) - on every index, we check all the previous indexes
	n := len(nums)
	memo := make([]int, n)

	for i := range memo { // "longest sequence" for every index is set to 1 at the beginning
		memo[i] = 1
	}

	maxSequenceLength := 1

	for i := 1; i < n; i++ { // skip the first element since its sequence is always 1
		for j := 0; j < i; j++ { // iterate earlier indexes
			fmt.Printf("i = %v, j = %v \n", i, j)

			if nums[j] < nums[i] { // we can extend the longest sequence with nums[i]
				newSequenceLength := memo[j] + 1 // longest sequence to j + nums[i]
				if memo[i] < newSequenceLength {
					fmt.Printf("nums[%v] = %v < nums[%v] = %v. Updated memo[%v] to %v. \n", j, nums[j], i, nums[i], i, newSequenceLength)

					memo[i] = newSequenceLength
				}
			}
		}

		// we can check after the complete position [i] has been updated
		if memo[i] > maxSequenceLength {
			fmt.Printf("Increased max increasing subsequence length from %d to %d. \n", maxSequenceLength, memo[i])
			maxSequenceLength = memo[i]
		}
	}

	fmt.Printf("Memo: %v \n", memo)

	return maxSequenceLength
}

func test(arr []int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("Array: %v \n", arr)

	maxSequenceLength := lengthOfLIS(arr)

	fmt.Printf("Max increasing subsequence length: %v \n", maxSequenceLength)

	if maxSequenceLength != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, maxSequenceLength)
	}
}

func test1() {
	a := []int{4, 2, 5, 3, 7}
	expected := 3

	test(a, expected)
}

func test2() {
	a := []int{10, 9, 2, 5, 3, 7, 101, 18}
	expected := 4

	test(a, expected)
}

func test3() {
	a := []int{0, 1, 0, 3, 2, 3}
	expected := 4

	test(a, expected)
}

func main() {
	test1()
	test2()
	test3()
}
