package main

import (
	"fmt"
)

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	result := min(height[left], height[right]) * (right - left)

	for left < right {
		// move from the lowest height within
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
	var height []int

	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	area := maxArea(height)

	fmt.Printf("Height: %v, max area: %v \n", height, area)
}
