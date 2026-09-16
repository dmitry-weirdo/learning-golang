package main

import (
	"fmt"
	"math/big"
)

func numberOfSets(n int, k int) int {
	return numberOfSets_combinatoric(n, k)
}

func numberOfSets_combinatoric(n int, k int) int {
	// we insert a fake gap points for (k - 1) intervals (the last segment doesn't need one)
	// So we're selecting (2 * k) points from a (n + k - 1) subset.

	// This is
	// (n + k - 1)! / (2 * k)! * (n + k - 1 - 2 * k)!
	// (n + k - 1)! / (2 * k)! * (n - k - 1)!
	const mod = 1_000_000_007
	modBigInt := big.NewInt(mod)

	var c big.Int
	c.Binomial(int64(n-1+k), int64(2*k))
	//fmt.Printf("c (big int): %v \n", c)

	c.Mod(&c, modBigInt) // !!! mod must be also executed on BigInt, NOT after converting to int64
	//fmt.Printf("c mod %v (big int): %v \n", mod, c)

	return int(c.Int64())
}

func test(n, k int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N: %v \n", n)
	fmt.Printf("K: %v \n", k)

	result := numberOfSets(n, k)

	fmt.Printf("Count of 2+ points %v segments in [0; %v]: %v \n", k, n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(4, 2, 5)
}

func test2() {
	test(3, 1, 3)
}

func test3() {
	test(30, 7, 796297179)
}

func test4() {
	// failing test-case 49 / 68
	test(52, 23, 963678472)
}

func main() {
	// 1621. Number of Sets of K Non-Overlapping Line Segments
	test1()
	test2()
	test3()
	test4()
}
