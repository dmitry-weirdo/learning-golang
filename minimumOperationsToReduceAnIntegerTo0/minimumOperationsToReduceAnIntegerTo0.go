package main

import "fmt"

func minOperations(n int) int {
	// If there is no consecutive 1 bits, we just subtract this power of 2.
	// If there are consecutive 1 bits, we add to this power of 2 and then subtract it

	// 54 = 110110
	// 0110 is consecutive 1-bits -> we add 1 to the least significant bits -> we add 2 = 0010, and it becomes 1000
	// 111000 (56) -> consecutive 1-bits -> we add 1 as 1000 (8) -> it becomes 1000000
	// 1000000 (64) - no consecutive bits -> we just subtract 64
	return minOperations_bitOperations(n)
}

func minOperations_bitOperations(n int) int {
	totalOperations := 0

	for n > 0 {
		if n&3 == 3 { // two last bits are 11 (consecutive ones)
			n++               // add 1 to the least significant bit
			totalOperations++ // we used one "add 1 to last significant 1-bit" operation

			// Note that it can set more than just two bits 011 to 100
			// If it is 0111, adding 1 will set last bits to 1000
		} else { // last 1 or 0 is non-consecutive -> 1 operation to remove this 1, 0 operations if it's 0
			if n%2 == 1 { // we only need to subtract the power of 2 if last bit is 1
				totalOperations++
			}

			// remove last bit
			n >>= 1
		}
	}

	return totalOperations
}

func test(x int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x) // todo: replace with your text if required

	result := minOperations(x)

	fmt.Printf("Min 2-power operations to reduce %v to 0: %v \n", x, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(54, 3)
}

func test2() {
	test(39, 3)
}

func test3() {
	test(64, 1)
}

func main() {
	// 2571. Minimum Operations to Reduce an Integer to 0
	test1()
	test2()
	test3()
}
