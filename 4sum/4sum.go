package main

import "slices"

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)

	// result are the values, not indexes
	result := make([][]int, 0)

	for i := 0; i < len(nums)-3; i++ { // 3 remaining values at the end are required for search values 2, 3, 4
		if (i > 0) && (nums[i] == nums[i-1]) {
			// skip duplicate a[i] values
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ { // 2 remaining values at the end are required for search values 3, 4
			if (j > i+1) && (nums[j] == nums[j-1]) {
				// skip duplicate a[j] values
				continue
			}

			// search with twoSum O(n) on the remaining array
			twoSum(nums, i, j, target, &result) // will append to result
		}
	}

	return result
}

func twoSum(arr []int, firstElementIndex int, secondElementIndex int, target int, result *[][]int) {
	firstElement := arr[firstElementIndex]
	secondElement := arr[secondElementIndex]

	left := secondElementIndex + 1
	right := len(arr) - 1

	for left < right {
		sum := firstElement + secondElement + arr[left] + arr[right]

		if sum == target { // found the result!
			*result = append(*result, []int{firstElement, secondElement, arr[left], arr[right]})

			// go to next values on both sides
			left++
			right--

			// skip duplicates from left, we compare to the successful left and move until the value is changed
			for left < right && arr[left] == arr[left-1] {
				left++
			}

			// skip duplicates from right, we compare to the successful right and move until the value is changed
			for left < right && arr[right] == arr[right+1] {
				right--
			}
		} else if sum < target {
			// we need to increase -> increase the smaller
			left++
		} else {
			// we need to decrease -> decrease the bigger
			right--
		}
	}
}

func main() {
	// 18. 4Sum
}
