package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	// this ticket sums to 0, but solution should be generalized to work for any target
	return threeSumWithTarget(nums, 0)
}

func threeSumWithTarget(nums []int, target int) [][]int {
	slices.Sort(nums)

	// result are the values, not indexes
	result := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ { // 2 remaining values are required to search values 2 and 3
		if (i > 0) && (nums[i] == nums[i-1]) {
			// skip duplicate a[i] values
			continue
		}

		// search with twoSum O(n) on the remaining array
		twoSum(nums, i, target, &result) // will append to result
	}

	return result
}

func twoSum(arr []int, firstElementIndex int, target int, result *[][]int) {
	firstElement := arr[firstElementIndex]

	left := firstElementIndex + 1
	right := len(arr) - 1

	for left < right {
		sum := firstElement + arr[left] + arr[right]

		if sum == target { // found the result!
			*result = append(*result, []int{firstElement, arr[left], arr[right]})

			// skip duplicates from left
			for left < right && arr[left] == arr[left+1] {
				left++
			}

			// skip duplicates from right
			for left < right && arr[right] == arr[right-1] {
				right--
			}
		}

		if sum < target {
			// we need to increase -> increase the smaller
			left++
		} else {
			// we need to decrease -> decrease the bigger
			right--
		}
	}
}

func test(arr []int, target int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Target sum: %v \n", target) // will always be 0

	result := threeSum(arr)

	fmt.Printf("Result triplets summing to %v: \n%v \n", target, result)
	fmt.Printf("Expected result: \n%v \n", expectedResult)

	// todo: compare results

	// return indexes are 1-based!
	/*	value1 := arr[result[0]-1]
		value2 := arr[result[1]-1]

		fmt.Printf("1-based indexes with targetSum = %v: %v \n", target, result)
		fmt.Printf("Values: %v + %v = %v. Target = %v. \n", value1, value2, value1+value2, target)
		fmt.Printf("Expected result:          %v \n", expectedResult)

		if value1+value2 != target {
			fmt.Printf("FAILURE: %v + %v = %v != target = %v \n", value1, value2, value1+value2, target)
		}

		if result[0] != expectedResult[0] {
			fmt.Printf("FAILURE: expected result[0] = %v, actual result[0] = %v \n", expectedResult[0], result[0])
			return
		}

		if result[1] != expectedResult[1] {
			fmt.Printf("FAILURE: expected result[1] = %v, actual result[1] = %v \n", expectedResult[1], result[1])
			return
		}*/
}

func test1() {
	arr := []int{0, 0, 0, 0}
	target := 0 // always 0 for this problem
	expected := [][]int{
		{0, 0, 0}, // no duplicates!
	}

	test(arr, target, expected)
}

func main() {
	// 15. 3Sum
	test1()
}
