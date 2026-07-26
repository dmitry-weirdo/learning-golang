package main

import "fmt"

func removeDuplicates(nums []int) int {
	return removeDuplicatesOptimized(nums)
	//return removeDuplicatesNaive(nums)
}

func removeDuplicatesOptimized(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	// we know that the array is sorted, so we can check the values backwards
	writePos := 2

	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[writePos-2] {
			// we've already written 2 of this value -> skip it
			// note that writePos is the position where we WILL write, i.e. we're checking 2 values before

			// 1, 1, 1, 3
			// i = 3, writePos = 2, so we're checking writePos[0] (and inherently writePos[1])
			// then we skip for i = 3, and for i = 4 we will write to writePos[2] -> 1, 1, 3, 3

			continue
		}

		nums[writePos] = nums[i]
		writePos++
	}

	return writePos
}

func removeDuplicatesNaive(nums []int) int {
	writePos := 1

	current := nums[0]
	currentStart := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] == current { // value stays the same
			if (i - currentStart) < 2 { // write the current value 2 times
				nums[writePos] = current
				writePos++
			}
		} else { // value changed -> write it to the result
			current = nums[i]
			currentStart = i

			nums[writePos] = current
			writePos++
		}
	}

	return writePos
}

func main() {
	// 80. Remove Duplicates from Sorted Array II
	arr := []int{1}
	fmt.Printf("Array before: %v \n", arr)

	removeDuplicatesOptimized(arr)

	fmt.Printf("Array after: %v \n", arr)
}
