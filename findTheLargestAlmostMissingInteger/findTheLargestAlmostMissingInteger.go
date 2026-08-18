package main

import "fmt"

func largestInteger(nums []int, k int) int {
	n := len(nums)

	// case 1: k == n -> every element is just in 1 array -> just select the biggest element in the array
	// e.g. [0, 0], k = 2 -> return is 0, not -1
	// we do NOT need the frequency map in this case.
	if k == n {
		maxInArray := -1

		for _, v := range nums {
			maxInArray = max(maxInArray, v)
		}

		return maxInArray
	}

	freq := make(map[int]int) // int to count

	for _, v := range nums {
		freq[v]++
	}

	// case 2: (1 < k < n) -> all the values except the first and the last will be in > 1 subarrays
	maxValue := -1 // values are in range [0; 50]

	if (1 < k) && (k < n) {
		validFirst := freq[nums[0]] == 1
		validLast := freq[nums[n-1]] == 1

		if validFirst {
			maxValue = max(maxValue, nums[0])
		}

		if validLast {
			maxValue = max(maxValue, nums[n-1])
		}

		return maxValue
	}

	// case 3: k == 1 -> we need to select the biggest value with (freq = 1) in the whole array
	if k == 1 {
		for key, v := range freq {
			if v == 1 && key > maxValue {
				maxValue = key
			}
		}

		return maxValue
	}

	// incorrect input, e.g. k <= 0 or k > n
	panic(fmt.Sprintf("This must never happen. Incorrect values k = %v, n = %v.", k, n))
}

func test(arr []int, k int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("K - subarray window size: %v \n", k)

	result := largestInteger(arr, k)

	fmt.Printf("Max element found in just one subarray of size %v: %v \n", k, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	// select of 1st and last if they have frequency 1
	test([]int{3, 9, 2, 1, 7}, 3, 7)
}

func test2() {
	// select of 1st and last if they have frequency 1. Last element 7 has freq = 2.
	test([]int{3, 9, 7, 2, 1, 7}, 4, 3)
}

func test3() {
	// select of all elements with freq = 1.
	test([]int{0, 0}, 1, -1)
}

func test4() {
	// select of all elements with freq = 1.
	test([]int{1, 2}, 1, 2)
}

func test5() {
	// full array window -> select biggest, don't care about duplicates
	test([]int{0, 0}, 2, 0)
}

func test6() {
	// full array window -> select biggest, don't care about duplicates
	test([]int{0, 1, 1, 1, 0}, 5, 1)
}

func main() {
	// 3471. Find the Largest Almost Missing Integer
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
