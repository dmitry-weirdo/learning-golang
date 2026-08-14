package main

import "fmt"

func closestPrimes(left int, right int) []int {
	primes := primesFromToIncluding(left, right)
	fmt.Printf("Prime numbers from %v to %v: %v \n", left, right, primes)

	if len(primes) < 2 { // not enough primes to get pairs
		return []int{-1, -1}
	}

	closestPair := []int{primes[0], primes[1]}

	for i := 1; i <= len(primes)-2; i++ {
		if (primes[i+1] - primes[i]) < (closestPair[1] - closestPair[0]) {
			closestPair[0] = primes[i]
			closestPair[1] = primes[i+1]
		}
	}

	return closestPair
}

func primesFromToIncluding(left int, n int) []int {
	// Sieve of Eratosthenes
	if n < 2 {
		return []int{}
	}

	// to not init all to true in a separate O(n) cycle, we use the negation
	isNotPrime := make([]bool, n+1)

	// 1 is NOT prime
	isNotPrime[1] = true

	// Remove multiples of each prime.
	for p := 2; p*p <= n; p++ {
		if !isNotPrime[p] {
			for multiple := p * p; multiple <= n; multiple += p {
				isNotPrime[multiple] = true
			}
		}
	}

	// Collect the primes in [left; right] range
	primes := []int{}
	for i := left; i <= n; i++ {
		if !isNotPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func test(left, right int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Left: %v \n", left)
	fmt.Printf("Right: %v \n", right)

	result := closestPrimes(left, right)

	fmt.Printf("Prime numbers in [%v; %v] range: %v \n", left, right, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	test(
		10,
		19,
		[]int{11, 13},
	)
}

func test2() {
	test(
		4,
		6,
		[]int{-1, -1}, // no pairs in between
	)
}

func test3() {
	test(
		5,
		7,
		[]int{5, 7}, // just 1 pair
	)
}

func main() {
	// 2523. Closest Prime Numbers in Range
	test1()
	test2()
	test3()
}
