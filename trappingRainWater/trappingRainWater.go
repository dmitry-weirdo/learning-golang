package main

import (
	"fmt"
)

func trap(height []int) int {
	// optimized to use just 2 pointers, not 2 pre-calculated arrays

	left := 0
	right := len(height) - 1

	leftMax := height[0]
	rightMax := height[len(height)-1]

	sum := 0

	// move left and right pointers until they reach the same point that should be already calculated in the previous step
	for left < right {
		if leftMax < rightMax {
			// leftMax will be the defining minimum
			left++

			// update leftMax with the new position increased from left
			leftMax = max(leftMax, height[left])

			// if leftMax was set to height[left], the amount on this position will be 0, even if height[left] > rightMax
			// else leftMax is defining the min of (leftMax, rightMax)
			sum += leftMax - height[left]
		} else {
			// rightMax will be the defining minimum
			right--

			// update rightMax with the new position decreased from right
			rightMax = max(rightMax, height[right])

			// if rightMax was set to height[right], the amount on this position will be 0, even if height[right] > leftMax
			// else rightMax is defining the min of (leftMax, rightMax)
			sum += rightMax - height[right]
		}
	}

	return sum
}

func trapWithPrefixArrays(height []int) int {
	n := len(height)

	if n < 3 { // no in-between space for water
		return 0
	}

	// 2 arrays will take 2 * O(n) space.
	// calculate max left heights (including self)
	left := make([]int, n)

	leftMax := height[0]
	for i := 0; i < n; i++ {
		leftMax = max(leftMax, height[i])
		left[i] = leftMax
	}

	// calculate max right heights (including self)
	right := make([]int, n)

	rightMax := height[n-1]
	for i := n - 1; i >= 0; i-- {
		rightMax = max(rightMax, height[i])
		right[i] = rightMax
	}

	// calculate result
	sum := 0

	for i := 0; i < n; i++ {
		// we subtract the current height since we're only putting water at the top of the current height
		// todo: what if this is negative?
		sum += min(left[i], right[i]) - height[i]
	}

	return sum
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := trap(arr)

	fmt.Printf("Max water trapped: %v \n", result)
	fmt.Printf("Expected max water trapped: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	expected := 6

	test(arr, expected)
}

func test2() {
	arr := []int{4, 2, 0, 3, 2, 5}
	expected := 9

	test(arr, expected)
}

func test3() {
	arr := []int{3, 0, 2, 0, 4}
	expected := 7

	test(arr, expected)
}

func main() {
	// 42. Trapping Rain Water
	test1()
	test2()
	test3()
}
