package main

import (
	"fmt"
	"slices"
)

func moveZeroes(nums []int) {
	writePos := 0

	for i, v := range nums {
		if v != 0 {
			// we can always swap and avoid this check
			// todo: this is still unclear why it will move zeroes to the end, but ok
			nums[writePos], nums[i] = nums[i], nums[writePos]

			writePos++
		}
	}
}

func moveZeroesNaive(nums []int) {
	writePos := 0

	for i, v := range nums {
		if v != 0 {
			// todo: we can avoid this check and
			if i != writePos { // do not overwrite itself
				nums[writePos] = v
				nums[i] = 0
			}

			writePos++
		}
	}
}

func test(nums []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", nums)

	moveZeroes(nums)

	fmt.Printf("Array after moving zeroes: %v \n", nums)

	if !slices.Equal(nums, expectedResult) {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, nums)
	}
}

func test1() {
	nums := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}

	test(nums, expected)
}

func test2() {
	nums := []int{0}
	expected := []int{0}

	test(nums, expected)
}

func test3() {
	nums := []int{1, 2, 3}
	expected := []int{1, 2, 3}

	test(nums, expected)
}

func main() {
	// 283. Move Zeroes
	test1()
	test2()
	test3()
}
