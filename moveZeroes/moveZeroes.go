package main

import (
	"fmt"
	"slices"
)

func moveZeroes(nums []int) {
	writePos := 0

	// The magic is:
	// - For the starting non-0 values, it will swap with self
	// - When 0 is encountered, writePos stays at this 0 and i skips up to next non-0 value.
	// - Then we swap the values, and the value at [i] is now 0.
	// - If there are further 0-values skipped, writePos is again at 0.
	// - If this was a single 0, we've just swapped the next non-0 value with previous 0, and this is again 0.

	// [1, 2, 3, 0, 10]
	//           W
	//           I

	// Pointer I skips the 0 value:
	// [1, 2, 3, 0, 10]
	//           W
	//              I

	// Swap is made - since it's the next element after the previous 0 position, W is again at 0:
	// [1, 2, 3, 10, 0]
	//               W
	//                  I

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
