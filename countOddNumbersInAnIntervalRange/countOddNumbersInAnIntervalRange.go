package main

import "fmt"

func countOdds(low int, high int) int {
	// Even, even: [8, 10], odd numbers 9. (10 - 8) / 2 = 1, this is the answer.
	// Odd, odd: [3, 5], odd numbers 3, 5. (5 - 3) / 2 = 1, we need to add 1.
	// Odd, even: [1, 4], odd numbers 1, 3. (4 - 1) / 2 = 1, we need to add 1.
	// Even, odd [2, 5], odd numbers 3, 5. (5 - 2) / 2 = 1, we need to add 1.

	// I.e. we need to add 1 to (high - low) / 2 always except the case when both low and high are even.

	if (low%2 == 0) && (high%2 == 0) { // both even
		return (high - low) / 2
	}

	return (high-low)/2 + 1
}

func test(l, r int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Left: %v \n", l)
	fmt.Printf("Right: %v \n", r)

	result := countOdds(l, r)

	fmt.Printf("Total odd numbers in range [%v; %v] : %v \n", l, r, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(3, 7, 3) // 3, 5, 7
}

func test2() {
	test(8, 10, 1) // 9
}

func test3() {
	test(3, 5, 2) // 3, 5
}

func test4() {
	test(1, 4, 2) // 1, 3
}

func test5() {
	test(2, 5, 2) // 3, 5
}

func test6() {
	test(1, 1, 1) // 1
}

func test7() {
	test(2, 2, 0) // no odd numbers
}

func main() {
	// 1523. Count Odd Numbers in an Interval Range
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
}
