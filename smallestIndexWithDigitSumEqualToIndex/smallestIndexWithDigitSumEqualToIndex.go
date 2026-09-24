package main

import "fmt"

func smallestIndex(nums []int) int {
	for i, v := range nums {
		if getSumOfDigits(v) == i {
			return i
		}
	}

	return -1
}

func getSumOfDigits(n int) int {
	sum := 0

	for n != 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := smallestIndex(arr)

	fmt.Printf("Smallest index where sum of digits is equal to index: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{1, 3, 2}, 2)
}

func test2() {
	test([]int{1, 10, 11}, 1) // 10 = 1, 11 = 2 both match, index 1 is smaller
}

func test3() {
	test([]int{1, 2, 3}, -1) // no target index
}

func test4() {
	// test-case
	test([]int{101, 135, 2}, 2)
}

func main() {
	// 3550. Smallest Index With Digit Sum Equal to Index
	test1()
	test2()
	test3()
	test4()
}
