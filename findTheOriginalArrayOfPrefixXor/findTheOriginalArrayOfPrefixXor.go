package main

import "fmt"

func findArray(pref []int) []int {
	// working on the pref array itself -> no need for time/space on a separate array
	// yes, it passes in 0-2 ms
	return findArray_inPlace(pref)

	//// allocating a separate array -> works, but slowly, 5+ ms
	//return findArray_naive(pref)
}

func findArray_inPlace(pref []int) []int {
	prev := pref[0]
	buf := 0

	for i := 1; i < len(pref); i++ {
		buf = pref[i] // save until we changed the pref[i]

		pref[i] = pref[i] ^ prev

		prev = buf
	}

	return pref
}

func findArray_naive(pref []int) []int {
	// x ^ a = b
	// -> x = a ^ b

	// Proof:
	// x ^ a = b
	// x ^ a ^ a = b ^ a
	// a ^ a == 0 (all bits double-negate)
	// b ^ a = a ^ b (xor is obviously commutative)
	// x ^ 0 = a ^ b
	// x ^ 0 = x (1-bits remain in place when xor-ed with 0, 0-bits are xor-ed with 0 which gives 0)

	// pref[i] = a[0] ^ a[1] ^ ... ^ a[i], i.e. inclusive
	// it means that pref[0] = a[0]

	// pref[1] = a[0] ^ a[1]
	// -> a[1] = pref[1] ^ a[0]

	// pref[2] = a[0] ^ a[1] ^ a[2]
	// -> a[2] = pref[2] ^ a[0] ^ a[1]
	// but a[0] ^ a[1] = pref[1]
	// a[2] = pref[2] ^ pref[1]

	n := len(pref)

	// todo: probably we can-do in place replacement of pref to have O(1) space.
	a := make([]int, n)
	a[0] = pref[0]

	for i := 1; i < n; i++ {
		a[i] = pref[i] ^ pref[i-1]
	}

	return a
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array of prefix xors: %v \n", arr)

	result := findArray(arr)

	fmt.Printf("Restored original array: %v \n", result)
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
		[]int{5, 2, 0, 3, 1},
		[]int{5, 7, 2, 3, 2},
	)
}

func test2() {
	test(
		[]int{13},
		[]int{13},
	)
}

func main() {
	// 2433. Find The Original Array of Prefix Xor
	test1()
	test2()
}
