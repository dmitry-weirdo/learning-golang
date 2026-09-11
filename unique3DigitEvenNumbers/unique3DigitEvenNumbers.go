package main

import (
	"fmt"
)

func totalNumbers(digits []int) int {
	m := make([]int, 10) // count of digits

	for _, v := range digits {
		m[v]++
	}

	if (m[0] == 0) && (m[2] == 0) && (m[4] == 0) && (m[6] == 0) && (m[8] == 0) { // no even digits -> cannot generate even numbers
		return 0
	}

	count := 0

	d := make([]int, 10)

	for i := 100; i <= 998; i += 2 { // only iterate even numbers
		digit2 := i / 100
		digit1 := (i % 100) / 10
		digit0 := i % 10
		//fmt.Printf("i: %v, digits: %v %v %v \n", i, digit2, digit1, digit0)

		// count the digits of the current number
		d[digit0]++
		d[digit1]++
		d[digit2]++

		// check whether there are enough digits for the current number
		if d[digit0] <= m[digit0] &&
			d[digit1] <= m[digit1] &&
			d[digit2] <= m[digit2] {
			count++
		}

		// backtrack the counts to 0
		d[digit0]--
		d[digit1]--
		d[digit2]--
	}

	return count
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array of digits: %v \n", arr)

	result := totalNumbers(arr)

	fmt.Printf("Count of distinct 3-digit numbers that can be generated from the array: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{1, 2, 3, 4},
		12, // 124, 132, 134, 142, 214, 234, 312, 314, 324, 342, 412, 432.
	)
}

func test2() {
	test(
		[]int{0, 2, 2},
		2, // 202, 220
	)
}

func test3() {
	test(
		[]int{6, 6, 6},
		1, // 666
	)
}

func test4() {
	test(
		[]int{1, 3, 5},
		0, // all is odd -> 0 even numbers could be generated
	)
}

func main() {
	// 3483. Unique 3-Digit Even Numbers
	// It's the same as "2094. Finding 3-Digit Even Numbers"
	test1()
	test2()
	test3()
	test4()
}
