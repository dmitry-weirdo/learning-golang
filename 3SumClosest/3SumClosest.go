package main

import (
	"fmt"
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	// external is similar to normal 3 sum
	slices.Sort(nums)

	closestSum := nums[0] + nums[1] + nums[2] // we're guaranteed that there are at least 3 elements
	if closestSum == target {
		return closestSum
	}

	for i := 0; i < len(nums)-2; i++ { // 2 remaining values are required to search values 2 and 3
		// we do NOT skip duplicates in this ticket, since we search for index pairs, NOT value pairs

		//if (i > 0) && (nums[i] == nums[i-1]) {
		//	// skip duplicate a[i] values
		//	continue
		//}

		// search from the next element
		closestSumForI := twoSumSmaller(nums, nums[i], i+1, target)

		if abs(closestSum-target) > abs(closestSumForI-target) {
			closestSum = closestSumForI

			if closestSum == target { // optimal result found -> no need to iterate anymore
				return closestSum
			}
		}
	}

	return closestSum
}

func twoSumSmaller(arr []int, currentElement int, firstElementIndex int, target int) int { // returns closest sum
	left := firstElementIndex
	right := len(arr) - 1

	// we're counting array index pairs that have sum < target
	closestSum := currentElement + arr[left] + arr[right]
	currentSum := 0

	for left < right {
		currentSum = currentElement + arr[left] + arr[right]

		if abs(closestSum-target) > abs(currentSum-target) { // we can improve dist to target from both sides -> check for every value
			closestSum = currentSum

			if closestSum == target { // optimal result found -> no need to iterate anymore
				return closestSum
			}
		}

		if currentSum >= target {
			// we cannot increase left since it will increase the sum
			// -> decrease right

			right--
		} else { // sum <= target
			// we cannot decrease right since it will decrease the sum
			// -> increase left

			left++
		}
	}

	return closestSum
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func test(arr []int, target int, expectedResult int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Target sum: %v \n", target) // will always be 0

	result := threeSumClosest(arr, target)

	fmt.Printf("Sum of a[i] + a[j] + a[k] closest to %v: %v \n", target, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{-1, 2, 1, -4}
	target := 1
	expected := 2 // -1 + 2 + 1 = 2

	test(arr, target, expected)
}

func test2() {
	arr := []int{0, 0, 0}
	target := 1
	expected := 0 // 0 + 0 + 0 = 0

	test(arr, target, expected)
}

func test3() {
	arr := []int{0, 0, 0}
	target := 10000
	expected := 0 // 0 + 0 + 0 = 0

	test(arr, target, expected)
}

func main() {
	// 16. 3Sum Closest
	// Very similar to "259. 3Sum Smaller".
	test1()
	test2()
	test3()
}
