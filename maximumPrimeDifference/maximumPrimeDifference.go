package main

import "fmt"

func maximumPrimeDifference(nums []int) int {
	// we can also get maxValue from the array
	const maxValue = 100

	primes := primesUpTo(maxValue)
	//fmt.Printf("Prime number up to %v: %v \n", maxValue, primes)

	p := make([]bool, maxValue+1) // shortcut to check whether a number is prime

	for _, v := range primes {
		p[v] = true
	}

	// search for first prime index
	i := 0
	for !p[nums[i]] { // we're not checking the borders since we're guaranteed that there is at least 1 prime number in the array
		i++
	}

	// search for last prime index
	j := len(nums) - 1
	for !p[nums[j]] { // we're not checking the borders since we're guaranteed that there is at least 1 prime number in the array
		j--
	}

	return j - i
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

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := maximumPrimeDifference(arr)

	fmt.Printf("Index diff of first and last prime number: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{4, 2, 9, 5, 3}, 3) // a[1] - a[4]
}

func test2() {
	test([]int{4, 8, 2, 8}, 0) // a[2] - a[2]
}

func main() {
	// 3115. Maximum Prime Difference
	test1()
	test2()
}
