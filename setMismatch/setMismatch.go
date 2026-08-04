package main

import "fmt"

func findErrorNums(nums []int) []int {
	return findErrorNums_optimized(nums)
	//return findErrorNums_trivial(nums)
}

func findErrorNums_optimized(nums []int) []int {
	// we're using the fact that all values are positive
	// array indexes are [0...n - 1]
	// expected values are [1...n]

	// If we find a value i, we negate the value [i - 1]
	// If it is already negative, then we know that i is the duplicate.
	// We leave a[i] as positive.
	// Therefore, all elements will be marked negative except the element of [j - 1], where j is the negative index

	duplicate := 0

	for _, v := range nums {
		n := abs(v)

		if nums[n-1] < 0 {
			duplicate = n
			// do not negate nums[n - 1] again, it is already negative
		} else {
			nums[n-1] = -nums[n-1]
		}
	}

	// a single [i] element that remained positive -> missing value is [i + 1]
	missing := 0

	for i, v := range nums {
		if v > 0 { // we found an index that was not marked to negative value
			missing = i + 1
			break
		}
	}

	return []int{duplicate, missing}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func findErrorNums_trivial(nums []int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	missing := 0
	duplicate := 0

	for i := 1; i <= len(nums) && (missing == 0 || duplicate == 0); i++ {
		if v, ok := m[i]; !ok {
			missing = i
		} else if v > 1 {
			duplicate = i
		}
	}

	return []int{duplicate, missing}
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)

	result := findErrorNums(arr)

	fmt.Printf("Array after finding duplicate and missing elements: %v \n", arr)
	fmt.Printf("Duplicate element, missing element: %v \n", result)
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
	arr := []int{1, 2, 2, 4}
	expected := []int{2, 3}

	test(arr, expected)
}

func test2() {
	arr := []int{1, 1}
	expected := []int{1, 2}

	test(arr, expected)
}

func main() {
	// 645. Set Mismatch
	test1()
	test2()
}
