package main

import "fmt"

func minSumOfLengths(arr []int, target int) int {
	// not faster, passes in 65-105 ms
	return minSumOfLengths_prefixSums_arrAsDp(arr, target)

	// passes in 57-125 ms, very unstable LeetCode
	//return minSumOfLengths_prefixSums_separateDpArray(arr, target)
}

func minSumOfLengths_prefixSums_arrAsDp(arr []int, target int) int {
	n := len(arr)

	// end position [i] to prefixSum
	// since we only have > 0 values, this sum will only increase and will be unique
	m := make(map[int]int)
	m[0] = 0

	prefixSum := 0

	// dp[i] - min subarray ending before or at index [i]
	// we can use arr itself as DP to save memory, see

	minLengthsSum := n + 1
	minLength := n + 1

	for i := range n {
		//fmt.Printf("i: %v, dp: %v \n", i, arr)

		prefixSum += arr[i]

		requiredDiff := prefixSum - target
		if j, ok := m[requiredDiff]; ok { // found the required targetSum
			length := i - j + 1
			//fmt.Printf("Found the required targetSum %v for indices [%v; %v]. Length of subarray: %v \n", target, j, i, length)

			if j >= 1 && arr[j-1] > 0 && arr[j-1] <= n { // there was an earlier subarray before of equal to j
				//fmt.Printf("DP[%v] = %v \n", j, arr[j-1])
				minLengthsSum = min(minLengthsSum, length+arr[j-1])
			}

			// check whether a subarray ending on [i] is the shortest so far
			minLength = min(minLength, length)
		}

		arr[i] = minLength

		m[prefixSum] = i + 1
	}

	//fmt.Printf("DP (min length of target subarray ending on index i or before): %v \n", arr)

	if minLengthsSum > n { // not a single pair of subarrays with target sum found
		return -1
	}

	return minLengthsSum
}

func minSumOfLengths_prefixSums_separateDpArray(arr []int, target int) int {
	n := len(arr)

	// end position [i] to prefixSum
	// since we only have > 0 values, this sum will only increase and will be unique
	m := make(map[int]int)
	m[0] = 0

	prefixSum := 0

	// dp[i] - min subarray ending before or at index [i]
	// we can use arr itself as DP to save memory, see minSumOfLengths_prefixSums_arrAsDp
	dp := make([]int, n)

	minLengthsSum := n + 1
	minLength := n + 1

	for i := range n {
		//fmt.Printf("i: %v, dp: %v \n", i, dp)

		prefixSum += arr[i]

		requiredDiff := prefixSum - target
		if j, ok := m[requiredDiff]; ok { // found the required targetSum
			length := i - j + 1
			//fmt.Printf("Found the required targetSum %v for indices [%v; %v]. Length of subarray: %v \n", target, j, i, length)

			if j >= 1 && dp[j-1] > 0 && dp[j-1] <= n { // there was an earlier subarray before of equal to j
				//fmt.Printf("DP[%v] = %v \n", j, dp[j-1])
				minLengthsSum = min(minLengthsSum, length+dp[j-1])
			}

			// check whether a subarray ending on [i] is the shortest so far
			minLength = min(minLength, length)
		}

		dp[i] = minLength

		m[prefixSum] = i + 1
	}

	//fmt.Printf("DP (min length of target subarray ending on index i or before): %v \n", dp)

	if minLengthsSum > n { // not a single pair of subarrays with target sum found
		return -1
	}

	return minLengthsSum
}

func test(arr []int, target int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Target subarray sum: %v \n", target)

	result := minSumOfLengths(arr, target)

	fmt.Printf("Min sum of lengths of subarrays with sum = %v: %v \n", target, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{3, 2, 2, 4, 3}, 3, 2) // [3] and [3]
}

func test2() {
	test([]int{7, 3, 4, 7}, 7, 2) // [7] and [7]
}

func test3() {
	test([]int{4, 3, 2, 6, 2, 3, 4}, 6, -1) // just one sub-array [6] with sum 6
}

func test4() {
	test([]int{1, 1}, 1, 2) // [1] and [1]
}

func main() {
	// 1477. Find Two Non-overlapping Sub-arrays Each With Target Sum
	test1()
	test2()
	test3()
	test4()
}
