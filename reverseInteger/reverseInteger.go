package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x == 0 { // corner-case -> do not remove the tail zeroes infinitely
		return 0
	}

	if x == math.MinInt32 { // corner-case -> negating it in int32 will cause overflow
		return 0
	}

	//fmt.Printf("Min int negated: %v \n", -math.MinInt32)
	//fmt.Printf("Min int: %v \n", math.MinInt32)
	//fmt.Printf("Max int: %v \n", math.MaxInt32)

	negative := x < 0

	x = abs(x)

	// there can be multiple leading zeroes
	for x%10 == 0 { // remove last zeroes since they will be leading zeroes
		x = x / 10
	}

	n := 0

	const overflowLimit = math.MaxInt32 / 10 // without the last digit

	for x > 0 {
		nextDigit := x % 10

		// remove last digit from the original number
		x = x / 10

		if n > overflowLimit { // any next digit will overflow
			return 0
		}

		// this handling is for pureness of code -> actually the numbers like 2147483648 reversed will not feat within int32, so it could not be an input
		// within 7, it will fit in both positive and negative
		// if it's 8, it will only fit the negative range
		// if it's > 8, it is an overflow
		if (n == overflowLimit) && (nextDigit > 7) {
			// Min int32: -2147483648
			// Reversed input: 8463847412
			if negative && (nextDigit == 8) { // we can handle -2147483648 negative, but not positive
				if x == 0 { // it is the last digit -> we have exactly MinInt32
					return math.MinInt32
				} else { // there are more digits -> it's an overflow
					return 0
				}
			} else { // overflow
				return 0
			}

			// Max int32: 2147483647
			// Reversed input: 7463847412
		}

		// add next digit
		n = 10*n + nextDigit
	}

	if negative {
		return -n
	}

	return n
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func test(x int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x)

	result := reverse(x)

	fmt.Printf("Reversed number: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(123, 321) // trivial case
}

func test2() {
	test(-123, -321) // negative value
}

func test3() {
	test(120, 21) // remove leading zeroes
}

func test4() {
	test(0, 0) // corner case -> don't remove leading zeroes, just return 0
}

func test5() {
	// since 9646324351 > MaxInt -> return 0
	test(1534236469, 0)
}

func test6() {
	// MinInt reversed will be > MaxInt -> return 0
	test(math.MinInt32, 0)
}

func test7() {
	// 7463847412 // MaxInt reversed
	// this won't fit into int32, so this is a theoretical test
	test(7463847412, math.MaxInt32)
}

func test8() {
	// 8463847412 // MaxInt + 1 reversed -> overflow
	// this won't fit into int32, so this is a theoretical test
	test(8463847412, 0)
}

func test9() {
	// -7463847412 // MinInt + 1 reversed
	// this won't fit into int32, so this is a theoretical test
	test(-7463847412, math.MinInt32+1)
}

func test10() {
	// -8463847412 // MinInt reversed
	// this won't fit into int32, so this is a theoretical test
	test(-8463847412, math.MinInt32)
}

func test11() {
	// -9463847412 // MinInt - 1 reversed -> overflow
	// this won't fit into int32, so this is a theoretical test
	test(-9463847412, 0)
}

func main() {
	// 7. Reverse Integer
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
	test9()
	test10()
	test11()
}
