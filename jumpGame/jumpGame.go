package main

import "fmt"

func canJump(nums []int) bool {
	return canJump_greedy_rightToLeft(nums)
}

func canJump_greedy_rightToLeft(nums []int) bool {
	n := len(nums)

	// last position is good -> if we land there, we win
	leftMostGoodPosition := n - 1

	// for all the indexes right-to-left check, whether they are good
	// if we found an index that is good, this is the new leftMostGoodPosition.
	// The target question is - is 0 a good position or not?
	for i := n - 1; i >= 0; i-- {
		if i+nums[i] >= leftMostGoodPosition {
			// this position is good -> this is our leftmost good position,
			// since from it we can jump through good positions to the last element
			leftMostGoodPosition = i
		}
	}

	// check whether on 0-th position, we could reach the leftMostGoodPosition from it
	return leftMostGoodPosition == 0
}

func canJump_greedy_naive(nums []int) bool {
	if len(nums) == 1 { // we're already at the last element
		return true
	}

	if nums[0] == 0 {
		return false
	}

	n := len(nums)

	i := 0

	// todo: this incorrect, since when jumping max, we can jump over the good position (like leading to the end of the array)
	// try to do greedy jumping, decrease if we jump to 0
	for i < n-1 {
		maxJump := i + nums[i]

		for maxJump >= n || // jumped over the array -> success
			((nums[maxJump] == 0) && (maxJump != n-1)) { // if we're at 0 element, but it's the last element, it's success
			maxJump--
		}

		if i >= maxJump { // can't go forward -> fail
			return false
		}

		i = maxJump
	}

	return i == n-1
}

func test(arr []int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := canJump(arr)

	fmt.Printf("Can jump from 0-th to last element: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{2, 3, 1, 1, 4}, true)
}

func test2() {
	test([]int{3, 2, 1, 0, 4}, false)
}

func test3() {
	test([]int{0, 1}, false)
}

func test4() {
	test([]int{0}, true)
}

func test5() {
	// we end at 0, but this is the end of the array
	test([]int{2, 0}, true)
}

func test6() {
	// tricky case -> jumping max = 3 from a[0] will lead to failure
	// However, jumping 2 from a[0] will lead to success
	test([]int{3, 0, 8, 2, 0, 0, 1}, true)
}

func main() {
	// 55. Jump Game
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
