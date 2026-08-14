package main

import "fmt"

func countPrimes(n int) int {
	// we don't need to collect primes, we just need the count
	return countPrimes_optimized(n)

	// reusing the primesUpToMethod
	// passes in 85-90 max
	//return countPrimes_naive(n)
}

func countPrimes_optimized(n int) int {
	// going O(n) once
	// it's actually slower to 80-85 ms again
	// probably the internal iteration on every digit is still slow
	//return countPrimesUpToExcludingSelf_2(n)

	// going up to sqrt(n), then full iteration up to(n)
	// yes, it speeds up to 60-70 ms
	// However, we can just to one O(n) run, and not up to square.
	return countPrimesUpToExcludingSelf_1(n)
}

func countPrimes_naive(n int) int {
	primes := primesUpTo(n)

	if len(primes) < 1 {
		return 0
	}

	if primes[len(primes)-1] == n { // we need the values < n, so if n is prime, don't count it
		return len(primes) - 1
	}

	return len(primes)
}

func primesUpTo(n int) []int {
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

	// Collect the primes.
	primes := []int{}
	for i := 2; i <= n; i++ {
		if !isNotPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func countPrimesUpToExcludingSelf_1(n int) int {
	// Sieve of Eratosthenes
	if n < 2 {
		return 0
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

	count := 0

	// Collect the primes.
	for i := 2; i <= n; i++ {
		if !isNotPrime[i] {
			count++
		}
	}

	// exclude n itself
	if !isNotPrime[n] {
		count--
	}

	return count
}

func countPrimesUpToExcludingSelf_2(n int) int {
	// Sieve of Eratosthenes
	if n < 2 {
		return 0
	}

	// to not init all to true in a separate O(n) cycle, we use the negation
	isNotPrime := make([]bool, n+1)

	// 1 is NOT prime
	isNotPrime[1] = true

	count := 0

	// we're going up to (n -1) to count in this cycle.
	// N self is not counted -> n, not n-1
	// Remove multiples of each prime.
	for p := 2; p <= n-1; p++ {
		if !isNotPrime[p] {
			count++

			for multiple := p * p; multiple <= n; multiple += p {
				isNotPrime[multiple] = true
			}
		}
	}

	return count
}

func test(n int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", n)

	result := countPrimes(n)

	fmt.Printf("Prime factors of %v that are < %v: %v \n", n, n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(10, 4) // 2, 3, 5, 7
}

func test2() {
	test(2, 0) // 2 does not count since we count primes < N
}

func test3() {
	test(0, 0)
}

func test4() {
	test(1, 0) // 1 is not prime, and we count < N anyway
}

func main() {
	// 204. Count Primes
	test1()
	test2()
	test3()
	test4()
}
