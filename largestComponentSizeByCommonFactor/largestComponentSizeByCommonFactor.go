package main

import "fmt"

func largestComponentSize(nums []int) int {
	// pre-compute the prime numbers up to max value to only iterate the prime numbers when calculating prime divisors
	// also, instead of map to group the values by unionFindParent, use an array [0..maxValue] instead of map
	/// this decreased time a bit - to 75-85 ms
	return largestComponentSize_unionFind_optimized(nums)

	// usual solution, no optmizations
	// passes in 100+ ms
	//return largestComponentSize_unionFind(nums)
}

func largestComponentSize_unionFind_optimized(nums []int) int {
	// find max number
	maxValue := maxInArray(nums)

	primesUpToMaxValue := primesUpTo(maxValue)
	//fmt.Printf("Primes up to %v: %v \n", maxValue, primesUpToMaxValue)

	// we will join positive numbers in range maxValue
	uf := newUnionFind(maxValue + 1) // we count from 0, values are [1; maxValue]

	// for every number, we split it to (unique prime divisors + number itself)
	// and union these groups together

	// Example:
	// 4 = 2, 4 -> merge (2, 4)
	// 6 = 2, 3, 6 -> merge (2, 3, 6) -> group is (2, 3, 4, 6)
	for _, v := range nums {
		primeFactorsAndSelf := uniquePrimeFactorsAndSelfWithPrecalculatedPrimeNumbers(v, primesUpToMaxValue)

		//fmt.Printf("%v -> self and prime factors -> %v \n", v, primeFactorsAndSelf)

		// union all prime factors and value itself
		firstPrimeFactor := primeFactorsAndSelf[0]

		for _, p := range primeFactorsAndSelf {
			uf.union(p, firstPrimeFactor)
		}
	}

	// The merged groups contain also the prime factors, that were not in the original array.
	// So we cannot return the max group size.
	// We're going through original values, get a merged group for every value and count original values by group.
	maxGroupSize := 1

	//m := make(map[int]int) // groupIndex -> count

	// map is slow -> let's use an array -> additional O(maxValue) space, but faster time
	m := make([]int, maxValue+1)

	// values in the array are unique, so we don't need to check whether this number was already visited and counted
	for _, v := range nums {
		group := uf.find(v)
		m[group]++

		maxGroupSize = max(maxGroupSize, m[group])
	}

	return maxGroupSize
}

func largestComponentSize_unionFind(nums []int) int {
	// find max number
	maxValue := maxInArray(nums)

	// we will join positive numbers in range maxValue
	uf := newUnionFind(maxValue + 1) // we count from 0, values are [1; maxValue]

	// for every number, we split it to (unique prime divisors + number itself)
	// and union these groups together

	// Example:
	// 4 = 2, 4 -> merge (2, 4)
	// 6 = 2, 3, 6 -> merge (2, 3, 6) -> group is (2, 3, 4, 6)
	for _, v := range nums {
		primeFactorsAndSelf := uniquePrimeFactorsAndSelf(v)

		//fmt.Printf("%v -> self and prime factors -> %v \n", v, primeFactorsAndSelf)

		// union all prime factors and value itself
		firstPrimeFactor := primeFactorsAndSelf[0]

		for _, p := range primeFactorsAndSelf {
			uf.union(p, firstPrimeFactor)
		}
	}

	// The merged groups contain also the prime factors, that were not in the original array.
	// So we cannot return the max group size.
	// We're going through original values, get a merged group for every value and count original values by group.
	maxGroupSize := 1

	m := make(map[int]int) // groupIndex -> count

	// values in the array are unique, so we don't need to check whether this number was already visited and counted
	for _, v := range nums {
		group := uf.find(v)
		m[group]++

		maxGroupSize = max(maxGroupSize, m[group])
	}

	return maxGroupSize
}

func uniquePrimeFactorsAndSelf(n int) []int {
	if n == 1 {
		return []int{1}
	}

	primeFactors := uniquePrimeFactors(n)

	if primeFactors[len(primeFactors)-1] != n { // value is NOT prime -> add it to result as well
		primeFactors = append(primeFactors, n)
	}

	return primeFactors
}

func uniquePrimeFactors(n int) []int { // gets unique prime factors
	// todo: handle n == 1

	// todo: we can use precalculated prime numbers, but it isn't necessary
	factors := make([]int, 0)

	// todo: fix the potential overflow on (d * d)
	for d := 2; d*d <= n; d++ { // iterate up to ceil(sqrt(n))
		if n%d == 0 {
			factors = append(factors, d) // only append this prime factor once

			// Remove all occurrences of d, so that all non-prime will drop as well (dividing by 2 will remove 4, 6 etc).
			for n%d == 0 {
				n /= d
			}
		}
	}

	// if the remaining value is prime (not 1 after all the divisions) -> add it as well
	// e.g. 7 -> 7 will remain
	// 26 -> 13 will remain
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

func uniquePrimeFactorsAndSelfWithPrecalculatedPrimeNumbers(n int, primeNumbers []int) []int {
	if n == 1 {
		return []int{1}
	}

	primeFactors := uniquePrimeFactorsWithPrecalculatedPrimeNumbers(n, primeNumbers)

	if primeFactors[len(primeFactors)-1] != n { // value is NOT prime -> add it to result as well
		primeFactors = append(primeFactors, n)
	}

	return primeFactors
}

func uniquePrimeFactorsWithPrecalculatedPrimeNumbers(n int, primeNumbers []int) []int { // gets unique prime factors
	// todo: handle n == 1

	factors := make([]int, 0)

	for _, d := range primeNumbers { // we only iterate over the prime numbers
		// todo: fix the potential overflow on (d * d)
		if d*d > n {
			break
		}

		if n%d == 0 {
			factors = append(factors, d) // only append this prime factor once

			// Remove all occurrences of d, so that all non-prime will drop as well (dividing by 2 will remove 4, 6 etc).
			for n%d == 0 {
				n /= d
			}
		}
	}

	// if the remaining value is prime (not 1 after all the divisions) -> add it as well
	// e.g. 7 -> 7 will remain
	// 26 -> 13 will remain
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

func maxInArray(arr []int) int {
	// we assume the array is non-empty
	m := arr[0]

	for _, v := range arr {
		m = max(m, v)
	}

	return m
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

type UnionFind struct {
	parents []int // if parent[i] = i, it is the root, else it's the index of the parent
	sizes   []int // sizes of the tree for every element
}

func newUnionFind(n int) UnionFind {
	parents := make([]int, n)
	sizes := make([]int, n)

	for i := range n {
		// every group is just a root
		parents[i] = i

		// every group has a size of 1
		sizes[i] = 1
	}

	return UnionFind{
		parents: parents,
		sizes:   sizes,
	}
}

func (uf UnionFind) find(x int) int { // recursive version
	if uf.parents[x] == x { // parent points to itself -> reached the root
		return x
	}

	// path compression -> set the root to every parents[i] in the chain
	uf.parents[x] = uf.find(uf.parents[x])

	return uf.parents[x]
}

func (uf UnionFind) print() {
	fmt.Printf("Parents: %v \n", uf.parents)
	fmt.Printf("Sizes: %v \n", uf.sizes)
}

func (uf UnionFind) union(x, y int) bool { // returns false if they're already in the same set
	// these find will perform path compression
	rootX := uf.find(x)
	rootY := uf.find(y)

	//fmt.Printf("root of %d: %d, root of %d: %d\n", x, rootX, y, rootY)

	// x and y are already in the same set -> nothing to merge
	if rootX == rootY {
		//fmt.Printf("Element %v and %v already belong to the same root %v. Nothing to merge. \n", x, y, rootX)
		return false
	}

	// merge the smaller group into the bigger group
	// todo: ideally, we should merge the tree with smaller depth into the tree with bigger depth
	if uf.sizes[rootX] < uf.sizes[rootY] { // merge x into y
		//fmt.Printf("sizes[%v] = %v < sizes[%v] = %v. Merging root %v into root %v \n", rootX, uf.sizes[rootX], rootY, uf.sizes[rootY], rootX, rootY)

		uf.parents[rootX] = rootY
		uf.sizes[rootY] += uf.sizes[rootX]
	} else { // merge y into x
		//fmt.Printf("sizes[%v] = %v >= sizes[%v] = %v. Merging root %v into root %v \n", rootX, uf.sizes[rootX], rootY, uf.sizes[rootY], rootY, rootX)

		uf.parents[rootY] = rootX
		uf.sizes[rootX] += uf.sizes[rootY]
	}

	return true
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := largestComponentSize(arr)

	fmt.Printf("Max group size of joined by common divisors > 1: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{4, 6, 15, 35}
	expected := 4 // all values joined

	test(arr, expected)
}

func test2() {
	arr := []int{20, 50, 9, 63}
	expected := 2 // (20, 50) - (9, 63)

	test(arr, expected)
}

func test3() {
	arr := []int{2, 3, 6, 7, 4, 12, 21, 39}
	expected := 8 // all values joined

	test(arr, expected)
}

func test4() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := 6 // (1) - (2, 3, 4. 6, 8, 9) - (5) - (7)

	test(arr, expected)
}

func main() {
	// 952. Largest Component Size by Common Factor
	test1()
	test2()
	test3()
	test4()
}
