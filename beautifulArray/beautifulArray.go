package main

import (
	"fmt"
	"math/bits"
	"slices"
)

func beautifulArray(n int) []int {
	return beautifulArray_divideAndConquer(n) // divide and conquer, O(n * log n) (same as quickSort or mergeSort)
	//return beautifulArray_bitsReversedSort(n) // nice and just sorting O(n * log n)
}

func beautifulArray_divideAndConquer(n int) []int {
	// 2*a[k] != a[i] + a[j]

	// if a[i] is even, a[j] is odd, we satisfy this condition
	// -> so we're trying to construct an array where left part is odd and right part is even
	// (could also be left even, right odd).

	// !!! we can apply any affine transformation y = qx + w to this inequality, and it will still be true:
	// q * (2*a[k]) + w != q * (a[i] + a[j]) + w
	// i.e. after applying any transformation to a "beautiful" array, it will remain beautiful.

	// to make left part consist of odd numbers only, we apply (2*x - 1) transformation to the left array
	// to make right part consist of even numbers only, we apply (2*x) transformation to the right array

	// cache the same values
	dp := make(map[int][]int)

	// base-case - for 1 we return an array of [1]
	dp[1] = []int{1}

	var dfs func(a int) []int

	dfs = func(a int) []int {
		// if ()

		// !!! Getting from memo will NOT work.
		// Example: N = 5, 1 2 3 4 5
		// N(3) ->

		if precalculated, ok := dp[a]; ok {
			fmt.Printf("N = %v, memo hit: returning dp[%v] = %v \n", a, a, precalculated)
			return copyArray(precalculated)
		}

		leftNumber := a/2 + a%2 // for odd number, it will be 5 -> 3, for even number, it will be 4 -> 2
		rightNumber := a - leftNumber

		left := dfs(leftNumber)
		right := dfs(rightNumber)

		fmt.Println()
		fmt.Printf("N: %v \n", a)
		fmt.Printf("Left array (%v): %v \n", leftNumber, left)
		fmt.Printf("Right array (%v): %v \n", rightNumber, right)

		// apply (2 * x - 1) transformation to left
		for i, v := range left {
			left[i] = 2*v - 1
		}

		// apply (2 * x) transformation to right
		for i, v := range right {
			right[i] = 2 * v
		}

		fmt.Printf("Left array after (2*x - 1): %v \n", left)
		fmt.Printf("Right array after (2*x): %v \n", right)

		result := append(left, right...) // !!! may reuse the memory of left

		// 100% separate copy
		//result := make([]int, 0, len(left)+len(right))
		//result = append(result, left...)
		//result = append(result, right...)

		dp[a] = copyArray(result)

		fmt.Printf("Set dp[%v] = %v \n", a, result)

		return result
	}

	return dfs(n)
}

func copyArray(arr []int) []int {
	arrayCopy := make([]int, len(arr))
	copy(arrayCopy, arr)
	return arrayCopy
}

func beautifulArray_bitsReversedSort(n int) []int {
	// See https://www.youtube.com/watch?v=O-1ucu8ErEo

	// super intellectual, we sort by reversed significant bytes.
	// last byte defines the odd/even etc

	// 1   2   3   4   5   6
	// 001 010 011 100 101 110 -> straight bytes
	// 100 010 110 001 101 011 -> reversed bytes

	// if we sort just for last original byte:
	// 2   4   6   1   3   5
	// 010 100 110 001 011 101 -> straight bytes, first 0, then 1

	// then sort by the 2nd last byte and 3rd last byte
	// 4   2   6   1   5   3
	// 100 010 110 001 101 011  -> straight bytes, first 0, then 1

	a := make([]uint16, n) // we only need 2 bytes for 1-1000 range

	for i := range n {
		a[i] = uint16(i + 1) // numbers start from 1
	}

	slices.SortFunc(a, func(x, y uint16) int {
		ux := bits.Reverse16(x)
		uy := bits.Reverse16(y)

		//fmt.Printf("x: %v \n", x)
		//fmt.Printf("x as binary: %b\n", x)
		//fmt.Printf("x reversed binary: %b\n", ux)

		//return int(ux) - int(uy) // does not pass tests, although also valid :facepalm:
		return int(ux) - int(uy)
	})

	b := make([]int, n)

	for i := range a {
		b[i] = int(a[i])
	}

	fmt.Printf("Sorted array: %v \n", b)

	return b
}

func main() {
	// 932. Beautiful Array

	n := 5
	result := beautifulArray(n)

	fmt.Printf("Beautiful array (%v): %v \n", n, result)

	// todo: no tests since the return order in LeetCode tests is not like in our algorithm
	// todo: we still can write tests for our ordering
}
