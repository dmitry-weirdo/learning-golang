package main

import "fmt"

func grayCode(n int) []int {
	size := 1 << n // 2^n

	result := make([]int, size)

	for i := range size { // iterate all values from 0 to 2^n
		result[i] = i ^ (i >> 1) // generate next value
	}

	return result
}

func test(n int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Number of bits: %v \n", n)

	result := grayCode(n)

	fmt.Printf("Gray code values for this number of bits: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	test(
		2,
		[]int{0, 1, 3, 2},
	)
}

func test2() {
	test(
		1,
		[]int{0, 1},
	)
}

func main() {
	// 89. Gray Code
	test1()
	test2()
}
