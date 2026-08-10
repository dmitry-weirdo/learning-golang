package main

import "fmt"

func smallestRepunitDivByK(k int) int {
	return smallestRepunitDivByK_iterateKRemainders(k) // no map, just iteration is faster - 0 ms
	//return smallestRepunitDivByK_mapOrRemainders(k) // slow, around 15 ms
}

func smallestRepunitDivByK_iterateKRemainders(k int) int {
	if (k%2 == 0) || (k%5 == 0) {
		// last digits of 111... cannot be even or 5
		return -1
	}

	remainder := 0

	// we can only iterate K+1 times until we hit a different remainder
	// i.e. we can iterate just K times, next iteration will be a guaranteed (Dirichlet's principle aka Pigeonhole principle)
	for i := 1; i <= k; i++ {
		remainder = (10*remainder + 1) % k

		if remainder == 0 { // result found -> return count of 1s
			return i
		}
	}

	return -1
}

func smallestRepunitDivByK_mapOrRemainders(k int) int {
	if (k%2 == 0) || (k%5 == 0) {
		// last digits of 111... cannot be even or 5
		return -1
	}

	// map to end early if the remainder reappeared
	m := make(map[int]bool)

	remainder := 0

	// we can only iterate K+1 times until we hit a different remainder
	// i.e. we can iterate just K times, next iteration will be a guaranteed (Dirichlet's principle aka Pigeonhole principle)
	for i := 1; i <= k; i++ {
		remainder = (10*remainder + 1) % k

		if remainder == 0 { // result found -> return count of 1s
			return i
		}

		if m[remainder] { // duplicate reminder found -> no solution
			return -1
		}

		m[remainder] = true
	}

	return -1
}

func test(k int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("K - divisor: %v \n", k)

	result := smallestRepunitDivByK(k)

	fmt.Printf("Number of digits of 111.... number that is divisible by %v: %v \n", k, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(1, 1)
}

func test2() {
	test(2, -1) // // 111... cannot end with even number
}

func test3() {
	test(3, 3)
}

func test4() {
	test(5, -1) // 111... cannot end with 5
}

func main() {
	// 1015. Smallest Integer Divisible by K
	test1()
	test2()
	test3()
	test4()
}
