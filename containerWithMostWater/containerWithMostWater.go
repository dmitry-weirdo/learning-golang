package main

import (
	"fmt"
)

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	result := min(height[left], height[right]) * (right - left)

	for left < right {
		// move from the lowest height within the range
		// Why: since (right - left) will decrease from the current state,
		// we can only grow the result by increasing the min(heightLeft; heightRight).
		// If we move from the smallest, we give the min a chance to increase.
		// If we'd move from the highest, the min will stay the same, no matter the change.

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

		currentArea := min(height[left], height[right]) * (right - left)
		fmt.Printf("Left: %v, right: %v, currentArea: %v \n", left, right, currentArea)

		if currentArea > result {
			fmt.Printf("Updated max height to %v \n", currentArea)

			result = currentArea
		}
	}

	return result
}

func main() {
	// 11. Container With Most Water
	var height []int

	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	area := maxArea(height)

	fmt.Printf("Height: %v, max area: %v \n", height, area)
}
