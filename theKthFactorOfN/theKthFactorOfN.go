package main

import (
	"fmt"
	"slices"
)

func kthFactor(n int, k int) int {
	factors := allFactors(n)
	//fmt.Printf("Factors of %v: %v \n", n, factors)

	if k > len(factors) { // if K > count of factors -> return -1
		return -1
	}

	// K is 1-based!
	return factors[k-1]
}

func allFactors(n int) []int { // not just prime factors, but all, including 1 and number itself
	var small, large []int

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			small = append(small, i)

			if i != n/i {
				large = append(large, n/i)
			}
		}
	}

	// large must be reversed
	slices.Reverse(large)

	small = append(small, large...)

	return small
}

func test(n, k int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number N: %v \n", n)
	fmt.Printf("K - number of the prime divisor: %v \n", k)

	result := kthFactor(n, k)

	fmt.Printf("%v-th prime factor of %v: %v \n", k, n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(12, 3, 3)
}

func test2() {
	test(7, 2, 7)
}

func test3() {
	test(4, 4, -1)
}

func main() {
	// 1492. The kth Factor of n
	test1()
	test2()
	test3()
}
