package main

import (
	"fmt"
)

func mySqrt(x int) int {
	return mySqrt_overflowSafe(x)
	//return mySqrt_naive(x)
}

func mySqrt_overflowSafe(x int) int {
	if x < 2 {
		return x
	}

	left := 0  // 0 is a possible root of 0
	right := x // square is <= than x / 2 - wrong for 1

	for left < right {
		mid := left + (right-left)/2 // overflow-safe

		if (x/mid == mid) && (x%mid == 0) { // not calculating mid * mid
			return mid
		}

		if (mid + 1) > x/(mid+1) { // we search the first value where adding 1 to it will overflow x
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func mySqrt_naive(x int) int {
	left := 0  // 0 is a possible root of 0
	right := x // square is <= than x / 2 - wrong for 1

	for left < right {
		// todo: if overflow, can use left + (right - left) / 2
		mid := (left + right) / 2

		// todo: potential overflow
		if mid*mid == x { // exact square match
			return mid
		}

		// todo: potential overflow
		if (mid+1)*(mid+1) > x { // we search the first value where adding 1 to it will overflow x
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func test(x int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("X: %v \n", x)

	result := mySqrt_overflowSafe(x)

	fmt.Printf("Rounded-down sqrt of %v: %v \n", x, result)
	fmt.Printf("Expected rounded-down sqrt of %v: %v \n", x, expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	x := 8
	expected := 2

	test(x, expected)
}

func test2() {
	x := 4
	expected := 2

	test(x, expected)
}

func test3() {
	x := 1
	expected := 1

	test(x, expected)
}

func test4() {
	x := 0
	expected := 0

	test(x, expected)
}

func main() {
	// 69. Sqrt(x)
	test1()
	test2()
	test3()
	test4()
}
