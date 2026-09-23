package main

import "fmt"

func minOperations(nums []int, x int) int {
	// we basically need to find max length subarray with windowSum = totalSum - x.
	totalSum := sumOfArray(nums)
	targetSubarraySum := totalSum - x

	//fmt.Printf("Array length: %v \n", len(nums))
	//fmt.Printf("Total array sum: %v \n", totalSum)
	//fmt.Printf("Target subarray sum: %v \n", targetSubarraySum)

	if targetSubarraySum < 0 {
		// array contains only positive values -> we cannot get to value > totalSum
		// notably, we can get 0 with all elements in the array
		return -1
	}

	if targetSubarraySum == 0 {
		// since all values in the array are positive -> we have to use all values in the array
		return len(nums)
	}

	maxWindowLength := -1

	left := 0
	sum := 0

	for right, v := range nums {
		sum += v

		//fmt.Println()
		//fmt.Printf("Right: %v, sum [%v; %v] = %v \n", right, left, right, sum)

		for sum > targetSubarraySum { // shrink from left
			// we can just subtract from left, since all values in the array are non-negative
			sum -= nums[left]
			left++
		}

		if sum == targetSubarraySum {
			windowLength := right - left + 1
			maxWindowLength = max(maxWindowLength, windowLength)
		}
	}

	if maxWindowLength < 0 {
		return -1
	}

	return len(nums) - maxWindowLength
}

func sumOfArray(arr []int) int {
	// we assume the array is non-empty
	sum := 0

	for _, v := range arr {
		sum += v
	}

	return sum
}

func test(arr []int, x int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("X: %v \n", x)

	result := minOperations(arr, x)

	fmt.Printf("Minimum array operations to reduce X = %v to 0: %v \n", x, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{1, 1, 4, 2, 3},
		5,
		2,
	)
}

func test2() {
	test(
		[]int{5, 6, 7, 8, 9},
		4,
		-1,
	)
}

func test3() {
	test(
		[]int{3, 2, 20, 1, 1, 3},
		10,
		5,
	)
}

func test4() {
	// test-case 5 / 97
	test(
		[]int{1, 1},
		3,
		-1,
	)
}

func test5() {
	// test-case 90 / 97
	test(
		[]int{8828, 9581, 49, 9818, 9974, 9869, 9991, 10000, 10000, 10000, 9999, 9993, 9904, 8819, 1231, 6309},
		134365,
		16, // we have to use the whole array
	)
}

func main() {
	// 1658. Minimum Operations to Reduce X to Zero
	test1()
	test2()
	test3()
	test4()
	test5()
}
