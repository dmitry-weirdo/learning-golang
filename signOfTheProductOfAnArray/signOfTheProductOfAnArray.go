package main

import "fmt"

func arraySign(nums []int) int {
	s := 1

	for _, v := range nums {
		if v == 0 {
			return 0
		}

		s *= sign(v)
	}

	return s
}

func sign(v int) int {
	switch {
	case v == 0:
		return 0
	case v < 0:
		return -1
	default:
		return 1
	}
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := arraySign(arr)

	fmt.Printf("Sign of the product of the array: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{-1, -2, -3, -4, 3, 2, 1}, 1)
}

func test2() {
	test([]int{1, 5, 0, 2, -3}, 0)
}

func test3() {
	test([]int{-1, 1, -1, 1, -1}, -1)
}

func main() {
	// 1822. Sign of the Product of an Array
	test1()
	test2()
	test3()
}
