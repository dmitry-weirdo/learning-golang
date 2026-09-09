package main

import "fmt"

func countCommas(n int64) int64 {
	if n < 1_000 {
		return 0
	}

	if n < 1_000_000 {
		return n - 1_000 + 1
	}

	if n < 1_000_000_000 {
		return (999_999 - 1_000 + 1) + 2*(n-1_000_000+1)
	}

	if n < 1_000_000_000_000 {
		return (999_999 - 1_000 + 1) + 2*(999_999_999-1_000_000+1) + 3*(n-1_000_000_000+1)
	}

	if n < 1_000_000_000_000_000 {
		return (999_999 - 1_000 + 1) + 2*(999_999_999-1_000_000+1) + 3*(999_999_999_999-1_000_000_000+1) + 4*(n-1_000_000_000_000+1)
	}

	// this is just one possible value 10^15 that adds 5 commas once.
	return (999_999 - 1_000 + 1) + 2*(999_999_999-1_000_000+1) + 3*(999_999_999_999-1_000_000_000+1) + 4*(999_999_999_999_999-1_000_000_000_000+1) + 5
}

func test(x int64, expectedResult int64) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x)

	result := countCommas(x)

	fmt.Printf("Commas used for numbers [1, %v]: %v \n", x, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(1002, 3)
}

func test2() {
	test(998, 0)
}

func test3() {
	test(1004590, 1008182)
}

func main() {
	// 3871. Count Commas in Range II
	test1()
	test2()
	test3()
}
