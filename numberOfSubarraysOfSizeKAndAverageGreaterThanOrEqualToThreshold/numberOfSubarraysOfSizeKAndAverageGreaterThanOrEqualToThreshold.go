package main

import "fmt"

func numOfSubarrays(arr []int, k int, threshold int) int {
	// we don't want to do average and floats
	// since threshold <= 10^4 and k <= 10^5, the sum will be <= 10^9 and will fit into int
	targetSum := threshold * k
	fmt.Printf("targetSum: %v \n", targetSum)

	sum := 0

	result := 0 // count of windows of size K with sum > targetSum

	left := 0

	for right := 0; right < len(arr); right++ {
		sum += arr[right]

		if right < k-1 { // first (k - 1) elements we just sum up
			// k = 3
			// 0, 1 - do nothing
			// 2 - array [0; 2] can be already added to the result
			continue
		}

		if right >= k { // move the window if right >= k
			// k = 3
			// 0, 1, 2 - don't move
			// 3 - move left from 0 to 1 -> we have indexes 1, 2, 3
			sum -= arr[left]

			left++
		}

		if sum >= targetSum {
			//fmt.Printf("Sum of arr[%v; %v] = %v >= %v \n", left, right, sum, targetSum)
			//fmt.Printf("Subarray of %v: %v \n", k, arr[left:right+1])

			result++
		}
	}

	return result
}

func test(arr []int, k, threshold int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("k (length of the subarray): %v \n", k)
	fmt.Printf("Threshold: %v \n", threshold)

	result := numOfSubarrays(arr, k, threshold)

	fmt.Printf("Total subarrays of size %v with sum >= %v: %v \n", k, threshold, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}
	k := 3
	threshold := 5
	expected := 6

	test(arr, k, threshold, expected)
}

func test2() {
	arr := []int{2, 2, 2, 2, 5, 5, 5, 8}
	k := 3
	threshold := 4
	expected := 3

	test(arr, k, threshold, expected)
}

func main() {
	// 1343. Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold
	test1()
	test2()
}
