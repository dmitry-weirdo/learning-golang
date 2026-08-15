package main

import (
	"fmt"
	"math"
)

func hammingWeight(n int) int {
	// optimized solution - iterates only on 1-value bits.
	return hammingWeight_optimized(n)

	// trivial solution O(log n), dividing by 2 and checking the last bit
	//return hammingWeight_trivial(n)
}

func hammingWeight_optimized(n int) int {
	// Magic observation: n & (n - 1) flips just the last 1 bit, but leaves all other bits in place.
	// Therefore, here the count of operations will be limited to the number of 1 bits onlz.

	// Example: n = 4 = 100
	// n - 1 = 3 = 011
	// n & n - 1 = 100 & 011 = 000
	// It flipped the 3rd bit, and other bits remain 0
	// Since result is 0, we can stop the iteration.

	sum := 0

	for n != 0 {
		sum++ // we have at least one significant bit

		// flip the last significant bit to 0
		n &= n - 1
	}

	return sum
}

func hammingWeight_trivial(n int) int {
	sum := 0

	for n > 0 {
		if n&1 == 1 {
			sum++
		}

		//n = n / 2
		n >>= 1
	}

	return sum
}

func test(x int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x)

	result := hammingWeight(x)

	fmt.Printf("Count of 1 bits in %v: %v \n", x, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(4, 1) // 100
}

func test2() {
	test(11, 3) // 1011
}

func test3() {
	test(128, 1) // 10000000
}

func test4() {
	test(2147483645, 30) // 1111111111111111111111111111101
}

func test5() {
	// 2147483647
	test(math.MaxInt32, 31) // 2^31 - 1 -> all ones in 31 bits
}

func main() {
	// 191. Number of 1 Bits
	test1()
	test2()
	test3()
	test4()
	test5()
}
