package main

import "fmt"

func minSubarray(nums []int, p int) int {
	// p - the divisor

	// (totalSum - subarraySum) % p = 0
	// totalSum % p = subarraySum % p
	// Let's assign k = totalSum % p

	// if k == 0, then answer is 0 -> the array itself is divisible by P

	// subarraySum % p = k
	// for indexes i < j:
	// (pS[j] - pS[i]) % p = k
	// since k % p = k, we can replace k > k % p
	// (pS[j] - pS[i]) % p = k % p
	// pS[i] % p = (pS[j] - k) % p

	// to avoid negative mods, we replace
	// (pS[j] - k) % p -> (pS[j] - k + p) % p

	// pS[i] % p = (pS[j] - k + p) % p
	// targetMod = (pS[j] - k + p) % p

	// And to minimize the (j - i) to have the smallest subarray for the current position i.
	// So we search for the biggest i < j where pS[i] % p == target
	// We keep the map of [modP] -> latest index before current position.
	// Then, if we found an (i = latestIndex[target]), the length of the subarray is (j - i).

	// We're using -1-based sums: pS[-1] = 0 and then pS[i] = a[0] + ... + a[i] (inclusive i)

	// I.e. we don't keep the pS array itself, we only need a map up to the current index j:
	// map[modP] -> last index i before current position with pS[i] % p = modP

	// calculate K - total array sum % p
	k := 0

	for _, v := range nums {
		k = (k + v) % p
	}

	//fmt.Printf("K (sum of the array %% %v) = %v \n", p, k)

	if k == 0 { // sum of the array is divisible by P -> nothing to remove
		return 0
	}

	// !! prefixSums -> -1 based
	m := make(map[int]int) // modP -> last index of pS[i] % p == modP before current position
	m[0] = -1              // with 0 elements (i = -1), we have modP = 0 since sum is 0

	minArrayLength := len(nums)

	prefixSumJModP := 0

	for j, v := range nums { // j > i, i - prefix indexes in the map
		// pS[j] % p
		prefixSumJModP = (prefixSumJModP + v) % p

		// calculate the target modP for the previous index, so that (pS[j] - pS[i]) % p = k
		// +p is the trick to avoid the negative remainders
		// targetMod (pS[j] - k + p) % p
		targetModP := (prefixSumJModP - k + p) % p

		if lastIndex, ok := m[targetModP]; ok { // targetMod found in earlier indexes
			subarrayLen := j - lastIndex // length of subarray = (j - i)

			minArrayLength = min(minArrayLength, subarrayLen)
		}

		// in any case, last index of the current mod is j
		m[prefixSumJModP] = j
	}

	if minArrayLength < len(nums) { // we found a subarray shorter than the whole array -> return length of this subarray
		return minArrayLength
	}

	return -1
}

func test(arr []int, p int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("P (modulus divisor): %v \n", p)

	result := minSubarray(arr, p)

	fmt.Printf("Min subarray size: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{3, 1, 4, 2}
	p := 6
	expected := 1 // remove 4, 3 + 1 + 2 = 6

	test(arr, p, expected)
}

func test2() {
	arr := []int{6, 3, 5, 2}
	p := 9
	expected := 2 // remove 5, 2, 6 + 3 = 9

	test(arr, p, expected)
}

func test3() {
	arr := []int{1, 2, 3}
	p := 3
	expected := 0 // nothing to remove, 1 + 2 + 3 = 6

	test(arr, p, expected)
}

func main() {
	// 1590. Make Sum Divisible by P
	test1()
	test2()
	test3()
}
