package main

import "fmt"

func getSum(a int, b int) int {
	return getSum_xor(a, b)

	// this fails on:
	// "You are not allowed to use += operator"
	//return getSum_naive(a, b)
}

func getSum_xor(a int, b int) int {
	// This is a black-magic solution.

	// XOR ^ - performs a sum without carry bits
	// AND & - leaves just the carry bits, and we need to move them to the next digit, i.e. * 2 (same as << 1)

	// then we continue it with the values of (a xor b), (a & b << 1), until the carry is 0

	// todo: understand why this black magic is also working for the negative numbers

	answer := a
	carry := b

	xor := 0
	carryBitsShiftedToNextBit := 0

	for carry != 0 {
		xor = answer ^ carry                              // sum without carries
		carryBitsShiftedToNextBit = (answer & carry) << 1 // only the carry bits carried to the next digit

		answer = xor
		carry = carryBitsShiftedToNextBit
	}

	return answer
}

func getSum_naive(a int, b int) int { // this fails for the negative number
	aBit := 0
	bBit := 0

	overflow := 0

	result := 0

	powerOfTwo := 1

	for a != 0 || b != 0 {
		aBit = a & 1
		bBit = b & 1

		fmt.Printf("a bit: %v, b bit: %v, overflow: %v, result = %v, powerOf2 = %v \n", aBit, bBit, overflow, result, powerOfTwo)

		a = a / 2
		b = b / 2

		overflow += aBit
		overflow += bBit

		result += (overflow & 1) * powerOfTwo
		powerOfTwo *= 2

		overflow /= 2
	}

	// add the remaining overflow if there is one
	if overflow > 0 {
		result += (overflow & 1) * powerOfTwo
	}

	return result
}

func test(x, y int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	// %032b is a 32-bit format,
	// We need to convert signed int to unsigned unit32
	// - it does NOT change the bit representation,
	// - it only changes the treatment of the same bits.
	// If we print %v for negative values, it will display -<positiveBits>,
	// this is not what we want.
	fmt.Printf("Number 1: %v = %032b \n", x, uint32(x))
	fmt.Printf("Number 2: %v = %032b \n", y, uint32(y))

	result := getSum(x, y)

	fmt.Printf("Sum x + y with just bitwise operations: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(1, 2, 3)
}

func test2() {
	test(2, 3, 5)
}

func test3() {
	test(1000, -1000, 0)
}

func test4() {
	test(1000, 1000, 2000)
}

func test5() {
	test(-1000, -1000, -2000)
}

func test6() {
	// -5
	// 11111111111111111111111111111011

	// 5
	// 00000000000000000000000000000101

	test(-5, 5, 0)
}

func main() {
	// 371. Sum of Two Integers
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
