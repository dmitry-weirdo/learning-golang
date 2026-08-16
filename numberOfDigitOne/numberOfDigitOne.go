package main

import "fmt"

func countDigitOne(n int) int {
	// we're counting of every digit appearances
	count := 0

	i := 1

	for i <= n {
		next10Power := i * 10

		// 1-st digit makes full 1-time iteration - (n / 10) complete times
		// n = 211 -> complete iterations = 21 * 1 = 21.
		// i.e. in every of 21 of tens iteration, we count last 1-digit once

		// 2-nd digit - (n / 100) complete times.
		// n = 211 -> complete iterations = 2 * 10 = 20.
		// i.e. on every 2 of hundreds iteration, we count 2nd 1-digit 10 times ( x10-x19 values )

		// etc.
		completeIterations := i * (n / next10Power)

		// (i - 1) is 0, 9, 99, 999 etc. - subtracting the values before 1x... in the current digit

		// 1st digit -> n % 10 is the remaining digits of full tenths
		// n = 211 -> n % 10 = 1
		// i.e. if the last digit is > 0, it will be 1, 2, ... 9 (although maximum we can have just 1)

		// 2nd digit -> n % 100 is the remaining of full hundreds.
		// n = 211 -> n % 100 = 11
		// we subtract 9 and get 2 - for x10 and x11.
		// for 19 we subtract 9 and will get 10 - for all values x10-x19
		// for > 19 we'll get > 10 -> then we should decrease to the most possible 10.

		// etc.
		partlyIterations := n%next10Power - (i - 1)

		if partlyIterations < 0 {
			partlyIterations = 0
		} else if partlyIterations > i {
			// we can only iterate 1, 10, 100, 1000... times for every next digit
			// for 1st digit - max 1 in x0 - x9 (just x1)
			// for 2nd digit - max 10 in x10-x19 values
			// for 3rd digit - max 100 in x100-x199 values
			// etc.
			partlyIterations = i
		}

		count += completeIterations
		count += partlyIterations

		i *= 10 // jump to the next digit
	}

	return count
}

func test(x int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x)

	result := countDigitOne(x)

	fmt.Printf("Count of decimal 1 digits in all numbers [1; %v]: %v \n", x, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(13, 6) // last digit in 1 and 11, 2nd digit in 10, 11, 12, 13
}

func test2() {
	test(0, 0)
}

func test3() {
	test(20, 12) // last digit in 1 and 11 (2 times), 2nd digit in 10-19 (10 times)
}

func test4() {
	test(21, 13) // last digit in 1, 11, 21 (3 times), 2nd digit in 10-19 (10 times)
}

func test5() {
	// 1st digit 21 * 1 + 1 = 22 times
	// 2nd digit in 20 + 2 (for 210, 211) = 22 times
	// 3rd digits 100-199 = 100 times
	test(211, 144)
}

func main() {
	// 233. Number of Digit One
	test1()
	test2()
	test3()
	test4()
	test5()
}
