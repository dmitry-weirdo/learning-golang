package main

import (
	"fmt"
	"slices"
)

func canMakeArithmeticProgression(arr []int) bool {
	slices.Sort(arr)

	// we're guaranteed to have 2 elements in the array
	diff := arr[1] - arr[0]

	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != diff {
			return false
		}
	}

	return true
}

func test(arr []int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := canMakeArithmeticProgression(arr) // todo: update to your function

	fmt.Printf("Array can be reordered to an arithmetic progression: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{3, 5, 1}, true)
}

func test2() {
	test([]int{1, 2, 4}, false)
}

func test3() {
	test([]int{1, 2}, true)
}

func main() {
	// 1502. Can Make Arithmetic Progression From Sequence
	test1()
	test2()
	test3()
}
