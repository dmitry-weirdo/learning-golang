package main

import "fmt"

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	writeIndex := len(nums) - 1 // we're writing from the maxes

	result := make([]int, len(nums))

	for left <= right {
		leftValue := abs(nums[left])
		rightValue := abs(nums[right])

		if leftValue >= rightValue {
			result[writeIndex] = leftValue * leftValue
			left++
		} else {
			result[writeIndex] = rightValue * rightValue
			right--
		}

		writeIndex--
	}

	return result
}

func abs(x int) int {
	// todo: not handling MinInt!
	if x > 0 {
		return x
	}

	return -x
}

func test(arr []int) {
	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("Array: %v \n", arr)

	result := sortedSquares(arr)

	fmt.Printf("Squares array: %v  \n", result)
}

func test1() {
	arr := []int{-4, -1, 0, 3, 10}

	test(arr)
}

func test2() {
	arr := []int{}

	test(arr)
}

func test3() {
	arr := []int{-10}

	test(arr)
}

func main() {
	test1()
	test2()
	test3()
}
